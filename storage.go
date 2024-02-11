package client

import "context"

type Storage interface {
	Create(ctx context.Context, client *Client) error
	FindAll(ctx context.Context) (clients []Client, err error)
	FindOne(ctx context.Context, id string) (Client, error)
	Update(ctx context.Context, client Client, args ...interface{}) error
	Delete(ctx context.Context, id string) error
}
