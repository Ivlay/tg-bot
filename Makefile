include .env
.PHONY:
.SILENT:

install:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2

lint:
	@golangci-lint run ./... -v

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

migrateup:
	migrate -path db/migrations -database "postgres://postgres:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgres://postgres:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" -verbose down

migrateforce:
	migrate -path db/migrations -database "postgres://postgres:${DB_PASSWORD}@localhost:5432/${DB_NAME}?sslmode=disable" force 1