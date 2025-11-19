package telegram

import (
	"birge/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleStart(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {

	reply := tgbotapi.NewMessage(msg.Chat.ID, "Всем салам, всем салам! Я - бот Бiрге. Я пришел помочь вам разделить счет.")
	bot.Send(reply)

}

func handleSplit(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {

	billID, err := service.NewBill(msg.Chat.ID, msg.From.ID)
	if err != nil {
		reply := tgbotapi.NewMessage(msg.Chat.ID, "Произошла ошибка при создании счета. Пожалуйста, попробуйте еще раз.")
		bot.Send(reply)
		return
	}
	reply := tgbotapi.NewMessage(msg.Chat.ID, "Отлично, давайте начнем разделять счет!\nДобавьте участников чека трапезы:")

	keyboard := JoinBillBtn(billID)
	reply.ReplyMarkup = keyboard
	bot.Send(reply)

}
func handleJoinBill(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery, billID int64) {
	err := service.AddUserToBill(billID, callbackQuery.From.ID, callbackQuery.From.UserName)
	if err != nil {
		callback := tgbotapi.NewCallback(callbackQuery.ID, "Произошла ошибка при добавлении вас к счету. Пожалуйста, попробуйте еще раз.")
		bot.Request(callback)
		return
	}

	callback := tgbotapi.NewCallback(callbackQuery.ID, "Вы успешно присоединились к счету")
	bot.Request(callback)

	participants, err := service.GetBillParticipants(billID)
	if err != nil {
		return
	}

	msgText := "Текущие участники счета:\n"
	for _, p := range participants {
		msgText += "- " + p.Username + "\n"
	}

	edit := tgbotapi.NewEditMessageText(
		callbackQuery.Message.Chat.ID,
		callbackQuery.Message.MessageID,
		msgText,
	)
	edit.ReplyMarkup = &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: JoinBillBtn(billID).InlineKeyboard,
	}
	bot.Send(edit)
}
