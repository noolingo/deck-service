package grpcserver

import (
	"context"
	"errors"

	"google.golang.org/grpc/metadata"
)

type Requester struct {
	UserID string
}

func fromContext(ctx context.Context) (*Requester, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no metadata provided")
	}

	userID := md["user_id"]

	if len(userID) == 0 {
		return nil, errors.New("invalid metadata")
	}

	return &Requester{
		UserID: userID[0],
	}, nil
}

func Auth(ctx context.Context) (r *Requester, err error) {
	r, err = fromContext(ctx)
	if err != nil {
		return nil, err
	}

	return r, nil
}
