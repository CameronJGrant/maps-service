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
		StatusID:      model.UnderConstruction,
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

// PatchSubmissionStatus implements patchSubmissionStatus operation.
//
// Update status following role restrictions.
//
// PATCH /submissions/{SubmissionID}/status
func (svc *Service) PatchSubmissionStatus(ctx context.Context, params api.PatchSubmissionStatusParams) error {
	pmap := datastore.Optional()
	pmap.AddNotNil("status", params.Status)
	err := svc.DB.Submissions().Update(ctx, params.SubmissionID, pmap)
	return err
}
