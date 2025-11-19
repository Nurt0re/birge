package telegram

import (
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Router struct {
	handler *Handler
}

func NewRouter(handler *Handler) *Router {
	return &Router{
		handler: handler,
	}
}

func (r *Router) RouteUpdate(update tgbotapi.Update) {
	if update.Message != nil {
		r.routeMessage(update.Message)
	}
	if update.CallbackQuery != nil {
		r.routeCallbackQuery(update.CallbackQuery)
	}
}

func (r *Router) routeMessage(message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		r.handler.handleStart(message)
	case "split":
		r.handler.handleSplit(message)
	default:
		// Handle other messages
	}

}

func (r *Router) routeCallbackQuery(callbackQuery *tgbotapi.CallbackQuery) {

	parts := strings.Split(callbackQuery.Data, ":")
	action := parts[0]
	switch action {
	case "join_bill":
		if len(parts) < 2 {
			callback := tgbotapi.NewCallback(callbackQuery.ID, "Неправильный формат")
			r.handler.bot.Request(callback)
			return
		}
		billID, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			callback := tgbotapi.NewCallback(callbackQuery.ID, "Неправильный формат ID счета")
			r.handler.bot.Request(callback)
			return
		}
		r.handler.handleJoinBill(callbackQuery, billID)
	case "mark_paid":
		r.handler.handleMarkPaid(callbackQuery)
	default:
		// Handle other callback queries
	}
}
