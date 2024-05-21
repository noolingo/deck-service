package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/noolingo/deck-service/internal/domain"
)

type deck struct {
	db *sql.DB
}

func (d *deck) NewDeck(ctx context.Context, deckID string, userID string, name string, description string) error {
	ins, err := d.db.PrepareContext(ctx,
		"insert into deck_descr(id, user_id, deck_name, deck_descr) values(?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = ins.ExecContext(ctx, deckID, userID, name, description)
	if err != nil {
		return err
	}
	return err
}

func (d *deck) ListDecks(ctx context.Context, userID string) ([]*domain.Deck, error) {
	var decks []*domain.Deck
	stmt := `SELECT id, deck_name, deck_descr from deck_descr where user_id=?`
	rows, err := d.db.QueryContext(ctx, stmt, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		deck := &domain.Deck{}
		err = rows.Scan(&deck.ID, &deck.Name, &deck.Description)
		if err != nil {
			return nil, err
		}
		decks = append(decks, deck)
	}
	return decks, err
}

func (d *deck) GetDeck(ctx context.Context, deckID string) (*domain.Deck, error) {
	deck := &domain.Deck{ID: deckID}
	err := d.db.QueryRowContext(ctx, "select deck_name, deck_descr from deck_descr where id=?", deckID).Scan(
		&deck.Name,
		&deck.Description,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return deck, nil

}

func (d *deck) RemoveDeck(ctx context.Context, deckID string) error {
	ins, err := d.db.PrepareContext(ctx, "delete from deck_descr where id=?")
	if err != nil {
		return err
	}
	_, err = ins.ExecContext(ctx, deckID)
	if err != nil {
		return err
	}
	ins, err = d.db.PrepareContext(ctx, "delete from decks where deck_id=?")
	if err != nil {
		return err
	}
	_, err = ins.ExecContext(ctx, deckID)
	return err
}

func (d *deck) AddCard(ctx context.Context, deckID string, cardID string) error {
	ins, err := d.db.PrepareContext(ctx,
		"insert into decks(deck_id, card_id) values (?,?)")
	if err != nil {
		return err
	}
	_, err = ins.ExecContext(ctx, deckID, cardID)
	return err
}

func (d *deck) RemoveCard(ctx context.Context, deckID string, cardID string) error {
	ins, err := d.db.PrepareContext(ctx, "delete from decks where deck_id=? and card_id=?")
	if err != nil {
		return err
	}
	_, err = ins.ExecContext(ctx, deckID, cardID)
	return err
}

func (d *deck) GetDeckCards(ctx context.Context, deckID string) ([]string, error) {
	var cards_IDs []string
	stmt := `SELECT card_id from decks where deck_id=?`
	rows, err := d.db.QueryContext(ctx, stmt, deckID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var card string
		err = rows.Scan(&card)
		if err != nil {
			return nil, err
		}
		cards_IDs = append(cards_IDs, card)
	}
	return cards_IDs, err
}
