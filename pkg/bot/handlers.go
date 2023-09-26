package bot

import (
	"fmt"
	"log"

	tgbot "github.com/Ivlay/go-telegram-bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
	new          = "new"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)

func (b *bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartMessage(message)
	case new:
		return b.handleNew(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("From %+v", message.From)

	if message.From.ID == 1652506912 {

		msg := tgbotapi.NewMessage(message.Chat.ID, "Пидарок пишет что-то тут")
		b.bot.Send(msg)
	}
}

func (b *bot) handleNew(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "In progress...")
	_, sendErr := b.bot.Send(msg)
	return sendErr
}

func (b *bot) handleStartMessage(message *tgbotapi.Message) error {
	u := tgbot.User{
		ChatId:    message.Chat.ID,
		FirstName: message.Chat.FirstName,
		UserId:    int(message.From.ID),
	}

	if message.Chat.UserName != "" {
		u.UserName = message.Chat.UserName
	} else {
		u.UserName = message.Chat.FirstName
	}

	userId, err := b.service.FindOrCreate(u)
	if err != nil {
		log.Fatal(err.Error())
	}

	fMsg := fmt.Sprintf("Здарова %s, твой id: %d\n", u.UserName, userId)

	msg := tgbotapi.NewMessage(message.Chat.ID, fMsg)
	msg.ReplyMarkup = numericKeyboard
	_, err = b.bot.Send(msg)
	return err
}

func (b *bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
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

func (b *bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know this command")
	_, err := b.Send(msg)
	return err
}
