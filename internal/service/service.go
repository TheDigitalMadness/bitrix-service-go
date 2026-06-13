package service

import (
	"context"

	bitrixModels "github.com/TheDigitalMadness/bitrix-service-go/internal/domain/bitrix"
)

type Client interface {
	AddDeal(ctx context.Context, fields bitrixModels.AddDealFields) error
	UpdateDeal(ctx context.Context, id int, fields bitrixModels.UpdateDealFields) error
}

type service struct {
	client Client
}

func New(client Client) *service {
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
