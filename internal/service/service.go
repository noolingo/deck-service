package service

import (
	"context"

	"github.com/noolingo/deck-service/internal/domain"
	"github.com/noolingo/deck-service/internal/repository"
	"github.com/sirupsen/logrus"
)

type Deck interface {
	Create(ctx context.Context, userID, name, description string) (string, error)
	Delete(ctx context.Context, userID, deckID string) error
	List(ctx context.Context, userID, deckID string) ([]*domain.Deck, error)
	Get(ctx context.Context, userID, deckID string) ([]string, error)
	CardAdd(ctx context.Context, userID, deckID, cardID string) error
	CardDelete(ctx context.Context, userID, deckID, cardID string) error
}

type Services struct {
	Deck Deck
}

type Params struct {
	Logger     *logrus.Logger
	Config     *domain.Config
	Repository *repository.Repository
}

func New(p *Params) *Services {
	return &Services{
		Deck: NewDeckService(p),
	}
}
