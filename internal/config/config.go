package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramAPIKey string
}

func NewConfig() *Config {
	godotenv.Load()
	value, err := getEnv("TELEGRAM_API_KEY")
	if err != nil {
		panic(err)
	}
	return &Config{
		TelegramAPIKey: value,
	}
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("could not find the env for %s", key)
	}
	return value, nil

}
