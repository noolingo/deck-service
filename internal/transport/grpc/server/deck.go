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

func (d DeckServer) Create(_ context.Context, _ *noolingo.CreateRequest) (*noolingo.CreateResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (d DeckServer) Delete(_ context.Context, _ *noolingo.DeleteRequest) (*common.Response, error) {
	panic("not implemented") // TODO: Implement
}

func (d DeckServer) List(_ context.Context, _ *noolingo.ListRequest) (*noolingo.ListResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (d DeckServer) Get(_ context.Context, _ *noolingo.GetRequest) (*noolingo.GetResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (d DeckServer) CardAdd(_ context.Context, _ *noolingo.CardAddRequest) (*noolingo.CardAddResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (d DeckServer) CardDelete(_ context.Context, _ *noolingo.CardDeleteRequest) (*common.Response, error) {
	panic("not implemented") // TODO: Implement
}
