package service

import (
	"context"

	"github.com/TheDigitalMadness/bitrix-service-go/internal/client/bitrix"
)

type Service interface {
	AddUser(ctx context.Context, email string) error
	AddProvider(ctx context.Context, email string) error
	FirstBuy(ctx context.Context, id int) error
	RepeatBuy(ctx context.Context, id int) error
}

type service struct {
	client bitrix.Client
}

func New(client bitrix.Client) Service {
	return &service{client: client}
}

func (s *service) AddUser(ctx context.Context, email string) error {
	return nil
}

func (s *service) AddProvider(ctx context.Context, email string) error {
	return nil
}

func (s *service) FirstBuy(ctx context.Context, id int) error {
	return nil
}

func (s *service) RepeatBuy(ctx context.Context, id int) error {
	return nil
}
