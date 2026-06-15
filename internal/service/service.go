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
	client     Client
	categoryID int
}

func New(client Client, categoryID int) *service {
	return &service{client: client, categoryID: categoryID}
}

// Adds a deal with user to registration stage
func (s *service) AddUser(ctx context.Context, email string) (int, error) {
	ID, err := s.client.AddDeal(ctx, &bitrixModels.AddDealFields{
		CategoryID: s.categoryID,
		StageID:    bitrixModels.REGISTRATION,
		Title:      email,
	})
	if err != nil {
		return 0, err
	}

	return ID, err
}

// Adds a deal with user to provider stage
func (s *service) AddProvider(ctx context.Context, email string) (int, error) {
	ID, err := s.client.AddDeal(ctx, &bitrixModels.AddDealFields{
		CategoryID: s.categoryID,
		StageID:    bitrixModels.PROVIDER,
		Title:      email,
	})
	if err != nil {
		return 0, err
	}

	return ID, err
}

// Moves a deal to first-buy stage
func (s *service) FirstBuy(ctx context.Context, id int) error {
	return s.client.UpdateDeal(ctx, id, &bitrixModels.UpdateDealFields{
		CategoryID: s.categoryID,
		StageID:    bitrixModels.FIRST_BUY,
	})
}

// Moves a deal to repeat-buy stage
func (s *service) RepeatBuy(ctx context.Context, id int) error {
	return s.client.UpdateDeal(ctx, id, &bitrixModels.UpdateDealFields{
		CategoryID: s.categoryID,
		StageID:    bitrixModels.REPEAT_BUY,
	})
}
