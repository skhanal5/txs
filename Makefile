PHONY: fmt build run rund help

fmt:
	go fmt ./...

build:
	go build -o txs ./cmd/txs/main.go

run: build
	./txs

cleanv:
	docker compose down -v

cleand:
	docker compose down

rund: cleand
	docker compose up --build -d

help:
	@echo "Available commands:"
	@echo "  make fmt   - Format the code"
	@echo "  make build - Build the project"
	@echo "  make run   - Run the project"
	@echo "  make cleanv - Remove Docker containers and volumes"
	@echo "  make cleand - Stop and remove Docker containers"
	@echo "  make rund  - Start Docker containers in detached mode"
	@echo "  make help  - Show this help message"