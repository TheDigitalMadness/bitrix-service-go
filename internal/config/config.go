package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HttpPort int `env:"PORT" env-default:"4000"`

	BitrixURL           string `env:"BITRIX_URL" env-required:"true"`
	BitrixUserID        string `env:"BITRIX_USER_ID" env-required:"true"`
	BitrixWebhookSecret string `env:"BITRIX_WEBHOOK_SECRET" env-required:"true"`

	BitrixClientHttpTimeoutInSeconds int `env:"BITRIX_CLIENT_HTTP_TIMEOUT" env-default:"6"`

	BitrixCategoryID int `env:"BITRIX_CATEGORY_ID" env-default:"2"`

	BitrixBaseUrl string `env:"-"`
}

// Read environment from .env file and build and return config
func MustReturnConfig() *Config {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}

	cfg.BitrixBaseUrl = fmt.Sprintf(
		"%s/rest/%s/%s",
		cfg.BitrixBaseUrl,
		cfg.BitrixUserID,
		cfg.BitrixWebhookSecret,
	)

	return &cfg
}
