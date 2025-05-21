# Makefile

BINARY_NAME=bin/pokedex
MAIN_PATH=./cmd/pokedex

.PHONY: build run clean

build:
	mkdir -p $(dir $(BINARY_NAME))
	go build -o $(BINARY_NAME) $(MAIN_PATH)

run:
	go run $(MAIN_PATH)

clean:
	rm -f $(BINARY_NAME)

test:
	go test -v ./internal/...
