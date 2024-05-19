package repository

import (
	"context"
	"database/sql"

	"github.com/noolingo/deck-service/internal/domain"
)

type Repository interface {
	NewDeck(ctx context.Context, id, name, description string) error
	ListDecks(ctx context.Context, userID string) ([]*domain.Deck, error)
	GetDeck(ctx context.Context, deckID string) (*domain.Deck, error)
	RemoveDeck(ctx context.Context, deckID string) error
	AddCard(ctx context.Context, deckID, cardID string) error
	RemoveCard(ctx context.Context, deckID, cardID string) error
	GetDeckCards(ctx context.Context, deckID string) ([]string, error)
}

func New(db *sql.DB) Repository {
	return &deck{db: db}
}
