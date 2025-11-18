package main

import (
	"birge/internal/config"
	"birge/internal/telegram"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	cfg := config.NewConfig()
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramAPIKey)
	if err != nil {
		panic(err)
	}

	log.Println("Bot started...")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		telegram.RouteUpdate(bot, update)
	}
}
