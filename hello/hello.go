package hello

import "context"

type WorldResponse struct {
	Message string
}

// World responds with a familiar message.
//
// encore:api public
func World(ctx context.Context) (*WorldResponse, error) {
	return &WorldResponse{Message: "Hello, world!"}, nil
}
