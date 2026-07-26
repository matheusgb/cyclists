.PHONY: run test lint build up down swag

run:
	go run main.go

swag:
	swag init -g main.go -o docs --parseDependency --parseInternal

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
