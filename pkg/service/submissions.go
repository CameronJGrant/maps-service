package service

import (
	"time"
	"context"
	"git.itzana.me/strafesnet/maps-service/pkg/api"
	"git.itzana.me/strafesnet/maps-service/pkg/model"
	"git.itzana.me/strafesnet/maps-service/pkg/datastore"
)

// POST /submissions
func (svc *Service) CreateSubmission(ctx context.Context, request api.OptSubmissionCreate) (*api.ID, error) {
	//I don't know how to read the http body
	submission, err := svc.DB.Submissions().Create(ctx, model.Submission{
		ID:            0,
		DisplayName:   request.Value.DisplayName.Value,
		Creator:       request.Value.Creator.Value,
		GameID:        request.Value.GameID.Value,
		Date:          time.Now(),
		Submitter:     request.Value.Submitter.Value,
		AssetID:       request.Value.AssetID.Value,
		AssetVersion:  request.Value.AssetVersion.Value,
		Completed:     false,
		TargetAssetID: request.Value.TargetAssetID.Value,
		StatusID:      model.StatusUnderConstruction,
	})
	if err != nil{
		return nil, err
	}
	return &api.ID{
		ID:api.NewOptInt64(submission.ID),
	}, nil
}

// GetSubmission implements getSubmission operation.
//
// Retrieve map with ID.
//
// GET /submissions/{SubmissionID}
func (svc *Service) GetSubmission(ctx context.Context, params api.GetSubmissionParams) (*api.Submission, error) {
	submission, err := svc.DB.Submissions().Get(ctx, params.SubmissionID)
	if err != nil{
		return nil, err
	}
	return &api.Submission{
		ID:             api.NewOptInt64(submission.ID),
		DisplayName:    api.NewOptString(submission.DisplayName),
		Creator:        api.NewOptString(submission.Creator),
		GameID:         api.NewOptInt32(submission.GameID),
		Date:           api.NewOptInt64(submission.Date.Unix()),
		Submitter:      api.NewOptInt64(submission.Submitter),
		AssetID:        api.NewOptInt64(submission.AssetID),
		AssetVersion:   api.NewOptInt64(submission.AssetVersion),
		Completed:      api.NewOptBool(submission.Completed),
		TargetAssetID:  api.NewOptInt64(submission.TargetAssetID),
		StatusID:       api.NewOptInt32(int32(submission.StatusID)),
	}, nil
}

// ListSubmissions implements listSubmissions operation.
//
// Get list of submissions.
//
// GET /submissions
func (svc *Service) ListSubmissions(ctx context.Context, request api.ListSubmissionsParams) ([]api.Submission, error) {
	filter := datastore.Optional()
	//fmt.Println(request)
	if request.Filter.IsSet() {
		filter.AddNotNil("display_name", request.Filter.Value.DisplayName)
		filter.AddNotNil("creator", request.Filter.Value.Creator)
		filter.AddNotNil("game_id", request.Filter.Value.GameID)
	}

	items, err := svc.DB.Submissions().List(ctx, filter, model.Page{
		Number: request.Page.GetPage(),
		Size:   request.Page.GetLimit(),
	})
	if err != nil{
		return nil, err
	}

	var resp []api.Submission
	for i := 0; i < len(items); i++ {
		resp = append(resp, api.Submission{
			ID:             api.NewOptInt64(items[i].ID),
			DisplayName:    api.NewOptString(items[i].DisplayName),
			Creator:        api.NewOptString(items[i].Creator),
			GameID:         api.NewOptInt32(items[i].GameID),
			Date:           api.NewOptInt64(items[i].Date.Unix()),
			Submitter:      api.NewOptInt64(items[i].Submitter),
			AssetID:        api.NewOptInt64(items[i].AssetID),
			AssetVersion:   api.NewOptInt64(items[i].AssetVersion),
			Completed:      api.NewOptBool(items[i].Completed),
			TargetAssetID:  api.NewOptInt64(items[i].TargetAssetID),
			StatusID:       api.NewOptInt32(int32(items[i].StatusID)),
		})
	}

	return resp, nil
}

