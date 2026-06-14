package main

import (
	"fmt"

	"github.com/TheDigitalMadness/bitrix-service-go/internal/client/bitrix"
	"github.com/TheDigitalMadness/bitrix-service-go/internal/config"
	httpController "github.com/TheDigitalMadness/bitrix-service-go/internal/controller/http"
	"github.com/TheDigitalMadness/bitrix-service-go/internal/service"
)

func main() {
	cfg := config.New()
	bitrixClient := bitrix.New(cfg.BitrixBaseUrl, cfg.BitrixClientHttpTimeoutInSeconds)
	service := service.New(bitrixClient, cfg.BitrixCategoryID)
	handler := httpController.New(service)

	r := httpController.NewRouter(handler)

	r.Run(fmt.Sprintf(":%d", cfg.HttpPort))
}
