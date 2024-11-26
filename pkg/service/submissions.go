package controller

import (
	"context"
	"git.itzana.me/strafesnet/maps-service/pkg/api"
)

// POST /submissions
func CreateSubmission(ctx context.Context) (*api.Submission, error) {
	return nil, nil
}

// GetSubmission implements getSubmission operation.
//
// Retrieve map with ID.
//
// GET /submissions/{SubmissionID}
func GetSubmission(ctx context.Context, params api.GetSubmissionParams) (*api.Submission, error) {
	return nil, nil
}

// ListSubmissions implements listSubmissions operation.
//
// Get list of submissions.
//
// GET /submissions
func ListSubmissions(ctx context.Context, params api.ListSubmissionsParams) ([]api.Submission, error) {
	return nil, nil
}

// PatchSubmissionCompleted implements patchSubmissionCompleted operation.
//
// Retrieve map with ID.
//
// PATCH /submissions/{SubmissionID}/completed
func PatchSubmissionCompleted(ctx context.Context, params api.PatchSubmissionCompletedParams) error {
	return nil
}

// PatchSubmissionModel implements patchSubmissionModel operation.
//
// Update model following role restrictions.
//
// PATCH /submissions/{SubmissionID}/model
func PatchSubmissionModel(ctx context.Context, params api.PatchSubmissionModelParams) error {
	return nil
}

// PatchSubmissionStatus implements patchSubmissionStatus operation.
//
// Update status following role restrictions.
//
// PATCH /submissions/{SubmissionID}/status
func PatchSubmissionStatus(ctx context.Context, params api.PatchSubmissionStatusParams) error {
	return nil
}
