package main

import (
	"flag"
	"log"
	"os"

	"github.com/Ivlay/go-telegram-bot/pkg/bot"
	"github.com/Ivlay/go-telegram-bot/pkg/htmlParser"
	"github.com/Ivlay/go-telegram-bot/pkg/repository"
	"github.com/Ivlay/go-telegram-bot/pkg/service"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var (
	debug = flag.Bool("debug", false, "-debug=true")
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

	token := os.Getenv("TG_TOKEN")
	if token == "" {
		log.Fatal("bot token must be provided")
	}

	repos := repository.New(db)
	parser := htmlParser.New("https://aj.ru/")
	service := service.New(repos, parser)
	bot, err := bot.New(service, token)
	if err != nil {
		log.Fatal("failed to create bot", err.Error())
	}

	bot.Debug = *debug

	// go bot.Run()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
