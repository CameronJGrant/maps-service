package service

import (
	"context"
	"git.itzana.me/strafesnet/maps-service/pkg/api"
)

var (
)

// CreateScriptPolicy implements createScriptPolicy operation.
//
// Create a new script policy.
//
// POST /script-policy
func (svc *Service) CreateScriptPolicy(ctx context.Context, req api.OptScriptPolicyCreate) (*api.ID, error){
	return nil,nil
}
// DeleteScriptPolicy implements deleteScriptPolicy operation.
//
// Delete the specified script policy by ID.
//
// DELETE /script-policy/id/{ScriptPolicyID}
func (svc *Service) DeleteScriptPolicy(ctx context.Context, params api.DeleteScriptPolicyParams) error{
	return nil
}
// GetScriptPolicy implements getScriptPolicy operation.
//
// Get the specified script policy by ID.
//
// GET /script-policy/id/{ScriptPolicyID}
func (svc *Service) GetScriptPolicy(ctx context.Context, params api.GetScriptPolicyParams) (*api.ScriptPolicy, error){
	return nil,nil
}
// GetScriptPolicyFromHash implements getScriptPolicyFromHash operation.
//
// Get the policy for the given hash of script source code.
//
// GET /script-policy/hash/{FromScriptHash}
func (svc *Service) GetScriptPolicyFromHash(ctx context.Context, params api.GetScriptPolicyFromHashParams) (*api.ScriptPolicy, error){
	return nil,nil
}
// UpdateScriptPolicy implements updateScriptPolicy operation.
//
// Update the specified script policy by ID.
//
// PATCH /script-policy/id/{ScriptPolicyID}
func (svc *Service) UpdateScriptPolicy(ctx context.Context, req api.OptScriptPolicyUpdate, params api.UpdateScriptPolicyParams) error{
	return nil
}
