package service

import (
	"time"
	"context"
	"git.itzana.me/strafesnet/maps-service/pkg/api"
	"git.itzana.me/strafesnet/maps-service/pkg/model"
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
		StatusID:      0,
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
		StatusID:       api.NewOptInt32(submission.StatusID),
	}, nil
}

// ListSubmissions implements listSubmissions operation.
//
// Get list of submissions.
//
// GET /submissions
func (svc *Service) ListSubmissions(ctx context.Context, params api.ListSubmissionsParams) ([]api.Submission, error) {
	return nil, nil
}

// PatchSubmissionCompleted implements patchSubmissionCompleted operation.
//
// Retrieve map with ID.
//
// PATCH /submissions/{SubmissionID}/completed
func (svc *Service) PatchSubmissionCompleted(ctx context.Context, params api.PatchSubmissionCompletedParams) error {
	return nil
}

// PatchSubmissionModel implements patchSubmissionModel operation.
//
// Update model following role restrictions.
//
// PATCH /submissions/{SubmissionID}/model
func (svc *Service) PatchSubmissionModel(ctx context.Context, params api.PatchSubmissionModelParams) error {
	return nil
}

// PatchSubmissionStatus implements patchSubmissionStatus operation.
//
// Update status following role restrictions.
//
// PATCH /submissions/{SubmissionID}/status
func (svc *Service) PatchSubmissionStatus(ctx context.Context, params api.PatchSubmissionStatusParams) error {
	return nil
}
