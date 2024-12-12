// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// ActionSubmissionPublish implements actionSubmissionPublish operation.
//
// (Internal endpoint) Role Validator changes status from Publishing -> Published.
//
// POST /submissions/{SubmissionID}/status/validator-published
func (UnimplementedHandler) ActionSubmissionPublish(ctx context.Context, params ActionSubmissionPublishParams) error {
	return ht.ErrNotImplemented
}

// ActionSubmissionReject implements actionSubmissionReject operation.
//
// Role Reviewer changes status from Submitted -> Rejected.
//
// POST /submissions/{SubmissionID}/status/reject
func (UnimplementedHandler) ActionSubmissionReject(ctx context.Context, params ActionSubmissionRejectParams) error {
	return ht.ErrNotImplemented
}

// ActionSubmissionRequestChanges implements actionSubmissionRequestChanges operation.
//
// Role Reviewer changes status from Validated|Accepted|Submitted -> ChangesRequested.
//
// POST /submissions/{SubmissionID}/status/request-changes
func (UnimplementedHandler) ActionSubmissionRequestChanges(ctx context.Context, params ActionSubmissionRequestChangesParams) error {
	return ht.ErrNotImplemented
}

// ActionSubmissionRevoke implements actionSubmissionRevoke operation.
//
// Role Submitter changes status from Submitted|ChangesRequested -> UnderConstruction.
//
// POST /submissions/{SubmissionID}/status/revoke
func (UnimplementedHandler) ActionSubmissionRevoke(ctx context.Context, params ActionSubmissionRevokeParams) error {
	return ht.ErrNotImplemented
}

// ActionSubmissionSubmit implements actionSubmissionSubmit operation.
//
// Role Submitter changes status from UnderConstruction|ChangesRequested -> Submitted.
//
// POST /submissions/{SubmissionID}/status/submit
func (UnimplementedHandler) ActionSubmissionSubmit(ctx context.Context, params ActionSubmissionSubmitParams) error {
	return ht.ErrNotImplemented
}

// ActionSubmissionTriggerPublish implements actionSubmissionTriggerPublish operation.
//
// Role Admin changes status from Validated -> Publishing.
//
// POST /submissions/{SubmissionID}/status/trigger-publish
func (UnimplementedHandler) ActionSubmissionTriggerPublish(ctx context.Context, params ActionSubmissionTriggerPublishParams) error {
	return ht.ErrNotImplemented
}

// ActionSubmissionTriggerValidate implements actionSubmissionTriggerValidate operation.
//
// Role Reviewer triggers validation and changes status from Submitted|Accepted -> Validating.
//
// POST /submissions/{SubmissionID}/status/trigger-validate
func (UnimplementedHandler) ActionSubmissionTriggerValidate(ctx context.Context, params ActionSubmissionTriggerValidateParams) error {
	return ht.ErrNotImplemented
}

// ActionSubmissionValidate implements actionSubmissionValidate operation.
//
// (Internal endpoint) Role Validator changes status from Validating -> Validated.
//
// POST /submissions/{SubmissionID}/status/validator-validated
func (UnimplementedHandler) ActionSubmissionValidate(ctx context.Context, params ActionSubmissionValidateParams) error {
	return ht.ErrNotImplemented
}

// CreateScript implements createScript operation.
//
// Create a new script.
//
// POST /scripts
func (UnimplementedHandler) CreateScript(ctx context.Context, req *ScriptCreate) (r *ID, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateScriptPolicy implements createScriptPolicy operation.
//
// Create a new script policy.
//
// POST /script-policy
func (UnimplementedHandler) CreateScriptPolicy(ctx context.Context, req *ScriptPolicyCreate) (r *ID, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateSubmission implements createSubmission operation.
//
// Create new submission.
//
// POST /submissions
func (UnimplementedHandler) CreateSubmission(ctx context.Context, req *SubmissionCreate) (r *ID, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteScript implements deleteScript operation.
//
// Delete the specified script by ID.
//
// DELETE /scripts/{ScriptID}
func (UnimplementedHandler) DeleteScript(ctx context.Context, params DeleteScriptParams) error {
	return ht.ErrNotImplemented
}

// DeleteScriptPolicy implements deleteScriptPolicy operation.
//
// Delete the specified script policy by ID.
//
// DELETE /script-policy/id/{ScriptPolicyID}
func (UnimplementedHandler) DeleteScriptPolicy(ctx context.Context, params DeleteScriptPolicyParams) error {
	return ht.ErrNotImplemented
}

// GetScript implements getScript operation.
//
// Get the specified script by ID.
//
// GET /scripts/{ScriptID}
func (UnimplementedHandler) GetScript(ctx context.Context, params GetScriptParams) (r *Script, _ error) {
	return r, ht.ErrNotImplemented
}

// GetScriptPolicy implements getScriptPolicy operation.
//
// Get the specified script policy by ID.
//
// GET /script-policy/id/{ScriptPolicyID}
func (UnimplementedHandler) GetScriptPolicy(ctx context.Context, params GetScriptPolicyParams) (r *ScriptPolicy, _ error) {
	return r, ht.ErrNotImplemented
}

// GetScriptPolicyFromHash implements getScriptPolicyFromHash operation.
//
// Get the policy for the given hash of script source code.
//
// GET /script-policy/hash/{FromScriptHash}
func (UnimplementedHandler) GetScriptPolicyFromHash(ctx context.Context, params GetScriptPolicyFromHashParams) (r *ScriptPolicy, _ error) {
	return r, ht.ErrNotImplemented
}

// GetSubmission implements getSubmission operation.
//
// Retrieve map with ID.
//
// GET /submissions/{SubmissionID}
func (UnimplementedHandler) GetSubmission(ctx context.Context, params GetSubmissionParams) (r *Submission, _ error) {
	return r, ht.ErrNotImplemented
}

// ListSubmissions implements listSubmissions operation.
//
// Get list of submissions.
//
// GET /submissions
func (UnimplementedHandler) ListSubmissions(ctx context.Context, params ListSubmissionsParams) (r []Submission, _ error) {
	return r, ht.ErrNotImplemented
}

// SetSubmissionCompleted implements setSubmissionCompleted operation.
//
// Retrieve map with ID.
//
// POST /submissions/{SubmissionID}/completed
func (UnimplementedHandler) SetSubmissionCompleted(ctx context.Context, params SetSubmissionCompletedParams) error {
	return ht.ErrNotImplemented
}

// UpdateScript implements updateScript operation.
//
// Update the specified script by ID.
//
// POST /scripts/{ScriptID}
func (UnimplementedHandler) UpdateScript(ctx context.Context, req *ScriptUpdate, params UpdateScriptParams) error {
	return ht.ErrNotImplemented
}

// UpdateScriptPolicy implements updateScriptPolicy operation.
//
// Update the specified script policy by ID.
//
// POST /script-policy/id/{ScriptPolicyID}
func (UnimplementedHandler) UpdateScriptPolicy(ctx context.Context, req *ScriptPolicyUpdate, params UpdateScriptPolicyParams) error {
	return ht.ErrNotImplemented
}

// UpdateSubmissionModel implements updateSubmissionModel operation.
//
// Update model following role restrictions.
//
// POST /submissions/{SubmissionID}/model
func (UnimplementedHandler) UpdateSubmissionModel(ctx context.Context, params UpdateSubmissionModelParams) error {
	return ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
