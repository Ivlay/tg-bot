package telegram

import (
	"fmt"
	"log"

	tgbot "github.com/Ivlay/go-telegram-bot"
	"github.com/Ivlay/go-telegram-bot/pkg/htmlParser"
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
		return b.handleUnkownCommand(message)
	}
}


func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("From %+v", message.From)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	b.bot.Send(msg)
}

type Qwe struct {
	Id int
}

func (b *Bot) handleNew(message *tgbotapi.Message) error {
	parser := htmlParser.NewParser("https://aj.ru/")
	message.Text = parser.GetPrice()
	
	u := tgbot.User{
		ChatId: message.Chat.ID,
		FirstName: message.Chat.FirstName,
		UserId: message.From.ID,
		UserName: message.Chat.UserName,
	}

	userId, err := b.service.CreateUser(u);

	fmt.Println("userId", userId)

	if err != nil {
		log.Fatalln(err)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
		_, sendErr := b.bot.Send(msg)
		return sendErr
}

func (b *Bot) handleStartMessage(message *tgbotapi.Message) error {
	user, ferr := b.service.GetUserByUserId(message.From.ID)
	if ferr != nil {
		log.Println("Error ->>", ferr)
	}

	fmt.Println("User -->", user)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
		_, err := b.bot.Send(msg)
		return err
}

func (b *Bot) handleUnkownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know this comand")
		_, err := b.bot.Send(msg)
		return err
}