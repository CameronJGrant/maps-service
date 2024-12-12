package service

import (
	"context"
	"errors"
	"git.itzana.me/strafesnet/maps-service/pkg/api"
	"git.itzana.me/strafesnet/maps-service/pkg/datastore"
	"github.com/nats-io/nats.go"
)

var (
	// ErrPermissionDenied caller does not have the required role
	ErrPermissionDenied = errors.New("Permission denied")
	// ErrUserInfo user info is missing for some reason
	ErrUserInfo = errors.New("Missing user info")
)

type Service struct {
	DB   datastore.Datastore
	Nats nats.JetStreamContext
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (svc *Service) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	status := 500
	if errors.Is(err, ErrPermissionDenied) {
		status = 403
	}
	if errors.Is(err, ErrUserInfo) {
		status = 401
	}
	return &api.ErrorStatusCode{
		StatusCode: status,
		Response: api.Error{
			Code:    int64(status),
			Message: err.Error(),
		},
	}
}
