package bot

import (
	"fmt"
	"log"

	tgbot "github.com/Ivlay/go-telegram-bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type replyKeyboardValue string

const (
	ReplyProducts = replyKeyboardValue("Товары")
)

func (b *bot) CmdStart(upd tgbotapi.Update) {
	name := upd.Message.From.UserName

	if name == "" {
		name = upd.Message.From.FirstName
	}

	u := tgbot.User{
		ChatId:    upd.Message.Chat.ID,
		FirstName: upd.Message.Chat.FirstName,
		UserId:    int(upd.Message.From.ID),
		UserName:  name,
	}

	_, err := b.service.FindOrCreate(u)
	if err != nil {
		log.Fatal(err.Error())
	}

	message := `Добро пожаловать в <b>AJ price check</b>, %s!`

	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf(message, name))
	reply.ParseMode = "html"

	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(string(ReplyProducts)),
		),
	)

	reply.ReplyMarkup = keyboard
	reply.DisableWebPagePreview = true

	if err := b.apiRequest(reply); err != nil {
		b.logger.Error("Failed to send start message", zap.Error(err))
	}
}
