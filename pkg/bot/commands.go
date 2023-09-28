package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type commandEntity struct {
	key    commandKey
	desc   string
	action func(upd tgbotapi.Update)
}

const (
	StartCmdKey          = commandKey("start")
	MySubscriptionCmdKey = commandKey("subscriptions")
	ProductsKey          = commandKey("products")
)

type commandKey string

func (b *bot) initCommands() error {
	commands := []commandEntity{
		{
			key:    StartCmdKey,
			desc:   "Запусить бота",
			action: b.CmdStart,
		},
		{
			key:    MySubscriptionCmdKey,
			desc:   "Посмотреть мои подписики",
			action: b.CmdSubscriptions,
		},
	}

	tgCommands := make([]tgbotapi.BotCommand, 0, len(commands))
	for _, cmd := range commands {
		b.commands[cmd.key] = cmd
		tgCommands = append(tgCommands, tgbotapi.BotCommand{
			Command:     "/" + string(cmd.key),
			Description: cmd.desc,
		})
	}

	config := tgbotapi.NewSetMyCommands(tgCommands...)

	return b.apiRequest(config)
}

func (b *bot) replyToCommand(text string) (commandEntity, bool) {
	switch replyKeyboardValue(text) {
	case ReplyProducts:
		cmd, ok := b.commands[ProductsKey]
		return cmd, ok
	}

	return commandEntity{}, false
}
