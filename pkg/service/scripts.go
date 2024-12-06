package service

import (
	"context"
	"fmt"

	"git.itzana.me/strafesnet/maps-service/pkg/api"
	"git.itzana.me/strafesnet/maps-service/pkg/datastore"
	"git.itzana.me/strafesnet/maps-service/pkg/model"
)

var (
)

// CreateScript implements createScript operation.
//
// Create a new script.
//
// POST /scripts
func (svc *Service) CreateScript(ctx context.Context, req *api.ScriptCreate) (*api.ID, error){
	userInfo, ok := ctx.Value("UserInfo").(*UserInfo)
	if !ok{
		return nil, ErrUserInfo
	}

	if !userInfo.Roles.ScriptWrite{
		return nil, ErrPermissionDenied
	}

	script, err := svc.DB.Scripts().Create(ctx, model.Script{
		ID:           0,
		Hash:         0, // TODO
		Source:       req.Source,
		SubmissionID: req.SubmissionID.Or(0),
	})
	if err != nil{
		return nil, err
	}

	return &api.ID{
		ID:script.ID,
	}, nil
}
// DeleteScript implements deleteScript operation.
//
// Delete the specified script by ID.
//
// DELETE /scripts/{ScriptID}
func (svc *Service) DeleteScript(ctx context.Context, params api.DeleteScriptParams) error{
	userInfo, ok := ctx.Value("UserInfo").(*UserInfo)
	if !ok{
		return ErrUserInfo
	}

	if !userInfo.Roles.ScriptWrite{
		return ErrPermissionDenied
	}

	return svc.DB.Scripts().Delete(ctx, params.ScriptID)
}
// GetScript implements getScript operation.
//
// Get the specified script by ID.
//
// GET /scripts/{ScriptID}
func (svc *Service) GetScript(ctx context.Context, params api.GetScriptParams) (*api.Script, error){
	_, ok := ctx.Value("UserInfo").(*UserInfo)
	if !ok{
		return nil, ErrUserInfo
	}

	// Read permission for scripts only requires you to be logged in

	script, err := svc.DB.Scripts().Get(ctx, params.ScriptID)
	if err != nil{
		return nil, err
	}

	return &api.Script{
		ID:           script.ID,
		Hash:         fmt.Sprintf("%x",script.Hash),
		Source:       script.Source,
		SubmissionID: script.SubmissionID,
	}, nil
}
// UpdateScript implements updateScript operation.
//
// Update the specified script by ID.
//
// PATCH /scripts/{ScriptID}
func (svc *Service) UpdateScript(ctx context.Context, req *api.ScriptUpdate, params api.UpdateScriptParams) error{
	userInfo, ok := ctx.Value("UserInfo").(*UserInfo)
	if !ok{
		return ErrUserInfo
	}

	if !userInfo.Roles.ScriptWrite{
		return ErrPermissionDenied
	}

	pmap := datastore.Optional()
	if source,ok:=req.Source.Get();ok{
		pmap.Add("source",source)
		panic("unimplemented")
		pmap.Add("from_script_hash",0)
	}
	if SubmissionID,ok:=req.SubmissionID.Get();ok{
		pmap.Add("submission_id",SubmissionID)
	}
	return svc.DB.Scripts().Update(ctx, req.ID, pmap)
}
