package commands

import "fmt"

func pokedex(args []string, cfg *Config) error {
	fmt.Println("Your Pokedex:")
	for name := range UserPokemon{
		fmt.Printf("  - %v\n", name)
	}

	return nil
}