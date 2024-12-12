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
	switch (submission_status) {
    	case SubmissionStatus.Published:
     		return "PUBLISHED"
     	case SubmissionStatus.Rejected:
      		return "REJECTED"
		case SubmissionStatus.Publishing:
			return "PUBLISHING"
		case SubmissionStatus.Validated:
			return "VALIDATED"
		case SubmissionStatus.Validating:
			return "VALIDATING"
		case SubmissionStatus.Accepted:
			return "ACCEPTED"
		case SubmissionStatus.ChangesRequested:
			return "CHANGES REQUESTED"
		case SubmissionStatus.Submitted:
			return "SUBMITTED"
		case SubmissionStatus.UnderConstruction:
			return "UNDER CONSTRUCTION"
		default:
			return "UNKNOWN"
    }
}

export {
	SubmissionStatus,
	SubmissionStatusToString,
	type SubmissionInfo
}