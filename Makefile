.PHONY: run test lint build up down

run:
	go run main.go

test:
	go test -v ./src/tests/...

lint:
	golangci-lint run ./...

build:
	go build -o cyclists .

up:
	docker compose up --build

down:
	docker compose down
