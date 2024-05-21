package grpcserver

import (
	"context"

	"github.com/noolingo/deck-service/internal/service"
	"github.com/noolingo/proto/codegen/go/common"
	"github.com/noolingo/proto/codegen/go/noolingo"
	"github.com/sirupsen/logrus"
)

type DeckServer struct {
	noolingo.UnimplementedDeckServer
	logger  *logrus.Logger
	service *service.Services
}

func newDeckServer(logger *logrus.Logger, service *service.Services) DeckServer {
	return DeckServer{logger: logger, service: service}
}

func newResponse(err error) (*common.Response, error) {
	response := &common.Response{
		Result: err == nil,
	}
	if err != nil {
		response.Error = &common.Error{
			Error: err.Error(),
		}
	}

	return response, err
}

func (d DeckServer) Create(ctx context.Context, req *noolingo.CreateRequest) (*noolingo.CreateResponse, error) {
	r, err := Auth(ctx)
	if err != nil {
		return nil, err
	}
	id, err := d.service.Deck.Create(ctx, r.UserID, req.Name, req.Description)
	resp, _ := newResponse(err)
	return &noolingo.CreateResponse{DeckID: id, Response: resp.Error}, err

}

func (d DeckServer) Delete(ctx context.Context, req *noolingo.DeleteRequest) (*common.Response, error) {
	r, err := Auth(ctx)
	if err != nil {
		return newResponse(err)
	}
	err = d.service.Deck.Delete(ctx, r.UserID, req.DeckId)
	return newResponse(err)

}

func (d DeckServer) List(ctx context.Context, req *noolingo.ListRequest) (*noolingo.ListResponse, error) {
	r, err := Auth(ctx)
	if err != nil {
		return nil, err
	}
	decks, err := d.service.Deck.List(ctx, r.UserID, req.Id)
	resp, _ := newResponse(err)
	var result []*noolingo.DeckObject
	for _, deck := range decks {
		result = append(result, &noolingo.DeckObject{
			Id:          deck.ID,
			Name:        deck.Name,
			Description: deck.Description,
		})
	}
	return &noolingo.ListResponse{Decks: result, Response: resp.Error}, err
}

func (d DeckServer) Get(ctx context.Context, req *noolingo.GetRequest) (*noolingo.GetResponse, error) {
	r, err := Auth(ctx)
	if err != nil {
		return nil, err
	}
	result, err := d.service.Deck.Get(ctx, r.UserID, req.DeckId)
	resp, _ := newResponse(err)
	return &noolingo.GetResponse{CardIds: result, Response: resp.Error}, err
}

func (d DeckServer) CardAdd(ctx context.Context, req *noolingo.CardAddRequest) (*common.Response, error) {
	r, err := Auth(ctx)
	if err != nil {
		return nil, err
	}
	err = d.service.Deck.CardAdd(ctx, r.UserID, req.DeckId, req.CardId)
	return newResponse(err)

}

func (d DeckServer) CardDelete(ctx context.Context, req *noolingo.CardDeleteRequest) (*common.Response, error) {
	r, err := Auth(ctx)
	if err != nil {
		return newResponse(err)
	}
	err = d.service.Deck.CardDelete(ctx, r.UserID, req.DeckId, req.CardId)
	return newResponse(err)
}
