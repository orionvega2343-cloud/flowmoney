package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Bot struct {
	ApiUrl string `yaml:"api_url"`
}

type Config struct {
	Bot Bot `yaml:"bot"`

	Token string `env:"BOT_TOKEN"`

	// ProxyUrl опционален: пусто — бот идёт в api.telegram.org напрямую.
	// Если Telegram недоступен напрямую (например, провайдер режет его без
	// VPN), сюда подставляется локальный прокси-порт Happ или другого
	// VPN-клиента: http://127.0.0.1:PORT или socks5://127.0.0.1:PORT.
	ProxyUrl string `env:"PROXY_URL"`
}

func MustLoad() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var cfg = Config{}

	err = cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal("Error loading config")
	}
	return &cfg
}
