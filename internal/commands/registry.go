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
			description: 	"Display 20 names of location areas in Pokemon world",
			callback: 		getMap,
		},
		"mapb": {
			name:					"mapb",
			description: 	"Display 20 previous names of location areas in Pokemon world",
			callback: 		getMapBack,
		},
	}
}

func Execute(commandName string, cfg *Config) error {
	registry := GetRegistry()

	cmd, exist := registry[commandName]
	if !exist {
		return fmt.Errorf("Unknown command '%v'.", commandName)
	}

	err := cmd.callback(cfg)
	if err != nil {
		return fmt.Errorf("Something went wrong. \n%w", err)
	}

	return nil
}