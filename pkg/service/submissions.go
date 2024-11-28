package service

import (
	"context"
	"git.itzana.me/strafesnet/maps-service/pkg/api"
)

// POST /submissions
func (svc *Service) CreateSubmission(ctx context.Context) (*api.ID, error) {
	return nil, nil
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
