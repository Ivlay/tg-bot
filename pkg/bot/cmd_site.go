package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

const (
	siteURL = "https://aj.ru/"
)

func (b *bot) CmdSite(upd tgbotapi.Update) {
	message := fmt.Sprintf("Все цены берутся с сайта <a href='%s'>AJ.ru</a>\n", siteURL)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Перейти", siteURL),
		),
	)

	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, message)
	reply.ReplyMarkup = keyboard
	reply.ParseMode = "html"

	if err := b.apiRequest(reply); err != nil {
		b.logger.Error("failed to send site command", zap.Error(err))
	}
}
