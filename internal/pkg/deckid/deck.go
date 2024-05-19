package deckid

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

func NewDeckID(userID string) string {
	tmp := ulid.Make()
	return fmt.Sprintf("%v%v", tmp.String(), userID)
}

func ExtractUserID(deckID string) (string, error) {
	if len(deckID) <= ulid.EncodedSize {
		return "", ulid.ErrDataSize
	}
	return deckID[ulid.EncodedSize:], nil
}
