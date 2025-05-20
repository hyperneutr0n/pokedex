package commands

import "fmt"

func GetRegistry() map[string]command {
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
		"map":{
			name:					"map",
			description: 	"Display names of location areas in Pokemon world",
			callback: 		getMap,
		},
	}
}

func Execute(commandString string) error {
	registry := GetRegistry()

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