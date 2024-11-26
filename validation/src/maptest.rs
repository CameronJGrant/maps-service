use crate::types::Games;

enum Status{
	Accepted,
	Rejected,//map council will not request changes, submit a new map
	ChangesRequested,//map council requests changes
	Submitted,//Submitted by owner for review by map council
	UnderConstruction,//default state upon map creation
}

enum MaptestType{
	// mapfixes change an existing map on staging game, so they know what map_id they upload to.
	Mapfix{
		//maps database entry id
		map_id:u64,
	},
	// map submissions create a new map entry in the staging map database when they are accepted.
	Submission{
		games:Games,
		creator:String,
		display_name:String,
	},
}

struct Map{
	//maptest database entry id
	id:u64,
	date:u64,
	// int64 UserID who created the submission
	// this user is allowed to change any data at any time, except for mapfix target map id
	owner:u64,
	model_id:u64,// asset id of the most recently submitted model
	maptest_type:MaptestType,
	status:Status,
}
