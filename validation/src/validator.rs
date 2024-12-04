use futures::StreamExt;
use crate::nats_types::ValidateRequest;

struct ModelVersion{
	model_id:u64,
	model_version:u64,
}

enum Valid{
	Untouched,
	Modified(ModelVersion),
}

#[allow(dead_code)]
#[derive(Debug)]
enum ValidateError{
	Get(rbx_asset::cookie::GetError),
	Json(serde_json::Error),
	ReadDom(ReadDomError),
	ApiGetReplacements(api::GetReplacementsError),
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
	nats:async_nats::Client,
	subscriber:async_nats::Subscriber,
	roblox_cookie:rbx_asset::cookie::CookieContext,
}

impl Validator{
	pub async fn new(
		nats:async_nats::Client,
		roblox_cookie:rbx_asset::cookie::CookieContext,
	)->Result<Self,async_nats::SubscribeError>{
		Ok(Self{
			subscriber:nats.subscribe("validate").await?,
			nats,
			roblox_cookie,
		})
	}
	pub async fn run(mut self){
		while let Some(message)=self.subscriber.next().await{
			self.validate_supress_error(message).await
		}
	}
	async fn validate_supress_error(&self,message:async_nats::Message){
		match self.validate(message).await{
			Ok(valid)=>{
				unimplemented!();
				// self.api.validate(validate_response).await.unwrap();
			},
			Err(e)=>{
				println!("there was an error, oopsie! {e}");
			}
		}
	}
	async fn validate(&self,message:async_nats::Message)->Result<Valid,ValidateError>{
		println!("validate {:?}",message);
		// decode json
		let validate_info:ValidateRequest=serde_json::from_slice(&message.payload).map_err(ValidateError::Json)?;

		// download map
		let data=self.roblox_cookie.get_asset(rbx_asset::cookie::GetAssetRequest{
			asset_id:validate_info.model_id,
			version:Some(validate_info.model_version),
		}).await.map_err(ValidateError::Get)?;

		// decode dom (slow!)
		let mut dom=read_dom(&mut std::io::Cursor::new(data)).map_err(ValidateError::ReadDom)?;

		/* VALIDATE MAP */

		// collect unique scripts
		let script_refs=get_script_refs(&dom);
		let mut script_map=std::collections::HashMap::<String,Option<String>>::new();
		for &script_ref in &script_refs{
			if let Some(script)=dom.get_by_ref(script_ref){
				if let Some(rbx_dom_weak::types::Variant::String(source))=script.properties.get("Source"){
					script_map.insert(source.clone(),None);
				}
			}
		}
		// ["local a=true","local b=true"]
		let script_list:Vec<&str>=script_map.keys().map(|s|s.as_str()).collect();
		let j=serde_json::to_string(&script_list).map_err(ValidateError::Json)?;

		// send all scripts to REST endpoint and receive the replacements
		let replacements=self.api.get_replacements(j).await.map_err(ValidateError::ApiGetReplacements)?;

		// assume the iteration order will not change if no keys are inserted or removed
		for (map_value,replacement) in script_map.values_mut().zip(replacements){
			*map_value=replacement;
		}

		// make the replacements
		let mut any_replaced=false;
		for &script_ref in &script_refs{
			if let Some(script)=dom.get_by_ref_mut(script_ref){
				if let Some(rbx_dom_weak::types::Variant::String(source))=script.properties.get_mut("Source"){
					if let Some(Some(replacement))=script_map.get(source.as_str()){
						any_replaced=true;
						*source=replacement.clone();
					}
				}
			}
		}

		// reply with validity
		Ok(if any_replaced{
			// serialize model (slow!)
			let mut data=Vec::new();
			rbx_binary::to_writer(&mut data,&dom,&[dom.root_ref()]).map_err(ValidateError::WriteDom)?;

			// upload a model lol
			let (model_id,model_version)=if let Some(model_id)=validate_info.validated_model_id{
				// upload to existing id
				let response=self.roblox_cookie.upload(rbx_asset::cookie::UploadRequest{
					assetid:model_id,
					name:None,
					description:None,
					ispublic:None,
					allowComments:None,
					groupId:None,
				},data).await.map_err(ValidateError::Upload)?;

				// response.AssetVersionId is not the right one
				let model_version=get_model_version_number();
				(response.AssetId,model_version)
			}else{
				// create new model
				let response=self.roblox_cookie.create(rbx_asset::cookie::CreateRequest{
					name:dom.root().name,
					description:"".to_owned(),
					ispublic:true,
					allowComments:true,
					groupId:None,
				},data).await.map_err(ValidateError::Create)?;

				// response.AssetVersionId is not the right one
				let model_version=get_model_version_number();
				(response.AssetId,model_version)
			};

			// tell the submission validate request to change the model
			Valid::Modified(ModelVersion{
				model_id,
				model_version,
			})
		}else{
			Valid::Untouched
		})
	}
}

#[allow(dead_code)]
#[derive(Debug)]
enum ReadDomError{
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
