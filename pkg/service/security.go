package service

import (
	"errors"
	"context"
	"git.itzana.me/strafesnet/maps-service/pkg/api"
	"git.itzana.me/strafesnet/go-grpc/auth"
)

var (
	// ErrMissingSessionID there is no session id
	ErrMissingSessionID = errors.New("SessionID missing")
	// ErrInvalidSession caller does not have a valid session
	ErrInvalidSession = errors.New("Session invalid")
)

var (
	// has SubmissionPublish
	RoleMapAdmin   int32 = 128
	// has SubmissionReview
	RoleMapCouncil int32 = 64
)

type Roles struct {
	// human roles
	SubmissionPublish bool
	SubmissionReview bool
	ScriptWrite bool
	// automated roles
	Maptest bool
	Validator bool
}

type UserInfo struct {
	Roles Roles
	UserID int64
}

func (usr UserInfo) IsSubmitter(submitter int64) bool{
	return usr.UserID == submitter
}

type SecurityHandler struct {
	Client auth.AuthServiceClient
}

func (svc SecurityHandler) HandleCookieAuth(ctx context.Context, operationName api.OperationName, t api.CookieAuth) (context.Context, error){
	sessionId := t.GetAPIKey()
	if sessionId == "" {
		return nil, ErrMissingSessionID
	}

	session, err := svc.Client.GetSessionUser(ctx, &auth.IdMessage{
		SessionID: sessionId,
	})
	if err != nil{
		return nil, err
	}

	role, err := svc.Client.GetGroupRole(ctx, &auth.IdMessage{
		SessionID: sessionId,
	})
	if err != nil{
		return nil, err
	}

	validate, err := svc.Client.ValidateSession(ctx, &auth.IdMessage{
		SessionID: sessionId,
	})
	if err != nil{
		return nil, err
	}
	if !validate.Valid{
		return nil, ErrInvalidSession
	}

	roles := Roles{}

	// fix this when roblox udpates group roles
	for _,r := range role.Roles{
		if RoleMapAdmin<=r.Rank{
			roles.SubmissionPublish = true
		}
		if RoleMapCouncil<=r.Rank{
			roles.SubmissionReview = true
		}
	}

	newCtx := context.WithValue(ctx, "UserInfo", UserInfo{
		Roles: roles,
		UserID: int64(session.UserID),
	})

	return newCtx, nil
}
