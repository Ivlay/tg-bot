package repository

import (
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/log/zapadapter"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const (
	usersTable         = "users"
	productsTable      = "products"
	productsListsTable = "products_lists"
)

type Config struct {
	Host     string
	Port     uint16
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config, logger *zap.Logger) (*sqlx.DB, error) {
	connConfig := pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     cfg.Port,
		User:     cfg.Username,
		Password: cfg.Password,
		Database: cfg.DBName,
	}

	connConfig.Logger = zapadapter.NewLogger(logger.Named("database"))

	connPoll, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     connConfig,
		AfterConnect:   nil,
		MaxConnections: 20,
		AcquireTimeout: 30 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	nativeDb := stdlib.OpenDBFromPool(connPoll)

	db := sqlx.NewDb(nativeDb, "pgx")

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
