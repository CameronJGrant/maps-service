// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// ActionSubmissionPublish implements actionSubmissionPublish operation.
	//
	// (Internal endpoint) Role Validator changes status from Publishing -> Published.
	//
	// POST /submissions/{SubmissionID}/status/validator-published
	ActionSubmissionPublish(ctx context.Context, params ActionSubmissionPublishParams) error
	// ActionSubmissionReject implements actionSubmissionReject operation.
	//
	// Role Reviewer changes status from Submitted -> Rejected.
	//
	// POST /submissions/{SubmissionID}/status/reject
	ActionSubmissionReject(ctx context.Context, params ActionSubmissionRejectParams) error
	// ActionSubmissionRequestChanges implements actionSubmissionRequestChanges operation.
	//
	// Role Reviewer changes status from Validated|Accepted|Submitted -> ChangesRequested.
	//
	// POST /submissions/{SubmissionID}/status/request-changes
	ActionSubmissionRequestChanges(ctx context.Context, params ActionSubmissionRequestChangesParams) error
	// ActionSubmissionRevoke implements actionSubmissionRevoke operation.
	//
	// Role Submitter changes status from Submitted|ChangesRequested -> UnderConstruction.
	//
	// POST /submissions/{SubmissionID}/status/revoke
	ActionSubmissionRevoke(ctx context.Context, params ActionSubmissionRevokeParams) error
	// ActionSubmissionSubmit implements actionSubmissionSubmit operation.
	//
	// Role Submitter changes status from UnderConstruction|ChangesRequested -> Submitted.
	//
	// POST /submissions/{SubmissionID}/status/submit
	ActionSubmissionSubmit(ctx context.Context, params ActionSubmissionSubmitParams) error
	// ActionSubmissionTriggerPublish implements actionSubmissionTriggerPublish operation.
	//
	// Role Admin changes status from Validated -> Publishing.
	//
	// POST /submissions/{SubmissionID}/status/trigger-publish
	ActionSubmissionTriggerPublish(ctx context.Context, params ActionSubmissionTriggerPublishParams) error
	// ActionSubmissionTriggerValidate implements actionSubmissionTriggerValidate operation.
	//
	// Role Reviewer triggers validation and changes status from Submitted|Accepted -> Validating.
	//
	// POST /submissions/{SubmissionID}/status/trigger-validate
	ActionSubmissionTriggerValidate(ctx context.Context, params ActionSubmissionTriggerValidateParams) error
	// ActionSubmissionValidate implements actionSubmissionValidate operation.
	//
	// (Internal endpoint) Role Validator changes status from Validating -> Validated.
	//
	// POST /submissions/{SubmissionID}/status/validator-validated
	ActionSubmissionValidate(ctx context.Context, params ActionSubmissionValidateParams) error
	// CreateScript implements createScript operation.
	//
	// Create a new script.
	//
	// POST /scripts
	CreateScript(ctx context.Context, req *ScriptCreate) (*ID, error)
	// CreateScriptPolicy implements createScriptPolicy operation.
	//
	// Create a new script policy.
	//
	// POST /script-policy
	CreateScriptPolicy(ctx context.Context, req *ScriptPolicyCreate) (*ID, error)
	// CreateSubmission implements createSubmission operation.
	//
	// Create new submission.
	//
	// POST /submissions
	CreateSubmission(ctx context.Context, req *SubmissionCreate) (*ID, error)
	// DeleteScript implements deleteScript operation.
	//
	// Delete the specified script by ID.
	//
	// DELETE /scripts/{ScriptID}
	DeleteScript(ctx context.Context, params DeleteScriptParams) error
	// DeleteScriptPolicy implements deleteScriptPolicy operation.
	//
	// Delete the specified script policy by ID.
	//
	// DELETE /script-policy/id/{ScriptPolicyID}
	DeleteScriptPolicy(ctx context.Context, params DeleteScriptPolicyParams) error
	// GetScript implements getScript operation.
	//
	// Get the specified script by ID.
	//
	// GET /scripts/{ScriptID}
	GetScript(ctx context.Context, params GetScriptParams) (*Script, error)
	// GetScriptPolicy implements getScriptPolicy operation.
	//
	// Get the specified script policy by ID.
	//
	// GET /script-policy/id/{ScriptPolicyID}
	GetScriptPolicy(ctx context.Context, params GetScriptPolicyParams) (*ScriptPolicy, error)
	// GetScriptPolicyFromHash implements getScriptPolicyFromHash operation.
	//
	// Get the policy for the given hash of script source code.
	//
	// GET /script-policy/hash/{FromScriptHash}
	GetScriptPolicyFromHash(ctx context.Context, params GetScriptPolicyFromHashParams) (*ScriptPolicy, error)
	// GetSubmission implements getSubmission operation.
	//
	// Retrieve map with ID.
	//
	// GET /submissions/{SubmissionID}
	GetSubmission(ctx context.Context, params GetSubmissionParams) (*Submission, error)
	// ListSubmissions implements listSubmissions operation.
	//
	// Get list of submissions.
	//
	// GET /submissions
	ListSubmissions(ctx context.Context, params ListSubmissionsParams) ([]Submission, error)
	// SetSubmissionCompleted implements setSubmissionCompleted operation.
	//
	// Retrieve map with ID.
	//
	// POST /submissions/{SubmissionID}/completed
	SetSubmissionCompleted(ctx context.Context, params SetSubmissionCompletedParams) error
	// UpdateScript implements updateScript operation.
	//
	// Update the specified script by ID.
	//
	// POST /scripts/{ScriptID}
	UpdateScript(ctx context.Context, req *ScriptUpdate, params UpdateScriptParams) error
	// UpdateScriptPolicy implements updateScriptPolicy operation.
	//
	// Update the specified script policy by ID.
	//
	// POST /script-policy/id/{ScriptPolicyID}
	UpdateScriptPolicy(ctx context.Context, req *ScriptPolicyUpdate, params UpdateScriptPolicyParams) error
	// UpdateSubmissionModel implements updateSubmissionModel operation.
	//
	// Update model following role restrictions.
	//
	// POST /submissions/{SubmissionID}/model
	UpdateSubmissionModel(ctx context.Context, params UpdateSubmissionModelParams) error
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
