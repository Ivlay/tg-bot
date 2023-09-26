package bot

import (
	"log"

	"github.com/Ivlay/go-telegram-bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type bot struct {
	*tgbotapi.BotAPI
	commands map[commandKey]commandEntity
	service  *service.Service
}

func New(service *service.Service, token string) (*bot, error) {
	api, aErr := tgbotapi.NewBotAPI(token)
	if (aErr) != nil {
		return nil, aErr
	}

	b := &bot{
		BotAPI:   api,
		service:  service,
		commands: make(map[commandKey]commandEntity),
	}

	if err := b.initCommands(); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *bot) Start() error {
	log.Printf("Authorized on account %s", b.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.initUpdateChannel()

	b.handleUpdates(updates)

	return nil
}

func (b *bot) initUpdateChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.GetUpdatesChan(u)
}

func (b *bot) apiRequest(c tgbotapi.Chattable) error {
	_, err := b.Request(c)

	return err
}
