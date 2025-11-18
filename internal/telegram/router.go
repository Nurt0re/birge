package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func RouteUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	if update.Message != nil {
		routeMessage(bot, update.Message)
	}
	if update.CallbackQuery != nil {
		routeCallbackQuery(bot, update.CallbackQuery)
	}
}

func routeMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		handleStart(bot, message)
	case "split":
		handleSplit(bot, message)
	default:
		// Handle other messages
	}

}

func routeCallbackQuery(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {}
