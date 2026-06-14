package bitrix

import (
	"context"
	"net/http"
	"time"

	bitrixModels "github.com/TheDigitalMadness/bitrix-service-go/internal/domain/bitrix"
)

type client struct {
	baseUrl string
	http    *http.Client
}

func New(baseUrl string, httpTimeoutInSeconds int) *client {
	return &client{
		baseUrl: baseUrl,
		http: &http.Client{
			Timeout: time.Duration(httpTimeoutInSeconds) * time.Second,
		},
	}
}

func (c *client) AddDeal(ctx context.Context, fields bitrixModels.AddDealFields) error {
	return nil
}

func (c *client) UpdateDeal(ctx context.Context, id int, fields bitrixModels.UpdateDealFields) error {
	return nil
}
