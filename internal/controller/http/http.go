package httpController

import (
	"context"
)

type Service interface {
	AddUser(ctx context.Context, email string) (int, error)
	AddProvider(ctx context.Context, email string) (int, error)
	FirstBuy(ctx context.Context, id int) error
	RepeatBuy(ctx context.Context, id int) error
}

type handler struct {
	service Service
}

func New(service Service) *handler {
	return &handler{
		service: service,
	}
}
