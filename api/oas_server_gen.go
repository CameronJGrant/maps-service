// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// GetUser implements getUser operation.
	//
	// Retrieve user with ID.
	//
	// GET /users/{UserID}
	GetUser(ctx context.Context, params GetUserParams) (*User, error)
	// GetUserRank implements getUserRank operation.
	//
	// Retrieve rank of user.
	//
	// GET /users/{UserID}/rank
	GetUserRank(ctx context.Context, params GetUserRankParams) (*Rank, error)
	// ListRanks implements listRanks operation.
	//
	// Get list of ranks.
	//
	// GET /ranks
	ListRanks(ctx context.Context, params ListRanksParams) ([]Rank, error)
	// ListTimes implements listTimes operation.
	//
	// Get list of times.
	//
	// GET /times
	ListTimes(ctx context.Context, params ListTimesParams) ([]Time, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
