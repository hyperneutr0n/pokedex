package commands

import "fmt"

func help(args []string, cfg *Config) error {
	registry := GetRegistry()
	fmt.Printf(`
Welcome to the Pokedex!

Usage:
`)

	for i, c := range registry {
		fmt.Printf("%v: %v\n", i, c.description)
	} 
	
	fmt.Println()
	return nil
}