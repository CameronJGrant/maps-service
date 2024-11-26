package controller

import (
	"context"
	"fmt"
	"git.itzana.me/strafesnet/maps-service/internal/datastore"
	"git.itzana.me/strafesnet/maps-service/internal/model"
	"git.itzana.me/strafesnet/go-grpc/maps"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Maps struct {
	*maps.UnimplementedMapsServiceServer
	Store datastore.Datastore
}

func (m Maps) Get(ctx context.Context, message *maps.IdMessage) (*maps.MapResponse, error) {
	item, err := m.Store.Maps().Get(ctx, message.GetID())
	if err != nil {
		if err == datastore.ErrNotExist {
			return nil, status.Error(codes.NotFound, "map does not exit")
		}
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to get map:").Error())
	}

	return &maps.MapResponse{
		ID:          item.ID,
		DisplayName: item.DisplayName,
		Creator:     item.Creator,
		GameID:      item.GameID,
		Date:        item.Date.Unix(),
	}, nil
}

func (m Maps) GetList(ctx context.Context, list *maps.IdList) (*maps.MapList, error) {
	items, err := m.Store.Maps().GetList(ctx, list.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to get maps").Error())
	}

	var resp maps.MapList
	for i := 0; i < len(items); i++ {
		resp.Maps = append(resp.Maps, &maps.MapResponse{
			ID:          items[i].ID,
			DisplayName: items[i].DisplayName,
			Creator:     items[i].Creator,
			GameID:      items[i].GameID,
			Date:        items[i].Date.Unix(),
		})
	}

	return &resp, nil
}

func (m Maps) Update(ctx context.Context, request *maps.MapRequest) (*maps.NullResponse, error) {
	updates := datastore.Optional()
	updates.AddNotNil("display_name", request.DisplayName)
	updates.AddNotNil("creator", request.Creator)
	updates.AddNotNil("game_id", request.GameID)
	if request.Date != nil {
		updates.AddNotNil("date", time.Unix(request.GetDate(), 0))
	}

	if err := m.Store.Maps().Update(ctx, request.GetID(), updates); err != nil {
		if err == datastore.ErrNotExist {
			return nil, status.Error(codes.NotFound, "map does not exit")
		}
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to update map:").Error())
	}

	return &maps.NullResponse{}, nil
}

func (m Maps) Create(ctx context.Context, request *maps.MapRequest) (*maps.IdMessage, error) {
	item, err := m.Store.Maps().Create(ctx, model.Map{
		ID:          request.GetID(),
		DisplayName: request.GetDisplayName(),
		Creator:     request.GetCreator(),
		GameID:      request.GetGameID(),
		Date:        time.Unix(request.GetDate(), 0),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to create map:").Error())
	}

	return &maps.IdMessage{ID: item.ID}, nil
}

func (m Maps) Delete(ctx context.Context, message *maps.IdMessage) (*maps.NullResponse, error) {
	if err := m.Store.Maps().Delete(ctx, message.GetID()); err != nil {
		if err == datastore.ErrNotExist {
			return nil, status.Error(codes.NotFound, "map does not exit")
		}
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to delete map:").Error())
	}

	return &maps.NullResponse{}, nil
}

func (m Maps) List(ctx context.Context, request *maps.ListRequest) (*maps.MapList, error) {
	filter := datastore.Optional()
	fmt.Println(request)
	if request.Filter != nil {
		filter.AddNotNil("display_name", request.GetFilter().DisplayName)
		filter.AddNotNil("creator", request.GetFilter().Creator)
		filter.AddNotNil("game_id", request.GetFilter().GameID)
	}

	items, err := m.Store.Maps().List(ctx, filter, model.Page{
		Number: request.GetPage().GetNumber(),
		Size:   request.GetPage().GetSize(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, errors.Wrap(err, "failed to get maps:").Error())
	}

	var resp maps.MapList
	for i := 0; i < len(items); i++ {
		resp.Maps = append(resp.Maps, &maps.MapResponse{
			ID:          items[i].ID,
			DisplayName: items[i].DisplayName,
			Creator:     items[i].Creator,
			GameID:      items[i].GameID,
			Date:        items[i].Date.Unix(),
		})
	}

	return &resp, nil
}
