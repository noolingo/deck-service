package service

import (
	"context"
	"errors"

	"github.com/noolingo/deck-service/internal/domain"
	"github.com/noolingo/deck-service/internal/repository"
	"github.com/sirupsen/logrus"
)

var (
	ErrNoDeckFound = errors.New("no such deck found")
)

type DeckService struct {
	logger     *logrus.Logger
	Config     *domain.Config
	repository repository.Repository
}

func NewDeckService(p *Params) *DeckService {
	return &DeckService{
		logger:     p.Logger,
		repository: *p.Repository,
		Config:     p.Config,
	}
}

func (d *DeckService) Create(ctx context.Context, userID string, name string, description string) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DeckService) Delete(ctx context.Context, userID string, deckID string) error {
	panic("not implemented") // TODO: Implement
}

func (d *DeckService) List(ctx context.Context, userID string, deckID string) ([]*domain.Deck, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DeckService) Get(ctx context.Context, userID string, deckID string) ([]string, error) {
	panic("not implemented") // TODO: Implement
}

func (d *DeckService) CardAdd(ctx context.Context, userID string, deckID string, cardID string) error {
	panic("not implemented") // TODO: Implement
}

func (d *DeckService) CardDelete(ctx context.Context, userID string, deckID string, cardID string) error {
	panic("not implemented") // TODO: Implement
}
