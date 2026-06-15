package bitrix

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

// Adds a deal to any stage of pre-configured category
func (c *client) AddDeal(ctx context.Context, fields *bitrixModels.AddDealFields) (int, error) {
	url := c.baseUrl + "/crm.deal.add"

	reqBody := map[string]any{
		"fields": map[string]any{
			"CATEGORY_ID": fields.CategoryID,
			"STAGE_ID":    fields.StageID,
			"TITLE":       fields.Title,
		},
	}
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Bad status: %s", resp.Status)
	}

	var ID int
	if err := json.NewDecoder(resp.Body).Decode(&ID); err != nil {
		return 0, err
	}

	return ID, nil
}

// Changes a deal stage in category by id
func (c *client) UpdateDeal(ctx context.Context, id int, fields *bitrixModels.UpdateDealFields) error {
	url := c.baseUrl + "/crm.deal.update"

	reqBody := map[string]any{
		"fields": map[string]any{
			"CATEGORY_ID": fields.CategoryID,
			"STAGE_ID":    fields.StageID,
		},
		"id": id,
	}
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
