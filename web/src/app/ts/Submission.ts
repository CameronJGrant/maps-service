const enum SubmissionStatus {
    Published,
    Rejected,
    Publishing,
    Validated,
    Validating,
    Accepted,
    ChangesRequested,
    Submitted,
    UnderConstruction
}

interface SubmissionInfo {
    readonly ID:            number,
    readonly DisplayName:   string,
    readonly Creator:       string,
    readonly GameID:        number,
    readonly Date:          number,
    readonly Submitter:     number,
    readonly AssetID:       number,
    readonly AssetVersion:  number,
    readonly Completed:     boolean,
    readonly TargetAssetID: number,
    readonly StatusID:      SubmissionStatus
}

function SubmissionStatusToString(submission_status: SubmissionStatus): string {
	let Review: string
	switch (submission_status) {
    	case SubmissionStatus.Published:
     		Review = "PUBLISHED"
       		break
     	case SubmissionStatus.Rejected:
      		Review = "REJECTED"
        	break
		case SubmissionStatus.Publishing:
			Review = "PUBLISHING"
			break
		case SubmissionStatus.Validated:
			Review = "VALIDATED"
			break
		case SubmissionStatus.Validating:
			Review = "VALIDATING"
			break
		case SubmissionStatus.Accepted:
			Review = "ACCEPTED"
			break
		case SubmissionStatus.ChangesRequested:
			Review = "CHANGES REQUESTED"
			break
		case SubmissionStatus.Submitted:
			Review = "SUBMITTED"
			break
		case SubmissionStatus.UnderConstruction:
			Review = "UNDER CONSTRUCTION"
			break
		default:
			Review = "UNKNOWN"
    }
    return Review
}

export {
	SubmissionStatus,
	SubmissionStatusToString,
	type SubmissionInfo
}