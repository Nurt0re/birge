package main

import (
	"birge/internal/config"
	"birge/internal/db"
	"birge/internal/repository"
	"birge/internal/service"
	"birge/internal/telegram"
	"log/slog"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		logger.Warn("Error loading .env file", "error", err)
	}
	cfg := config.NewConfig()
	db, err := db.ConnectPostgres(cfg)
	if err != nil {
		logger.Error("Failed to connect to Postgres", "error", err)
	}
	defer db.Close()

	repo := repository.NewRepository(db, logger)
	svc := service.NewService(repo, logger)

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramAPIKey)
	if err != nil {
		logger.Error("Failed to initialize bot", "error", err)
		panic(err)
	}
	logger.Info("Bot started", "username", bot.Self.UserName)

	// Initialize handler and router
	handler := telegram.NewHandler(bot, svc, logger)
	router := telegram.NewRouter(handler)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		router.RouteUpdate(update)
	}
}
