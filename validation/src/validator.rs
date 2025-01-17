use futures::TryStreamExt;

use crate::nats_types::ValidateRequest;

const SCRIPT_CONCURRENCY:usize=16;

enum Policy{
	Allowed,
	Blocked,
	Delete,
	Replace(String),
}

#[allow(dead_code)]
#[derive(Debug)]
pub enum ValidateError{
	Blocked,
	NotAllowed,
	Get(rbx_asset::cookie::GetError),
	Json(serde_json::Error),
	ReadDom(ReadDomError),
	ApiGetScriptPolicy(api::Error),
	ApiGetScript(api::Error),
	ApiUpdateSubmissionModel(api::Error),
	ApiActionSubmissionValidate(api::Error),
	WriteDom(rbx_binary::EncodeError),
	Upload(rbx_asset::cookie::UploadError),
	Create(rbx_asset::cookie::CreateError),
}
impl std::fmt::Display for ValidateError{
	fn fmt(&self,f:&mut std::fmt::Formatter<'_>)->std::fmt::Result{
		write!(f,"{self:?}")
	}
}
impl std::error::Error for ValidateError{}

pub struct Validator{
	roblox_cookie:rbx_asset::cookie::CookieContext,
	api:api::Context,
}

impl Validator{
	pub const fn new(
		roblox_cookie:rbx_asset::cookie::CookieContext,
		api:api::Context,
	)->Self{
		Self{
			roblox_cookie,
			api,
		}
	}
	pub async fn validate(&self,message:async_nats::jetstream::Message)->Result<(),ValidateError>{
		println!("validate {:?}",message.message.payload);
		// decode json
		let validate_info:ValidateRequest=serde_json::from_slice(&message.payload).map_err(ValidateError::Json)?;

		// download map
		let data=self.roblox_cookie.get_asset(rbx_asset::cookie::GetAssetRequest{
			asset_id:validate_info.ModelID,
			version:Some(validate_info.ModelVersion),
		}).await.map_err(ValidateError::Get)?;

		// decode dom (slow!)
		let mut dom=read_dom(&mut std::io::Cursor::new(data)).map_err(ValidateError::ReadDom)?;

		/* VALIDATE MAP */

		// collect unique scripts
		let script_refs=get_script_refs(&dom);
		let mut script_map=std::collections::HashMap::<String,Policy>::new();
		for &script_ref in &script_refs{
			if let Some(script)=dom.get_by_ref(script_ref){
				if let Some(rbx_dom_weak::types::Variant::String(source))=script.properties.get("Source"){
					script_map.insert(source.clone(),Policy::Blocked);
				}
			}
		}

		// send all script hashes to REST endpoint and retrieve the replacements
		futures::stream::iter(script_map.iter_mut().map(Ok))
		.try_for_each_concurrent(Some(SCRIPT_CONCURRENCY),|(source,replacement)|async{
			// get the hash
			let mut hasher=siphasher::sip::SipHasher::new();
			std::hash::Hasher::write(&mut hasher,source.as_bytes());
			let hash=std::hash::Hasher::finish(&hasher);

			// fetch the script policy
			let script_policy=self.api.get_script_policy_from_hash(api::ScriptPolicyHashRequest{
				hash:format!("{:x}",hash),
			}).await.map_err(ValidateError::ApiGetScriptPolicy)?;

			// write the policy to the script_map, fetching the replacement code if necessary
			*replacement=match script_policy.Policy{
				api::Policy::Allowed=>Policy::Allowed,
				api::Policy::Blocked=>Policy::Blocked,
				api::Policy::Delete=>Policy::Delete,
				api::Policy::Replace=>{
					let script=self.api.get_script(api::GetScriptRequest{
						ScriptID:script_policy.ToScriptID,
					}).await.map_err(ValidateError::ApiGetScript)?;
					Policy::Replace(script.Source)
				},
			};

			Ok(())
		})
		.await?;

		// make the replacements
		let mut modified=false;
		for &script_ref in &script_refs{
			if let Some(script)=dom.get_by_ref_mut(script_ref){
				if let Some(rbx_dom_weak::types::Variant::String(source))=script.properties.get_mut("Source"){
					match script_map.get(source.as_str()){
						Some(Policy::Blocked)=>return Err(ValidateError::Blocked),
						None=>return Err(ValidateError::NotAllowed),
						Some(Policy::Allowed)=>(),
						Some(Policy::Delete)=>{
							modified=true;
							// delete script
							dom.destroy(script_ref);
						},
						Some(Policy::Replace(replacement))=>{
							modified=true;
							*source=replacement.clone();
						},
					}
				}
			}
		}

		// if the model was validated, the submission must be changed to use the modified model
		if modified{
			// serialize model (slow!)
			let mut data=Vec::new();
			rbx_binary::to_writer(&mut data,&dom,dom.root().children()).map_err(ValidateError::WriteDom)?;

			// upload a model lol
			let model_id=if let Some(model_id)=validate_info.ValidatedModelID{
				// upload to existing id
				let response=self.roblox_cookie.upload(rbx_asset::cookie::UploadRequest{
					assetid:model_id,
					name:None,
					description:None,
					ispublic:None,
					allowComments:None,
					groupId:None,
				},data).await.map_err(ValidateError::Upload)?;

				response.AssetId
			}else{
				// create new model
				let response=self.roblox_cookie.create(rbx_asset::cookie::CreateRequest{
					name:dom.root().name.clone(),
					description:"".to_owned(),
					ispublic:true,
					allowComments:true,
					groupId:None,
				},data).await.map_err(ValidateError::Create)?;

				response.AssetId
			};

			// update the submission to use the validated model
			self.api.update_submission_model(api::UpdateSubmissionModelRequest{
				ID:validate_info.SubmissionID,
				ModelID:model_id,
				ModelVersion:1, //TODO
			}).await.map_err(ValidateError::ApiUpdateSubmissionModel)?;
		};

		// update the submission model status to validated
		self.api.action_submission_validate(
			api::SubmissionID(validate_info.SubmissionID)
		).await.map_err(ValidateError::ApiActionSubmissionValidate)?;

		Ok(())
	}
}

