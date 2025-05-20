package main

import "fmt"

func getRegistry() map[string]command {
	return map[string]command{
		"exit": {
			name:					"exit",
			description:	"Exit the pokedex",
			callback:			exit,
		},
		"help": {
			name:					"help",
			description: 	"Display a help message",
			callback: 		help,
		},
	}
}

func routing(commandString string) error {
	registry := getRegistry()

	commandToExecute, exist := registry[commandString]
	if !exist {
		return fmt.Errorf("Unknown command '%v'.", commandString)
	}

	err := commandToExecute.callback()
	if err != nil {
		return fmt.Errorf("Something went wrong. \n%w", err)
	}

	return nil
}