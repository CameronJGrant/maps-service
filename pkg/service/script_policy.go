package service

import (
	"context"
	"fmt"
	"strconv"

	"git.itzana.me/strafesnet/maps-service/pkg/api"
	"git.itzana.me/strafesnet/maps-service/pkg/datastore"
	"git.itzana.me/strafesnet/maps-service/pkg/model"
)

// CreateScriptPolicy implements createScriptPolicy operation.
//
// Create a new script policy.
//
// POST /script-policy
func (svc *Service) CreateScriptPolicy(ctx context.Context, req *api.ScriptPolicyCreate) (*api.ID, error){
	userInfo, ok := ctx.Value("UserInfo").(*UserInfo)
	if !ok{
		return nil, ErrUserInfo
	}

	if !userInfo.Roles.ScriptWrite{
		return nil, ErrPermissionDenied
	}

	from_script, err := svc.DB.Scripts().Get(ctx,req.FromScriptID)
	if err != nil{
		return nil, err
	}

	// the existence of ToScriptID does not need to be validated because it's checked by a foreign key constraint.

	script, err := svc.DB.ScriptPolicy().Create(ctx, model.ScriptPolicy{
		ID:             0,
		FromScriptHash: from_script.Hash,
		ToScriptID:     req.ToScriptID,
		Policy:         model.Policy(req.Policy),
	})
	if err != nil{
		return nil, err
	}

	return &api.ID{
		ID:script.ID,
	}, nil
}
// DeleteScriptPolicy implements deleteScriptPolicy operation.
//
// Delete the specified script policy by ID.
//
// DELETE /script-policy/id/{ScriptPolicyID}
func (svc *Service) DeleteScriptPolicy(ctx context.Context, params api.DeleteScriptPolicyParams) error{
	userInfo, ok := ctx.Value("UserInfo").(*UserInfo)
	if !ok{
		return ErrUserInfo
	}

	if !userInfo.Roles.ScriptWrite{
		return ErrPermissionDenied
	}

	return svc.DB.ScriptPolicy().Delete(ctx, params.ScriptPolicyID)
}
// GetScriptPolicy implements getScriptPolicy operation.
//
// Get the specified script policy by ID.
//
// GET /script-policy/id/{ScriptPolicyID}
func (svc *Service) GetScriptPolicy(ctx context.Context, params api.GetScriptPolicyParams) (*api.ScriptPolicy, error){
	_, ok := ctx.Value("UserInfo").(*UserInfo)
	if !ok{
		return nil, ErrUserInfo
	}

	// Read permission for script policy only requires you to be logged in

	policy, err := svc.DB.ScriptPolicy().Get(ctx, params.ScriptPolicyID)
	if err != nil{
		return nil, err
	}

	return &api.ScriptPolicy{
		ID:             policy.ID,
		FromScriptHash: fmt.Sprintf("%x",policy.FromScriptHash),
		ToScriptID:     policy.ToScriptID,
		Policy:         int32(policy.Policy),
	}, nil
}
// GetScriptPolicyFromHash implements getScriptPolicyFromHash operation.
//
// Get the policy for the given hash of script source code.
//
// GET /script-policy/hash/{FromScriptHash}
func (svc *Service) GetScriptPolicyFromHash(ctx context.Context, params api.GetScriptPolicyFromHashParams) (*api.ScriptPolicy, error){
	_, ok := ctx.Value("UserInfo").(*UserInfo)
	if !ok{
		return nil, ErrUserInfo
	}

	// Read permission for script policy only requires you to be logged in

	// parse hash from hex
	hash, err := strconv.ParseUint(params.FromScriptHash, 16, 64)
	if err != nil{
		return nil, err
	}

	policy, err := svc.DB.ScriptPolicy().GetFromHash(ctx, hash)
	if err != nil{
		return nil, err
	}

	return &api.ScriptPolicy{
		ID:             policy.ID,
		FromScriptHash: fmt.Sprintf("%x",policy.FromScriptHash),
		ToScriptID:     policy.ToScriptID,
		Policy:         int32(policy.Policy),
	}, nil
}
// UpdateScriptPolicy implements updateScriptPolicy operation.
//
// Update the specified script policy by ID.
//
// PATCH /script-policy/id/{ScriptPolicyID}
func (svc *Service) UpdateScriptPolicy(ctx context.Context, req *api.ScriptPolicyUpdate, params api.UpdateScriptPolicyParams) error{
	userInfo, ok := ctx.Value("UserInfo").(*UserInfo)
	if !ok{
		return ErrUserInfo
	}

	if !userInfo.Roles.ScriptWrite{
		return ErrPermissionDenied
	}

	pmap := datastore.Optional()
	if from_script_id,ok:=req.FromScriptID.Get();ok{
		from_script, err := svc.DB.Scripts().Get(ctx,from_script_id)
		if err != nil{
			return err
		}
		pmap.Add("from_script_hash",from_script.Hash)
	}
	if to_script_id,ok:=req.ToScriptID.Get();ok{
		pmap.Add("to_script_id",to_script_id)
	}
	if policy,ok:=req.Policy.Get();ok{
		pmap.Add("policy",policy)
	}
	return svc.DB.ScriptPolicy().Update(ctx, req.ID, pmap)
}
