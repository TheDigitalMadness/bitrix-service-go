package main

import (
	"context"
	"fmt"

	"github.com/TheDigitalMadness/bitrix-service-go/internal/client/bitrix"
	"github.com/TheDigitalMadness/bitrix-service-go/internal/config"
	httpController "github.com/TheDigitalMadness/bitrix-service-go/internal/controller/http"
	"github.com/TheDigitalMadness/bitrix-service-go/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func mainWithDi() {
	var module = fx.Options(
		fx.Provide(
			config.New,
			httpController.NewRouter,
			httpController.New,
			newService,
			newBitrixClient,
		),
		fx.Invoke(
			registerRunners,
		),
	)

	fx.New(
		module,
	).Run()
}

func newBitrixClient(cfg *config.Config) service.Client {
	return bitrix.New(cfg.BitrixBaseUrl, cfg.BitrixClientHttpTimeoutInSeconds)
}

func newService(cfg *config.Config, client service.Client) httpController.Service {
	return service.New(client, cfg.BitrixCategoryID)
}

func registerRunners(cfg *config.Config, lc fx.Lifecycle, engine *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := engine.Run(fmt.Sprintf(":%d", cfg.HttpPort)); err != nil {
					panic(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
