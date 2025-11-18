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

	keyboard := AddParticipantsBtn()
	reply.ReplyMarkup = keyboard
	bot.Send(reply)
	_ = billID
}
