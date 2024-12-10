package service

import (
	"errors"
	"context"
	"git.itzana.me/strafesnet/maps-service/pkg/api"
	"git.itzana.me/strafesnet/maps-service/pkg/model"
	"git.itzana.me/strafesnet/maps-service/pkg/datastore"
)

var (
	// ErrInvalidSourceStatus current submission status cannot change to destination status
	ErrInvalidSourceStatus = errors.New("Invalid source status")
)

// POST /submissions
func (svc *Service) CreateSubmission(ctx context.Context, request *api.SubmissionCreate) (*api.ID, error) {
	userInfo, ok := ctx.Value("UserInfo").(UserInfo)
	if !ok{
		return nil, ErrUserInfo
	}

	submission, err := svc.DB.Submissions().Create(ctx, model.Submission{
		ID:            0,
		DisplayName:   request.DisplayName,
		Creator:       request.Creator,
		GameID:        request.GameID,
		Submitter:     int64(userInfo.UserID),
		AssetID:       request.AssetID,
		AssetVersion:  request.AssetVersion,
		Completed:     false,
		TargetAssetID: request.TargetAssetID.Value,
		StatusID:      model.StatusUnderConstruction,
	})
	if err != nil{
		return nil, err
	}
	return &api.ID{
		ID:submission.ID,
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
		ID:             submission.ID,
		DisplayName:    submission.DisplayName,
		Creator:        submission.Creator,
		GameID:         submission.GameID,
		CreatedAt:      submission.CreatedAt.Unix(),
		UpdatedAt:      submission.UpdatedAt.Unix(),
		Submitter:      submission.Submitter,
		AssetID:        submission.AssetID,
		AssetVersion:   submission.AssetVersion,
		Completed:      submission.Completed,
		TargetAssetID:  api.NewOptInt64(submission.TargetAssetID),
		StatusID:       int32(submission.StatusID),
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
			ID:             items[i].ID,
			DisplayName:    items[i].DisplayName,
			Creator:        items[i].Creator,
			GameID:         items[i].GameID,
			CreatedAt:      items[i].CreatedAt.Unix(),
			UpdatedAt:      items[i].UpdatedAt.Unix(),
			Submitter:      items[i].Submitter,
			AssetID:        items[i].AssetID,
			AssetVersion:   items[i].AssetVersion,
			Completed:      items[i].Completed,
			TargetAssetID:  api.NewOptInt64(items[i].TargetAssetID),
			StatusID:       int32(items[i].StatusID),
		})
	}

	return resp, nil
}

// PatchSubmissionCompleted implements patchSubmissionCompleted operation.
//
// Retrieve map with ID.
//
// POST /submissions/{SubmissionID}/completed
func (svc *Service) SetSubmissionCompleted(ctx context.Context, params api.SetSubmissionCompletedParams) error {
	userInfo, ok := ctx.Value("UserInfo").(UserInfo)
	if !ok{
		return ErrUserInfo
	}

	// check if caller has MaptestGame role (request must originate from a maptest roblox game)
	if !userInfo.Roles.Maptest{
		return ErrPermissionDenied
	}

	pmap := datastore.Optional()
	pmap.Add("completed", true)
	err := svc.DB.Submissions().Update(ctx, params.SubmissionID, pmap)
	return err
}

// PatchSubmissionModel implements patchSubmissionModel operation.
//
// Update model following role restrictions.
//
// POST /submissions/{SubmissionID}/model
func (svc *Service) UpdateSubmissionModel(ctx context.Context, params api.UpdateSubmissionModelParams) error {
	userInfo, ok := ctx.Value("UserInfo").(UserInfo)
	if !ok{
		return ErrUserInfo
	}

	// read submission (this could be done with a transaction WHERE clause)
	submission, err := svc.DB.Submissions().Get(ctx, params.SubmissionID)
	if err != nil{
		return err
	}

	// check if caller is the submitter
	if !userInfo.IsSubmitter(submission.Submitter){
		return ErrPermissionDenied
	}

	// check if Status is ChangesRequested|Submitted|UnderConstruction
	pmap := datastore.Optional()
	pmap.AddNotNil("asset_id", params.ModelID)
	pmap.AddNotNil("asset_version", params.VersionID)
	//always reset completed when model changes
	pmap.Add("completed",false)
	return svc.DB.Submissions().IfStatusThenUpdate(ctx, params.SubmissionID, []model.Status{model.StatusChangesRequested,model.StatusSubmitted,model.StatusUnderConstruction}, pmap)
}

