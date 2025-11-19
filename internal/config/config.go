package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	TelegramAPIKey string
	PostgresConfig PostgresConfig `yaml:"db"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string
	DBName   string `yaml:"dbname"`
}

func NewConfig() *Config {

	telegramAPIKey, err := getEnv("TELEGRAM_API_KEY")
	if err != nil {
		panic(err)
	}
	dbPassword, err := getEnv("DB_PASSWORD")
	if err != nil {
		panic(err)
	}
	cfg, err := loadConfig("internal/config/config.yaml")
	if err != nil {
		panic(err)
	}
	return &Config{
		TelegramAPIKey: telegramAPIKey,
		PostgresConfig: PostgresConfig{
			Host:     cfg.PostgresConfig.Host,
			Port:     cfg.PostgresConfig.Port,
			User:     cfg.PostgresConfig.User,
			DBName:   cfg.PostgresConfig.DBName,
			Password: dbPassword,
		},
	}
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("could not find the env for %s", key)
	}
	return value, nil

}
func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not load config from yaml: %w", err)
	}
	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal config from yaml: %w", err)
	}
	return &cfg, nil
}
