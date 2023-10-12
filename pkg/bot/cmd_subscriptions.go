package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (b *bot) CmdSubscriptions(upd tgbotapi.Update) {
	pp, err := b.service.Product.GetByUserId(int(upd.Message.From.ID))
	if err != nil {
		b.logger.Error("Filed to get user subscriptions", zap.Error(err))
		return
	}

	if len(pp) == 0 {
		message := `У вас пока нет активных подписок`
		reply := tgbotapi.NewMessage(upd.Message.Chat.ID, message)
		if err := b.apiRequest(reply); err != nil {
			b.logger.Error("Filed to send subscriptions message when no found", zap.Error(err))
		}
	} else {
		message := ""

		for _, val := range pp {
			if val.Price == 0 {
				message += fmt.Sprintf("%s - цены пока что нет\n", val.Title)
			} else {
				message += fmt.Sprintf("%s - %d ₽\n", val.Title, val.Price)
			}
		}

		reply := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf("<b>Ваши текущие подписки:</b>\n%s", message))

		reply.ParseMode = "html"

		if err := b.apiRequest(reply); err != nil {
			b.logger.Error("Failed to send subscription message with values", zap.Error(err))
		}
	}
}
