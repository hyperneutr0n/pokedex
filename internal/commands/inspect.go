package commands

import "fmt"

func inspect(args []string, cfg *Config) error {
	if _, exist := UserPokemon[args[0]]; !exist {
		return fmt.Errorf("Pokemon %s hasn't been caught.", args[0])
	}
	for _, pokemon := range UserPokemon {
		fmt.Println("Name: " + pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for stat, val := range pokemon.Stats {
			fmt.Printf("  -%s: %v\n", stat, val)
		}
		fmt.Println("Types:")
		for _, pokeType := range pokemon.Types {
			fmt.Printf("  - %s\n", pokeType)
		}
	}
	return nil
}