// PatchSubmissionCompleted implements patchSubmissionCompleted operation.
//
// Retrieve map with ID.
//
// PATCH /submissions/{SubmissionID}/completed
func (svc *Service) PatchSubmissionCompleted(ctx context.Context, params api.PatchSubmissionCompletedParams) error {
	// check if caller has MaptestGame role (request must originate from a maptest roblox game)
	pmap := datastore.Optional()
	pmap.Add("completed", true)
	err := svc.DB.Submissions().Update(ctx, params.SubmissionID, pmap)
	return err
}

// PatchSubmissionModel implements patchSubmissionModel operation.
//
// Update model following role restrictions.
//
// PATCH /submissions/{SubmissionID}/model
func (svc *Service) PatchSubmissionModel(ctx context.Context, params api.PatchSubmissionModelParams) error {
	// check if caller has Submitter role
	// check if Status is ChangesRequested|Submitted|UnderConstruction
	// PROBLEM how to deal with async? data may become out of date
	pmap := datastore.Optional()
	pmap.AddNotNil("asset_id", params.ModelID)
	pmap.AddNotNil("asset_version", params.VersionID)
	//always reset completed when model changes
	pmap.Add("completed",false)
	err := svc.DB.Submissions().Update(ctx, params.SubmissionID, pmap)
	return err
}

// ActionSubmissionPublish invokes actionSubmissionPublish operation.
//
// Role Validator changes status from Publishing -> Published.
//
// PATCH /submissions/{SubmissionID}/status/publish
func (svc *Service) ActionSubmissionPublish(ctx context.Context, params api.ActionSubmissionPublishParams) error {
	return nil
}
// ActionSubmissionReject invokes actionSubmissionReject operation.
//
// Role Reviewer changes status from Submitted -> Rejected.
//
// PATCH /submissions/{SubmissionID}/status/reject
func (svc *Service) ActionSubmissionReject(ctx context.Context, params api.ActionSubmissionRejectParams) error {
	return nil
}
// ActionSubmissionRequestChanges invokes actionSubmissionRequestChanges operation.
//
// Role Reviewer changes status from Validated|Accepted|Submitted -> ChangesRequested.
//
// PATCH /submissions/{SubmissionID}/status/request-changes
func (svc *Service) ActionSubmissionRequestChanges(ctx context.Context, params api.ActionSubmissionRequestChangesParams) error {
	return nil
}
// ActionSubmissionRevoke invokes actionSubmissionRevoke operation.
//
// Role Submitter changes status from Submitted|ChangesRequested -> UnderConstruction.
//
// PATCH /submissions/{SubmissionID}/status/revoke
func (svc *Service) ActionSubmissionRevoke(ctx context.Context, params api.ActionSubmissionRevokeParams) error {
	return nil
}
// ActionSubmissionSubmit invokes actionSubmissionSubmit operation.
//
// Role Submitter changes status from UnderConstruction|ChangesRequested -> Submitted.
//
// PATCH /submissions/{SubmissionID}/status/submit
func (svc *Service) ActionSubmissionSubmit(ctx context.Context, params api.ActionSubmissionSubmitParams) error {
	return nil
}
// ActionSubmissionTriggerPublish invokes actionSubmissionTriggerPublish operation.
//
// Role Admin changes status from Validated -> Publishing.
//
// PATCH /submissions/{SubmissionID}/status/trigger-publish
func (svc *Service) ActionSubmissionTriggerPublish(ctx context.Context, params api.ActionSubmissionTriggerPublishParams) error {
	return nil
}
// ActionSubmissionTriggerValidate invokes actionSubmissionTriggerValidate operation.
//
// Role Reviewer triggers validation and changes status from Submitted|Accepted -> Validating.
//
// PATCH /submissions/{SubmissionID}/status/trigger-validate
func (svc *Service) ActionSubmissionTriggerValidate(ctx context.Context, params api.ActionSubmissionTriggerValidateParams) error {
	return nil
}
// ActionSubmissionValidate invokes actionSubmissionValidate operation.
//
// Role Validator changes status from Validating -> Validated.
//
// PATCH /submissions/{SubmissionID}/status/validate
func (svc *Service) ActionSubmissionValidate(ctx context.Context, params api.ActionSubmissionValidateParams) error {
	return nil
}
