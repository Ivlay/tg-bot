package telegram

import (
	"fmt"
	"log"

	tgbot "github.com/Ivlay/go-telegram-bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart = "start"
	new = "new"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartMessage(message)
	case new:
		return b.handleNew(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("From %+v", message.From)

	if message.From.ID == 1652506912 {
	
		msg := tgbotapi.NewMessage(message.Chat.ID, "Пидарок пишет что-то тут")
		b.bot.Send(msg)
	}
}

func (b *Bot) handleNew(message *tgbotapi.Message) error {
	message.Text = b.service.HtmlParser.GetPrice("#macbook_pro_16_2021_")

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
		_, sendErr := b.bot.Send(msg)
		return sendErr
}

func (b *Bot) handleStartMessage(message *tgbotapi.Message) error {
	u := tgbot.User{
		ChatId: message.Chat.ID,
		FirstName: message.Chat.FirstName,
		UserId: message.From.ID,
	}

	if message.Chat.UserName != "" {
		u.UserName = message.Chat.UserName
	} else {
		u.UserName = message.Chat.FirstName
	}

	userId, errr := b.service.FindOrCreate(u)
	if errr != nil {
		log.Default()
	}

	fMsg := fmt.Sprintf("Здарова пидарок, твой id: %d\n", userId)

	msg := tgbotapi.NewMessage(message.Chat.ID, fMsg)
		_, err := b.bot.Send(msg)
		return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know this command")
		_, err := b.bot.Send(msg)
		return err
}
