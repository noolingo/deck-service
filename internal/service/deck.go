package service

import (
	"context"
	"errors"

	"github.com/noolingo/deck-service/internal/domain"
	"github.com/noolingo/deck-service/internal/pkg/deckid"
	"github.com/noolingo/deck-service/internal/repository"
	"github.com/noolingo/proto/codegen/go/apierrors"
	"github.com/sirupsen/logrus"
)

var (
	ErrNoDeckFound = errors.New("no such deck found")
)

type DeckService struct {
	logger     *logrus.Logger
	Config     *domain.Config
	repository repository.Repository
	clients    GrpcClients
}

func NewDeckService(p *Params) *DeckService {
	return &DeckService{
		clients:    p.GrpcClient,
		logger:     p.Logger,
		repository: *p.Repository,
		Config:     p.Config,
	}
}

func (d *DeckService) Create(ctx context.Context, userID string, name string, description string) (string, error) {
	deckID := deckid.NewDeckID(userID)
	err := d.repository.NewDeck(ctx, deckID, userID, name, description)
	if err != nil {
		return "", err
	}
	return deckID, nil
}

func (d *DeckService) Delete(ctx context.Context, userID string, deckID string) error {
	userID2, err := deckid.ExtractUserID(deckID)
	if err != nil {
		return err
	}
	if userID2 != userID {
		return apierrors.ErrForbidden
	}
	err = d.repository.RemoveDeck(ctx, deckID)
	return err
}

func (d *DeckService) List(ctx context.Context, userID string, deckID string) ([]*domain.Deck, error) {
	res, err := d.repository.ListDecks(ctx, userID)
	return res, err
}

func (d *DeckService) Get(ctx context.Context, userID string, deckID string) ([]string, error) {
	userID2, err := deckid.ExtractUserID(deckID)
	if err != nil {
		return nil, err
	}
	if userID2 != userID {
		return nil, apierrors.ErrForbidden
	}
	cardIDs, err := d.repository.GetDeckCards(ctx, deckID)
	return cardIDs, err
}

// сделать запрос на CardService при помощи "клиента"
func (d *DeckService) CardAdd(ctx context.Context, userID string, deckID string, cardID string) error {
	userID2, err := deckid.ExtractUserID(deckID)
	if err != nil {
		return err
	}
	if userID2 != userID {
		return apierrors.ErrForbidden
	}
	deck, err := d.repository.GetDeck(ctx, deckID)
	if err != nil {
		return err
	}
	if deck == nil {
		return apierrors.ErrInvalidPayload
	}
	if err = d.clients.CardExistsByID(ctx, cardID); err != nil {
		return err
	}
	err = d.repository.AddCard(ctx, deckID, cardID)
	return err
}

func (d *DeckService) CardDelete(ctx context.Context, userID string, deckID string, cardID string) error {
	userID2, err := deckid.ExtractUserID(deckID)
	if err != nil {
		return err
	}
	if userID2 != userID {
		return apierrors.ErrForbidden
	}
	err = d.repository.RemoveCard(ctx, deckID, cardID)
	return err
}
