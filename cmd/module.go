package main

import (
	httpController "github.com/TheDigitalMadness/bitrix-service-go/internal/controller/http"
	"github.com/TheDigitalMadness/bitrix-service-go/internal/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		httpController.NewRouter,
		httpController.New,
		service.New,
	),
	fx.Invoke(StartServer),
)

func StartServer() {
	// Тут буду сервак запускать
}
