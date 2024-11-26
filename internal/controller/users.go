package controller

import (
	"context"
	"git.itzana.me/strafesnet/data-service/internal/datastore"
	"git.itzana.me/strafesnet/data-service/internal/model"
	"git.itzana.me/strafesnet/go-grpc/users"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Users struct {
	*users.UnimplementedUsersServiceServer
	Store datastore.Datastore
}

func (u Users) Get(ctx context.Context, request *users.IdMessage) (*users.UserResponse, error) {
	ur, err := u.Store.Users().Get(ctx, request.ID)
	if err != nil {
		if err == datastore.ErrNotExist {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to get user").Error())
	}

	return &users.UserResponse{
		ID:       ur.ID,
		Username: ur.Username,
		StateID:  ur.StateID,
	}, nil
}

func (u Users) GetList(ctx context.Context, list *users.IdList) (*users.UserList, error) {
	uList, err := u.Store.Users().GetList(ctx, list.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to get users").Error())
	}

	var resp users.UserList
	for i := 0; i < len(uList); i++ {
		resp.Users = append(resp.Users, &users.UserResponse{
			ID:       uList[i].ID,
			Username: uList[i].Username,
			StateID:  uList[i].StateID,
		})
	}

	return &resp, nil
}

func (u Users) Update(ctx context.Context, request *users.UserRequest) (*users.NullResponse, error) {
	updates := datastore.Optional()
	updates.AddNotNil("state_id", request.StateID)
	updates.AddNotNil("username", request.Username)

	if err := u.Store.Users().Update(ctx, request.GetID(), updates); err != nil {
		if err == datastore.ErrNotExist {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to update user").Error())
	}

	return &users.NullResponse{}, nil
}

func (u Users) Create(ctx context.Context, request *users.UserRequest) (*users.IdMessage, error) {
	us, err := u.Store.Users().Create(ctx, model.User{
		ID:       request.GetID(),
		Username: request.GetUsername(),
		StateID:  request.GetStateID(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to create user").Error())
	}

	return &users.IdMessage{ID: us.ID}, nil
}

func (u Users) Delete(ctx context.Context, request *users.IdMessage) (*users.NullResponse, error) {
	if err := u.Store.Users().Delete(ctx, request.GetID()); err != nil {
		if err == datastore.ErrNotExist {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to delete user").Error())
	}

	return &users.NullResponse{}, nil
}

func (u Users) List(ctx context.Context, request *users.ListRequest) (*users.UserList, error) {
	filters := datastore.Optional()
	if request.Filter != nil {
		filters.AddNotNil("id", request.GetFilter().ID)
		filters.AddNotNil("state_id", request.GetFilter().StateID)
		filters.AddNotNil("username", request.GetFilter().Username)
	}

	uList, err := u.Store.Users().List(ctx, filters, model.Page{
		Number: request.GetPage().GetNumber(),
		Size:   request.GetPage().GetSize(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to get filtered users").Error())
	}

	var uResp users.UserList
	for i := 0; i < len(uList); i++ {
		uResp.Users = append(uResp.Users, &users.UserResponse{
			ID:       uList[i].ID,
			Username: uList[i].Username,
			StateID:  uList[i].StateID,
		})
	}

	return &uResp, nil
}
