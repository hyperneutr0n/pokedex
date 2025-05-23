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
		"explore": {
			name: 				"explore",
			description: 	"Explore an area of choice, passed in by the first argument",
			callback: 		explore,
		},
		"catch": {
			name:					"catch",
			description: 	"Catch a pokemon, pass the pokemon name in the first argument",
			callback: 		catch,
		},
		"inspect": {
			name:					"inspect",
			description: 	"Inspect a pokemon that you have caught. Pass it on the first argument",
			callback: 		inspect,
		},
		"pokedex": {
			name:					"pokedex",
			description: 	"See all the pokemon you've caught",
			callback: 		pokedex,
		},
	}
}

func Execute(cliArguments []string, cfg *Config) error {
	registry := GetRegistry()

	cmd, exist := registry[cliArguments[0]]
	if !exist {
		return fmt.Errorf("Unknown command '%v'.", cliArguments[0])
	}

	err := cmd.callback(cliArguments[1:], cfg)
	if err != nil {
		return fmt.Errorf("Something went wrong. \n%w", err)
	}

	return nil
}