# Makefile

BINARY_NAME=bin/pokedex
MAIN_PATH=./cmd/pokedex
COMMAND_PATH=internal/commands

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

command:
	touch $(COMMAND_PATH)/$(NAME)
	echo "package commands" > $(COMMAND_PATH)/$(NAME)
	code $(COMMAND_PATH)/$(NAME)