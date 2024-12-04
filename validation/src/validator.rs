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

		// validate map
		// validate(dom)

		// reply with validity
		Ok(Valid::Untouched)
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
