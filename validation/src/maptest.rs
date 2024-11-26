use std::collections::HashSet;

type GameID=u32;
type StyleID=u32;
type TimeID=u64;

// the status is supposed to walk up the phases ascending
enum Status{
	// Phase 3: Retired
	Published,// pipeline is complete, map was published
	Rejected,// map council will not request changes, submit a new map
	// Phase 2: Staging
	Publishing,// Admin clicked the publish button, publishing is in progress and should only take a second
	Validated,// Validator validated the map, expose a "Publish" button to Admins
	Accepted,// Reviewer accepts the map, trigger validation once and expose a "Validate" button to Reviewers
	// Phase 1: Maptest
	ChangesRequested,// map council requests changes
	Submitted,// Submitted by owner for review by map council, expose "Accept" / "Reject" / "ChangesRequested" to reviewers
	UnderConstruction,// default state upon pipeline creation
}

enum MaptestType{
	// mapfixes change an existing map on staging game, so they know what map_id they upload to.
	Mapfix{
		//main game target asset id
		asset_id:u64,
		// How is this data entered?  It should be suggested by map maker and finalized by time deleter
		//affected styles to be wiped upon publish
		delete_styles:Option<HashSet<(GameID,StyleID)>>,
		//affected times to be individually removed
		delete_times:Option<HashSet<TimeID>>,
	},
	// map submissions create a new map entry in the staging map database when they are accepted.
	Submission{
		games:HashSet<GameID>,
		creator:String,
		display_name:String,
	},
}

struct Map{
	//maptest database entry id
	id:u64,
	date:u64,
	// int64 UserID who created the submission
	// this user is allowed to change any data in the maptest phase, except for mapfix target map id
	owner:u64,
	model_id:u64,// asset id of the most recently submitted model
	model_version:u64,// snapshot of the model to prevent tampering
	is_completed:bool,// has this asset version been completed by any player
	maptest_type:MaptestType,
	status:Status,
}
