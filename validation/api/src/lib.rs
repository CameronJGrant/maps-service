#[derive(Debug)]
pub enum Error{
	ParseError(url::ParseError),
	Reqwest(reqwest::Error),
}
impl std::fmt::Display for Error{
	fn fmt(&self,f:&mut std::fmt::Formatter<'_>)->std::fmt::Result{
		write!(f,"{self:?}")
	}
}
impl std::error::Error for Error{}

#[derive(serde::Deserialize)]
pub struct ScriptID(i64);

#[allow(nonstandard_style)]
pub struct GetScriptRequest{
	pub ScriptID:ScriptID,
}
#[allow(nonstandard_style)]
#[derive(serde::Deserialize)]
pub struct ScriptResponse{
	pub ID:i64,
	pub Hash:String,
	pub Source:String,
	pub SubmissionID:i64,
}

#[derive(serde::Deserialize)]
#[repr(i32)]
pub enum Policy{
	Allowed=0,
	Blocked=1,
	Delete=2,
	Replace=3,
}

pub struct ScriptPolicyHashRequest{
	pub hash:String,
}
#[allow(nonstandard_style)]
#[derive(serde::Deserialize)]
pub struct ScriptPolicyResponse{
	pub ID:i64,
	pub FromScriptHash:String,
	pub ToScriptID:ScriptID,
	pub Policy:Policy
}

#[allow(nonstandard_style)]
pub struct UpdateSubmissionModelRequest{
	pub ID:i64,
	pub ModelID:u64,
	pub ModelVersion:u64,
}

pub struct SubmissionID(pub i64);

#[derive(Clone)]
pub struct Context{
	base_url:String,
	client:reqwest::Client,
}

pub type ReqwestError=reqwest::Error;

impl Context{
	pub fn new(
		base_url:String,
		// cert:reqwest::Certificate,
		// identity:reqwest::Identity,
	)->reqwest::Result<Self>{
		Ok(Self{
			base_url,
			client:reqwest::Client::builder()
				.use_rustls_tls()
				//.tls_built_in_root_certs(false)
				//.add_root_certificate(cert)
				//.identity(identity)
				.https_only(true)
				.build()?,
		})
	}
	async fn get(&self,url:impl reqwest::IntoUrl)->Result<reqwest::Response,reqwest::Error>{
		self.client.get(url)
		.send().await
	}
	async fn patch(&self,url:impl reqwest::IntoUrl)->Result<reqwest::Response,reqwest::Error>{
		self.client.patch(url)
		.send().await
	}
	pub async fn get_script(&self,config:GetScriptRequest)->Result<ScriptResponse,Error>{
		let url_raw=format!("{}/scripts/{}",self.base_url,config.ScriptID.0);
		let url=reqwest::Url::parse(url_raw.as_str()).map_err(Error::ParseError)?;

		self.get(url).await.map_err(Error::Reqwest)?
		.error_for_status().map_err(Error::Reqwest)?
		.json().await.map_err(Error::Reqwest)
	}
	pub async fn get_script_policy_from_hash(&self,config:ScriptPolicyHashRequest)->Result<ScriptPolicyResponse,Error>{
		let url_raw=format!("{}/script-policy/hash/{}",self.base_url,config.hash);
		let url=reqwest::Url::parse(url_raw.as_str()).map_err(Error::ParseError)?;

		self.get(url).await.map_err(Error::Reqwest)?
		.error_for_status().map_err(Error::Reqwest)?
		.json().await.map_err(Error::Reqwest)
	}
	pub async fn update_submission_model(&self,config:UpdateSubmissionModelRequest)->Result<(),Error>{
		let url_raw=format!("{}/submissions/{}/model",self.base_url,config.ID);
		let mut url=reqwest::Url::parse(url_raw.as_str()).map_err(Error::ParseError)?;

		{
			url.query_pairs_mut()
				.append_pair("ModelID",config.ModelID.to_string().as_str())
				.append_pair("ModelVersion",config.ModelVersion.to_string().as_str());
		}

		self.patch(url).await.map_err(Error::Reqwest)?
		.error_for_status().map_err(Error::Reqwest)?;

		Ok(())
	}
	pub async fn action_submission_validate(&self,config:SubmissionID)->Result<(),Error>{
		let url_raw=format!("{}/submissions/{}/status/validate",self.base_url,config.0);
		let url=reqwest::Url::parse(url_raw.as_str()).map_err(Error::ParseError)?;

		self.patch(url).await.map_err(Error::Reqwest)?
		.error_for_status().map_err(Error::Reqwest)?;

		Ok(())
	}
}
