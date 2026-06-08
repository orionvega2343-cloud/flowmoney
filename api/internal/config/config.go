package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type DB struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Username string `yaml:"user"`
	Name     string `yaml:"name"`
	SslMode  string `yaml:"ssl_mode"`
	Pass     string `env:"DB_PASS"`
}

type Server struct {
	Port int `yaml:"port"`
}

type Jwt struct {
	JWTSecret string `env:"JWT_SECRET"`
}

type Config struct {
	DB     DB     `yaml:"db"`
	Server Server `yaml:"server"`
	Jwt    Jwt    `yaml:"jwt"`
}

func MustLoad() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var c Config

	err = cleanenv.ReadConfig("config.yml", &c)
	if err != nil {
		log.Fatal("Error loading config")
	}
	return &c

}
