package clients

import (
	"context"

	"github.com/noolingo/proto/codegen/go/noolingo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	cardClient noolingo.CardsClient
}

func NewClients(cardsService string) (*Clients, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(cardsService, opts...)
	if err != nil {
		return nil, err
	}
	return &Clients{cardClient: noolingo.NewCardsClient(conn)}, nil
}

func (c *Clients) CardExistsByID(ctx context.Context, cardID string) error {
	_, err := c.cardClient.SearchByID(ctx, &noolingo.SearchByIDRequest{Id: []string{cardID}})
	if err != nil {
		return err
	}
	return nil
}
