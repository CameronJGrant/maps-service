use crate::types::Games;

// This data structure just exists to remember the hash of the model file as it's uploaded
// so that the hash can be checked later to see if the file is changed
struct AssetVersion{
	//roblox asset id
	asset_id:u64,
	//hash of map file before upload
	hash:u128,
}

struct Git;

struct Map{
	id:u64,//database entry
	date:u64,
	games:Games,
	creator:String,
	display_name:String,
	//the staging model is created at the same time as the maps database entry
	staging:AssetVersion,
	//the main model does not exist before the submission is published for the first time
	main:Option<AssetVersion>,
	changelog:Git,//An entire git repo, ideally the xml of the roblox map
}
