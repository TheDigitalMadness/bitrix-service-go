package config

type Config struct {
	HttpPort                         int
	BitrixBaseUrl                    string
	BitrixClientHttpTimeoutInSeconds int
}

// Типа реализовал получение данных из конфига
func New() *Config {
	return &Config{
		HttpPort:                         4000,
		BitrixBaseUrl:                    "https://b24-5ji7no.bitrix24.ru/rest/1/TopSecret",
		BitrixClientHttpTimeoutInSeconds: 6,
	}
}