// ActionSubmissionPublish invokes actionSubmissionPublish operation.
//
// Role Validator changes status from Publishing -> Published.
//
// POST /submissions/{SubmissionID}/status/publish
func (svc *Service) ActionSubmissionPublish(ctx context.Context, params api.ActionSubmissionPublishParams) error {
	userInfo, ok := ctx.Value("UserInfo").(UserInfo)
	if !ok{
		return ErrUserInfo
	}

	// check if caller has required role
	if !userInfo.Roles.Validator{
		return ErrPermissionDenied
	}

	// transaction
	smap := datastore.Optional()
	smap.Add("status_id",model.StatusPublished)
	return svc.DB.Submissions().IfStatusThenUpdate(ctx, params.SubmissionID, []model.Status{model.StatusPublishing}, smap)
}
// ActionSubmissionReject invokes actionSubmissionReject operation.
//
// Role Reviewer changes status from Submitted -> Rejected.
//
// POST /submissions/{SubmissionID}/status/reject
func (svc *Service) ActionSubmissionReject(ctx context.Context, params api.ActionSubmissionRejectParams) error {
	userInfo, ok := ctx.Value("UserInfo").(UserInfo)
	if !ok{
		return ErrUserInfo
	}

	// check if caller has required role
	if !userInfo.Roles.SubmissionReview{
		return ErrPermissionDenied
	}

	// transaction
	smap := datastore.Optional()
	smap.Add("status_id",model.StatusRejected)
	return svc.DB.Submissions().IfStatusThenUpdate(ctx, params.SubmissionID, []model.Status{model.StatusSubmitted}, smap)
}
// ActionSubmissionRequestChanges invokes actionSubmissionRequestChanges operation.
//
// Role Reviewer changes status from Validated|Accepted|Submitted -> ChangesRequested.
//
// POST /submissions/{SubmissionID}/status/request-changes
func (svc *Service) ActionSubmissionRequestChanges(ctx context.Context, params api.ActionSubmissionRequestChangesParams) error {
	userInfo, ok := ctx.Value("UserInfo").(UserInfo)
	if !ok{
		return ErrUserInfo
	}

	// check if caller has required role
	if !userInfo.Roles.SubmissionReview{
		return ErrPermissionDenied
	}

	// transaction
	smap := datastore.Optional()
	smap.Add("status_id",model.StatusChangesRequested)
	return svc.DB.Submissions().IfStatusThenUpdate(ctx, params.SubmissionID, []model.Status{model.StatusValidated,model.StatusAccepted,model.StatusSubmitted}, smap)
}
// ActionSubmissionRevoke invokes actionSubmissionRevoke operation.
//
// Role Submitter changes status from Submitted|ChangesRequested -> UnderConstruction.
//
// POST /submissions/{SubmissionID}/status/revoke
func (svc *Service) ActionSubmissionRevoke(ctx context.Context, params api.ActionSubmissionRevokeParams) error {
	userInfo, ok := ctx.Value("UserInfo").(UserInfo)
	if !ok{
		return ErrUserInfo
	}

	// read submission (this could be done with a transaction WHERE clause)
	submission, err := svc.DB.Submissions().Get(ctx, params.SubmissionID)
	if err != nil{
		return err
	}

	// check if caller is the submitter
	if !userInfo.IsSubmitter(submission.Submitter){
		return ErrPermissionDenied
	}

	// transaction
	smap := datastore.Optional()
	smap.Add("status_id",model.StatusUnderConstruction)
	return svc.DB.Submissions().IfStatusThenUpdate(ctx, params.SubmissionID, []model.Status{model.StatusSubmitted,model.StatusChangesRequested}, smap)
}
// ActionSubmissionSubmit invokes actionSubmissionSubmit operation.
//
// Role Submitter changes status from UnderConstruction|ChangesRequested -> Submitted.
//
// POST /submissions/{SubmissionID}/status/submit
func (svc *Service) ActionSubmissionSubmit(ctx context.Context, params api.ActionSubmissionSubmitParams) error {
	userInfo, ok := ctx.Value("UserInfo").(UserInfo)
	if !ok{
		return ErrUserInfo
	}

	// read submission (this could be done with a transaction WHERE clause)
	submission, err := svc.DB.Submissions().Get(ctx, params.SubmissionID)
	if err != nil{
		return err
	}

	// check if caller is the submitter
	if !userInfo.IsSubmitter(submission.Submitter){
		return ErrPermissionDenied
	}

	// transaction
	smap := datastore.Optional()
	smap.Add("status_id",model.StatusSubmitted)
	return svc.DB.Submissions().IfStatusThenUpdate(ctx, params.SubmissionID, []model.Status{model.StatusUnderConstruction,model.StatusChangesRequested}, smap)
}
// ActionSubmissionTriggerPublish invokes actionSubmissionTriggerPublish operation.
//
// Role Admin changes status from Validated -> Publishing.
//
// POST /submissions/{SubmissionID}/status/trigger-publish
func (svc *Service) ActionSubmissionTriggerPublish(ctx context.Context, params api.ActionSubmissionTriggerPublishParams) error {
	userInfo, ok := ctx.Value("UserInfo").(UserInfo)
	if !ok{
		return ErrUserInfo
	}

	// check if caller has required role
	if !userInfo.Roles.SubmissionPublish{
		return ErrPermissionDenied
	}

	// transaction
	smap := datastore.Optional()
	smap.Add("status_id",model.StatusPublishing)
	return svc.DB.Submissions().IfStatusThenUpdate(ctx, params.SubmissionID, []model.Status{model.StatusValidated}, smap)
}
// ActionSubmissionTriggerValidate invokes actionSubmissionTriggerValidate operation.
//
// Role Reviewer triggers validation and changes status from Submitted|Accepted -> Validating.
//
// POST /submissions/{SubmissionID}/status/trigger-validate
func (svc *Service) ActionSubmissionTriggerValidate(ctx context.Context, params api.ActionSubmissionTriggerValidateParams) error {
	userInfo, ok := ctx.Value("UserInfo").(UserInfo)
	if !ok{
		return ErrUserInfo
	}

	// check if caller has required role
	if !userInfo.Roles.SubmissionReview{
		return ErrPermissionDenied
	}

	// transaction
	smap := datastore.Optional()
	smap.Add("status_id",model.StatusValidating)
	return svc.DB.Submissions().IfStatusThenUpdate(ctx, params.SubmissionID, []model.Status{model.StatusSubmitted,model.StatusAccepted}, smap)
}
// ActionSubmissionValidate invokes actionSubmissionValidate operation.
//
// Role Validator changes status from Validating -> Validated.
//
// POST /submissions/{SubmissionID}/status/validate
func (svc *Service) ActionSubmissionValidate(ctx context.Context, params api.ActionSubmissionValidateParams) error {
	userInfo, ok := ctx.Value("UserInfo").(UserInfo)
	if !ok{
		return ErrUserInfo
	}

	// check if caller has required role
	if !userInfo.Roles.Validator{
		return ErrPermissionDenied
	}

	// transaction
	smap := datastore.Optional()
	smap.Add("status_id",model.StatusValidated)
	return svc.DB.Submissions().IfStatusThenUpdate(ctx, params.SubmissionID, []model.Status{model.StatusValidating}, smap)
}
