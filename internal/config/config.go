package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	HttpPort                         int
	BitrixBaseUrl                    string
	BitrixClientHttpTimeoutInSeconds int
}

// https://b24-5ji7no.bitrix24.ru
// 1
// TopSecret
func New() *Config {
	return &Config{
		HttpPort:      4000,
		BitrixBaseUrl: fmt.Sprintf("%s/rest/%s/%s", os.Getenv("BITRIX_URL"), os.Getenv("BITRIX_USER_ID"), os.Getenv("BITRIX_WEBHOOK_SECRET")),
		BitrixClientHttpTimeoutInSeconds: func() int {
			val, err := strconv.ParseInt(os.Getenv("BITRIX_CLIENT_HTTP_TIMEOUT"), 10, 64)
			if err != nil || val <= 0 {
				val = 6
			}
			return int(val)
		}(),
	}
}
