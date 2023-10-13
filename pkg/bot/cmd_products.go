package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (b *bot) CmdProducts(upd tgbotapi.Update) {
	pp, err := b.service.Product.GetAllProducts()
	if err != nil {
		b.logger.Error("Error while get all products", zap.Error(err))
		return
	}

	if len(pp) > 0 {
		message := "<b>У нас есть несколько товаров:</b>\n"

		buttons := tgbotapi.NewInlineKeyboardRow()

		for _, p := range pp {
			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s, %d ₽\n", p.Title, p.Price), p.PriceId))
		}

		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			buttons)

		reply := tgbotapi.NewMessage(upd.Message.Chat.ID, message)
		reply.ReplyMarkup = keyboard
		reply.ParseMode = "html"

		if err := b.apiRequest(reply); err != nil {
			b.logger.Error("failed to send products command", zap.Error(err))
		}
	}
}
