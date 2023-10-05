package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (b *bot) CheckPrice() {
	ids, err := b.service.Product.Update()
	if err != nil {
		b.logger.Error("error while update products", zap.Error(err))
	}

	if len(ids) > 0 {
		uu, err := b.service.Product.GetByUserIds(ids)
		if err != nil {
			b.logger.Error("error while get user subscriptions", zap.Error(err))
			return
		}

		if uu != nil {
			defaultGreeting := "Привет! Цена на товар изменилась:\n"
			products := ""
			for _, u := range uu {
				for _, p := range u.Products {
					products += fmt.Sprintf("%s - %d ₽, <s>%d ₽</s>\n", p.Title, p.Price, p.OldPrice)
				}
				reply := tgbotapi.NewMessage(u.ChatId, defaultGreeting+products)
				reply.ParseMode = "html"
				if err := b.apiRequest(reply); err != nil {
					b.logger.Error("Failed to send start message", zap.Error(err))
				}
			}

		}
	}
}
