package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddParticipantsBtn() tgbotapi.InlineKeyboardMarkup {
	button := tgbotapi.NewInlineKeyboardButtonData("Добавить участника", "add_participant")
	row := tgbotapi.NewInlineKeyboardRow(button)
	return tgbotapi.NewInlineKeyboardMarkup(row)
}
func MarkPaidBtn() tgbotapi.InlineKeyboardMarkup {
	btn := tgbotapi.NewInlineKeyboardButtonData("I paid 💸", "mark_paid")
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(btn),
	)
}

func JoinBillBtn(billID int64) tgbotapi.InlineKeyboardMarkup {
	btn := tgbotapi.NewInlineKeyboardButtonData("Присоединиться к счету", fmt.Sprintf("join_bill:%d", billID))
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(btn),
	)
}