#[allow(dead_code)]
#[derive(Debug)]
pub enum ReadDomError{
	Binary(rbx_binary::DecodeError),
	Xml(rbx_xml::DecodeError),
	Read(std::io::Error),
	Seek(std::io::Error),
	UnknownFormat([u8;8]),
}
impl std::fmt::Display for ReadDomError{
	fn fmt(&self,f:&mut std::fmt::Formatter<'_>)->std::fmt::Result{
		write!(f,"{self:?}")
	}
}
impl std::error::Error for ReadDomError{}

fn read_dom<R:std::io::Read+std::io::Seek>(input:&mut R)->Result<rbx_dom_weak::WeakDom,ReadDomError>{
	let mut first_8=[0u8;8];
	std::io::Read::read_exact(input,&mut first_8).map_err(ReadDomError::Read)?;
	std::io::Seek::rewind(input).map_err(ReadDomError::Seek)?;
	match &first_8[0..4]{
		b"<rob"=>{
			match &first_8[4..8]{
				b"lox!"=>rbx_binary::from_reader(input).map_err(ReadDomError::Binary),
				b"lox "=>rbx_xml::from_reader(input,rbx_xml::DecodeOptions::default()).map_err(ReadDomError::Xml),
				_=>Err(ReadDomError::UnknownFormat(first_8)),
			}
		},
		_=>Err(ReadDomError::UnknownFormat(first_8)),
	}
}

fn class_is_a(class:&str,superclass:&str)->bool{
	if class==superclass{
		return true
	}
	let class_descriptor=rbx_reflection_database::get().classes.get(class);
	if let Some(descriptor)=&class_descriptor{
		if let Some(class_super)=&descriptor.superclass{
			return class_is_a(&class_super,superclass)
		}
	}
	false
}

fn recursive_collect_superclass(objects:&mut std::vec::Vec<rbx_dom_weak::types::Ref>,dom:&rbx_dom_weak::WeakDom,instance:&rbx_dom_weak::Instance,superclass:&str){
	for &referent in instance.children(){
		if let Some(c)=dom.get_by_ref(referent){
			if class_is_a(c.class.as_str(),superclass){
				objects.push(c.referent());//copy ref
			}
			recursive_collect_superclass(objects,dom,c,superclass);
		}
	}
}

fn get_script_refs(dom:&rbx_dom_weak::WeakDom)->Vec<rbx_dom_weak::types::Ref>{
	let mut scripts=std::vec::Vec::new();
	recursive_collect_superclass(&mut scripts,dom,dom.root(),"LuaSourceContainer");
	scripts
}
