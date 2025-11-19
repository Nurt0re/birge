package telegram

import (
	"birge/internal/service"
	"context"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot     *tgbotapi.BotAPI
	service *service.Service
	log     *slog.Logger
}

func NewHandler(bot *tgbotapi.BotAPI, service *service.Service, log *slog.Logger) *Handler {
	return &Handler{
		bot:     bot,
		service: service,
		log:     log,
	}
}
func (h *Handler) handleStart(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, "Всем салам, всем салам! Я - бот Бiрге. Я пришел помочь вам разделить счет.")
	_, err := h.bot.Send(reply)
	if err != nil {
		h.log.Error("Failed to send start message", "error", err)
	}
}

func (h *Handler) handleSplit(msg *tgbotapi.Message) {
	ctx := context.Background()
	billID, err := h.service.BillService.NewBill(ctx, msg.Chat.ID, msg.From.ID, msg.From.UserName)
	if err != nil {
		h.log.Error("Failed to create bill", "error", err, "chat_id", msg.Chat.ID)
		reply := tgbotapi.NewMessage(msg.Chat.ID, "Произошла ошибка при создании счета. Пожалуйста, попробуйте еще раз.")
		h.bot.Send(reply)
		return
	}

	h.log.Info("Bill created", "bill_id", billID, "creator", msg.From.ID)

	reply := tgbotapi.NewMessage(msg.Chat.ID, "Отлично, давайте начнем разделять счет!\n\n👥 Участники:\n• "+msg.From.FirstName+" (создатель)\n\nНажмите кнопку ниже, чтобы присоединиться:")
	keyboard := JoinBillBtn(billID)
	reply.ReplyMarkup = keyboard
	h.bot.Send(reply)
}
func (h *Handler) handleJoinBill(callbackQuery *tgbotapi.CallbackQuery, billID int64) {
	ctx := context.Background()
	err := h.service.BillService.AddUserToBill(ctx, billID, callbackQuery.From.ID, callbackQuery.From.UserName)
	if err != nil {
		h.log.Error("Failed to add user to bill", "error", err, "bill_id", billID, "user_id", callbackQuery.From.ID)
		callback := tgbotapi.NewCallback(callbackQuery.ID, "Произошла ошибка при добавлении вас к счету. Пожалуйста, попробуйте еще раз.")
		h.bot.Request(callback)
		return
	}

	h.log.Info("User joined bill", "bill_id", billID, "user_id", callbackQuery.From.ID, "username", callbackQuery.From.UserName)

	callback := tgbotapi.NewCallback(callbackQuery.ID, "Вы успешно присоединились к счету")
	h.bot.Request(callback)

	participants, err := h.service.BillService.GetBillParticipants(ctx, billID)
	if err != nil {
		h.log.Error("Failed to get participants", "error", err, "bill_id", billID)
		return
	}

	msgText := "Отлично, давайте начнем разделять счет!\n\n👥 Участники:\n"
	for _, p := range participants {
		msgText += "• " + p.Username + "\n"
	}
	msgText += "\nНажмите кнопку ниже, чтобы присоединиться:"

	edit := tgbotapi.NewEditMessageText(
		callbackQuery.Message.Chat.ID,
		callbackQuery.Message.MessageID,
		msgText,
	)
	edit.ReplyMarkup = &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: JoinBillBtn(billID).InlineKeyboard,
	}
	h.bot.Send(edit)
}

func (h *Handler) handleMarkPaid(callbackQuery *tgbotapi.CallbackQuery) {
	callback := tgbotapi.NewCallback(callbackQuery.ID, "Функция в разработке")
	h.bot.Request(callback)
}
