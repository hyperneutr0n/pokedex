# Makefile

BINARY_NAME=bin/pokedex
MAIN_PATH=./cmd/pokedex

.PHONY: build run clean

build:
	go build -o $(BINARY_NAME) $(MAIN_PATH)

run:
	go run $(MAIN_PATH)

clean:
	rm -f $(BINARY_NAME)
