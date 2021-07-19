PROJECT_NAME=$(shell basename "$(PWD)")

build:
	go build -o bin/$(PROJECT_NAME) cmd/$(PROJECT_NAME)/main.go

run:
	go run cmd/$(PROJECT_NAME)/main.go

test:
	go test ./... -v

.PHONY: run build test