package repository

import (
	"context"
	"database/sql"

	"github.com/noolingo/deck-service/internal/domain"
)

type deck struct {
	db *sql.DB
}

func (d *deck) NewDeck(ctx context.Context, deckID string, userID string, name string, description string) error {
	_, err := d.db.PrepareContext(ctx,
		"insert into deck_descr")
	return err

}

func (d *deck) ListDecks(ctx context.Context, userID string) ([]*domain.Deck, error) {
	panic("not implemented") // TODO: Implement
}

func (d *deck) GetDeck(ctx context.Context, deckID string) (*domain.Deck, error) {
	panic("not implemented") // TODO: Implement
}

func (d *deck) RemoveDeck(ctx context.Context, deckID string) error {
	panic("not implemented") // TODO: Implement
}

func (d *deck) AddCard(ctx context.Context, deckID string, cardID string) error {
	panic("not implemented") // TODO: Implement
}

func (d *deck) RemoveCard(ctx context.Context, deckID string, cardID string) error {
	panic("not implemented") // TODO: Implement
}

func (d *deck) GetDeckCards(ctx context.Context, deckID string) ([]string, error) {
	panic("not")
}
