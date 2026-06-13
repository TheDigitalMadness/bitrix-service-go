package bitrix

import (
	"context"
	"net/http"
	"time"
)

type AddDealFields struct {
	CategoryID int
	StageID    string
	Title      string
}

type UpdateDealFields struct {
	CategoryID int
	StageID    string
}

type Client interface {
	AddDeal(ctx context.Context, fields AddDealFields) error
	UpdateDeal(ctx context.Context, id int, fields UpdateDealFields) error
}

type client struct {
	baseUrl string
	http    *http.Client
}

func New(baseUrl string) Client {
	return &client{
		baseUrl: baseUrl,
		http: &http.Client{
			Timeout: 6 * time.Second,
		},
	}
}

func (c *client) AddDeal(ctx context.Context, fields AddDealFields) error {
	return nil
}

func (c *client) UpdateDeal(ctx context.Context, id int, fields UpdateDealFields) error {
	return nil
}
