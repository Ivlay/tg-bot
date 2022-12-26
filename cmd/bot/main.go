package main

import (
	"log"
	"os"

	"github.com/Ivlay/go-telegram-bot/pkg/htmlParser"
	"github.com/Ivlay/go-telegram-bot/pkg/repository"
	"github.com/Ivlay/go-telegram-bot/pkg/service"
	"github.com/Ivlay/go-telegram-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to init db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	parser := htmlParser.NewParser("https://aj.ru/")
	service := service.NewService(repos, parser)

	service.Product.Prepare()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot, service)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
