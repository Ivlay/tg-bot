package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ivlay/go-telegram-bot/pkg/bot"
	"github.com/Ivlay/go-telegram-bot/pkg/htmlParser"
	"github.com/Ivlay/go-telegram-bot/pkg/repository"
	"github.com/Ivlay/go-telegram-bot/pkg/service"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	debug    = flag.Bool("debug", false, "-debug=true")
	schedule = flag.String("schedule", "@every 6h", "-schedule=@every 6h")
)

func init() {
	flag.Parse()
}

func Log(message time.Time) {
	fmt.Printf("Message from cron %s\n", message)
}

func main() {
	logger := mustLogger(*debug)
	// cron := cron.New(cron.WithLogger(cron.DefaultLogger))

	if err := initConfig(); err != nil {
		log.Fatalf("error init config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetUint16("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  viper.GetString("db.sslmode"),
	}, logger)

	if err != nil {
		log.Fatalf("failed to init db: %s", err.Error())
	}

	token := os.Getenv("TG_TOKEN")
	if token == "" {
		log.Fatal("bot token must be provided")
	}

	parser := htmlParser.New("https://aj.ru/")
	repos := repository.New(db)
	service := service.New(repos, parser)
	bot, err := bot.New(logger, service, token)
	if err != nil {
		log.Fatal("failed to create bot", err.Error())
	}

	service.Product.Prepare()

	bot.Debug = *debug

	go bot.Run()

	bot.CheckPrice()

	// cron.Start()

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)

	<-stopCh

	logger.Info("Bot stopped")
}

func mustLogger(debug bool) *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.DisableStacktrace = true

	logLevel := zapcore.InfoLevel
	if debug {
		logLevel = zapcore.DebugLevel
	}
	cfg.Level = zap.NewAtomicLevelAt(logLevel)

	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err.Error())
	}

	return logger
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
