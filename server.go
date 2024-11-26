package main

import (
	"context"
	"log"
	"net/http"

	"git.itzana.me/strafesnet/maps-service/api"
	"git.itzana.me/strafesnet/maps-service/internal/controller/submissions"
	"git.itzana.me/strafesnet/maps-service/internal/controller/users"
)

type apiServer struct {
	submissions *submissions.Submissions
	users *users.Users
}

// GetUserRank implements api.Handler.
func (m *apiServer) GetSubmission(ctx context.Context, params api.GetSubmissionParams) (*api.Submission, error) {
	return m.submissions.Get(ctx,params)
}

// NewError implements api.Handler.
func (m *apiServer) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	return &api.ErrorStatusCode{
		StatusCode: 500,
		Response:   api.Error{Message: err.Error()},
	}
}

// GetTimes implements api.Handler.
func (m *apiServer) ListTimes(ctx context.Context, params api.ListTimesParams) ([]api.Time, error) {
	client := times.NewTimesServiceClient(m.client)

	// Call the List method using params
	req := &times.ListRequest{
		Page: &times.Pagination{
			Number: params.Page.GetPage(),
			Size:   params.Page.GetLimit(),
		},
	}

	if filter, ok := params.Filter.Get(); ok {
		if id := filter.GetID(); id.IsSet() {
			req.Filter.ID = &id.Value
		}
		if time := filter.GetTime(); time.IsSet() {
			req.Filter.Time = &time.Value
		}
		if userID := filter.GetUserID(); userID.IsSet() {
			req.Filter.UserID = &userID.Value
		}
		if mapID := filter.GetMapID(); mapID.IsSet() {
			req.Filter.MapID = &mapID.Value
		}
		if styleID := filter.GetStyleID(); styleID.IsSet() {
			req.Filter.StyleID = &styleID.Value
		}
		if modeID := filter.GetModeID(); modeID.IsSet() {
			req.Filter.ModeID = &modeID.Value
		}
		if gameID := filter.GetGameID(); gameID.IsSet() {
			req.Filter.GameID = &gameID.Value
		}
	}

	response, err := client.List(ctx, req)
	if err != nil {
		return nil, err
	}

	return convertTimes(response.Times), nil
}

// GetUser implements api.Handler.
func (m *apiServer) GetUser(ctx context.Context, params api.GetUserParams) (*api.User, error) {
	client := users.NewUsersServiceClient(m.client)

	response, err := client.Get(ctx, &users.IdMessage{
		ID: params.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &api.User{
		ID:       api.NewOptInt64(response.ID),
		Username: api.NewOptString(response.Username),
		StateID:  api.NewOptInt32(response.StateID),
	}, nil
}

// ListRanks implements api.Handler.
func (m *apiServer) ListRanks(ctx context.Context, params api.ListRanksParams) ([]api.Rank, error) {
	client := ranks.NewRanksServiceClient(m.client)

	req := &ranks.ListRequest{
		Page: &ranks.Pagination{
			Number: params.Page.GetPage(),
			Size:   params.Page.GetLimit(),
		},
	}

	if filter, ok := params.Filter.Get(); ok {
		if gameID, ok := filter.GetGameID().Get(); ok {
			req.GameID = gameID
		}

		if modeID, ok := filter.GetModeID().Get(); ok {
			req.ModeID = modeID
		}

		if styleID, ok := filter.GetStyleID().Get(); ok {
			req.StyleID = styleID
		}

		if sort, ok := filter.GetSort().Get(); ok {
			req.Sort = sort
		}
	}

	response, err := client.List(ctx, req)
	if err != nil {
		return nil, err
	}

	ranks := make([]api.Rank, len(response.Ranks))
	for i, r := range response.Ranks {
		ranks[i] = api.Rank{
			ID: api.NewOptInt64(r.ID),
			User: api.NewOptUser(api.User{
				ID:       api.NewOptInt64(r.User.ID),
				Username: api.NewOptString(r.User.Username),
				StateID:  api.NewOptInt32(r.User.StateID),
			}),
			StyleID:   api.NewOptInt32(r.StyleID),
			ModeID:    api.NewOptInt32(r.ModeID),
			GameID:    api.NewOptInt32(r.GameID),
			Rank:      api.NewOptFloat64(r.Rank),
			Skill:     api.NewOptFloat64(r.Skill),
			UpdatedAt: api.NewOptInt64(r.UpdatedAt),
		}
	}

	return ranks, nil
}

func main() {
	// new grpc client
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	svc := &apiServer{
		client: conn,
	}

	srv, err := api.NewServer(
		svc,
		api.WithPathPrefix("/v2"),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}

func convertTime(t *times.TimeResponse) api.Time {
	return api.Time{
		ID:   api.NewOptInt64(t.ID),
		Time: api.NewOptInt64(t.Time),
		User: api.NewOptUser(api.User{
			ID:       api.NewOptInt64(t.User.ID),
			Username: api.NewOptString(t.User.Username),
			StateID:  api.NewOptInt32(t.User.StateID),
		}),
		Map: api.NewOptMap(api.Map{
			ID:          api.NewOptInt64(t.Map.ID),
			DisplayName: api.NewOptString(t.Map.DisplayName),
			Creator:     api.NewOptString(t.Map.Creator),
			Date:        api.NewOptInt64(t.Map.Date),
		}),
		Date:    api.NewOptInt64(t.Date),
		StyleID: api.NewOptInt32(t.StyleID),
		ModeID:  api.NewOptInt32(t.ModeID),
		GameID:  api.NewOptInt32(t.GameID),
	}
}

func convertTimes(timeList []*times.TimeResponse) []api.Time {
	times := make([]api.Time, len(timeList))
	for i, t := range timeList {
		times[i] = convertTime(t)
	}
	return times
}
