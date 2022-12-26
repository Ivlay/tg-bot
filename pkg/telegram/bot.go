package telegram

import (
	"log"

	"github.com/Ivlay/go-telegram-bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	service *service.Service
}

func NewBot(bot *tgbotapi.BotAPI, services *service.Service) *Bot {
	return &Bot{
		bot:     bot,
		service: services,
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.initUpdateChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				err := b.handleCommand(update.Message)
				if err != nil {
					log.Fatal(err)
				}
				continue
			}

			b.handleMessage(update.Message)
		}
	}
}

func (b *Bot) initUpdateChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
