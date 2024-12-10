export const enum SubmissionStatus {
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

export interface SubmissionInfo {
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