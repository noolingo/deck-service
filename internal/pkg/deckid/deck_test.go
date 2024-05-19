package deckid

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDeckID(t *testing.T) {
	userID := "1"
	deckId := NewDeckID(userID)
	userID2, err := ExtractUserID(deckId)
	require.NoError(t, err)
	require.Equal(t, userID2, userID)
}
