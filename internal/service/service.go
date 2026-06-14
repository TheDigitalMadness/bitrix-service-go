package service

import (
	"context"

	bitrixModels "github.com/TheDigitalMadness/bitrix-service-go/internal/domain/bitrix"
)

type Client interface {
	AddDeal(ctx context.Context, fields *bitrixModels.AddDealFields) (int, error)
	UpdateDeal(ctx context.Context, id int, fields *bitrixModels.UpdateDealFields) error
}

type service struct {
	client Client
}

func New(client Client) *service {
	return &service{client: client}
}

func (s *service) AddUser(ctx context.Context, email string) (int, error) {
	ID, err := s.client.AddDeal(ctx, &bitrixModels.AddDealFields{
		CategoryID: 2,
		StageID:    bitrixModels.REGISTRATION,
		Title:      email,
	})
	if err != nil {
		return 0, err
	}

	return ID, err
}

func (s *service) AddProvider(ctx context.Context, email string) (int, error) {
	ID, err := s.client.AddDeal(ctx, &bitrixModels.AddDealFields{
		CategoryID: 2,
		StageID:    bitrixModels.PROVIDER,
		Title:      email,
	})
	if err != nil {
		return 0, err
	}

	return ID, err
}

func (s *service) FirstBuy(ctx context.Context, id int) error {
	return s.client.UpdateDeal(ctx, id, &bitrixModels.UpdateDealFields{
		CategoryID: 2,
		StageID:    bitrixModels.FIRST_BUY,
	})
}

func (s *service) RepeatBuy(ctx context.Context, id int) error {
	return s.client.UpdateDeal(ctx, id, &bitrixModels.UpdateDealFields{
		CategoryID: 2,
		StageID:    bitrixModels.REPEAT_BUY,
	})
}
