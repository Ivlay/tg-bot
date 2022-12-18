include .env
.PHONY:
.SILENT:

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