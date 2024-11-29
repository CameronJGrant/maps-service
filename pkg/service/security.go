package service

import (
	"errors"
	"context"
	"git.itzana.me/strafesnet/maps-service/pkg/api"
	"git.itzana.me/strafesnet/go-grpc/auth"
	"google.golang.org/grpc"
)

var (
	// ErrMissingSessionID there is no session id
	ErrMissingSessionID = errors.New("SessionID missing")
	// ErrInvalidSession caller does not have a valid session
	ErrInvalidSession = errors.New("Session invalid")
)

var (
	RoleAdmin = 128
	RoleReviewer = 64
)

type Roles struct {
	Admin bool
	Reviewer bool
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
	Client *grpc.ClientConn
}

func (svc SecurityHandler) HandleCookieAuth(ctx context.Context, operationName api.OperationName, t api.CookieAuth) (context.Context, error){
	sessionId := t.GetAPIKey()
	if sessionId == "" {
		return nil, ErrMissingSessionID
	}

	client := auth.NewAuthServiceClient(svc.Client)

	session, err := client.GetSessionUser(ctx, &auth.IdMessage{
		SessionID: sessionId,
	})
	if err != nil{
		return nil, err
	}

	role, err := client.GetGroupRole(ctx, &auth.IdMessage{
		SessionID: sessionId,
	})
	if err != nil{
		return nil, err
	}

	validate, err := client.ValidateSession(ctx, &auth.IdMessage{
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
	for r := range role.Roles{
		if RoleAdmin<=r{
			roles.Admin = true
		}
		if RoleReviewer<=r{
			roles.Reviewer = true
		}
	}

	newCtx := context.WithValue(ctx, "UserInfo", UserInfo{
		Roles: roles,
		UserID: int64(session.UserID),
	})

	return newCtx, nil
}
