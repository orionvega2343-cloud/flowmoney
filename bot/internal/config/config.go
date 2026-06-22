package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	ApiUrl string `yaml:"api_url" env:"API_URL"`
	Token  string `env:"BOT_TOKEN"`
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
