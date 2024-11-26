package service

import (
	"context"
	"git.itzana.me/strafesnet/maps-service/pkg/api"
	"git.itzana.me/strafesnet/maps-service/pkg/datastore"
)

type Service struct {
	DB datastore.Datastore
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (svc *Service) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	return &api.ErrorStatusCode{
		StatusCode: 500,
		Response:   api.Error{Message: err.Error()},
	}
}
