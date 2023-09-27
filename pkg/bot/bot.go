package bot

import (
	"github.com/Ivlay/go-telegram-bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

// 1652506912
type bot struct {
	*tgbotapi.BotAPI
	logger   *zap.Logger
	commands map[commandKey]commandEntity
	service  *service.Service
}

func New(logger *zap.Logger, service *service.Service, token string) (*bot, error) {
	api, aErr := tgbotapi.NewBotAPI(token)
	if (aErr) != nil {
		return nil, aErr
	}

	logger = logger.Named("bot")
	b := &bot{
		BotAPI:   api,
		service:  service,
		logger:   logger,
		commands: make(map[commandKey]commandEntity),
	}

	if err := b.initCommands(); err != nil {
		return nil, err
	}

	b.logger.Info("bot created", zap.String("username", api.Self.UserName))
	return b, nil
}

func (b *bot) apiRequest(c tgbotapi.Chattable) error {
	_, err := b.Request(c)

	return err
}
