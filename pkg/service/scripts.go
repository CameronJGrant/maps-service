package service

import (
	"context"
	"git.itzana.me/strafesnet/maps-service/pkg/api"
)

var (
)

// CreateScript implements createScript operation.
//
// Create a new script.
//
// POST /scripts
func (svc *Service) CreateScript(ctx context.Context, req api.OptScriptCreate) (*api.ID, error){
	return nil,nil
}
// DeleteScript implements deleteScript operation.
//
// Delete the specified script by ID.
//
// DELETE /scripts/{ScriptID}
func (svc *Service) DeleteScript(ctx context.Context, params api.DeleteScriptParams) error{
	return nil
}
// GetScript implements getScript operation.
//
// Get the specified script by ID.
//
// GET /scripts/{ScriptID}
func (svc *Service) GetScript(ctx context.Context, params api.GetScriptParams) (*api.Script, error){
	return nil,nil
}
// UpdateScript implements updateScript operation.
//
// Update the specified script by ID.
//
// PATCH /scripts/{ScriptID}
func (svc *Service) UpdateScript(ctx context.Context, req api.OptScriptUpdate, params api.UpdateScriptParams) error{
	return nil
}
