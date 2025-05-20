package commands

import "fmt"

func help() error {
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