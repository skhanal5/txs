PHONY: fmt build run help

fmt:
	go fmt ./...

build:
	go build -o txs ./cmd/txs/main.go

run: build
	./txs

help:
	@echo "Available commands:"
	@echo "  make fmt   - Format the code"
	@echo "  make build - Build the project"
	@echo "  make run   - Run the project"
	@echo "  make help  - Show this help message"