package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	result := strings.Fields(strings.ToLower(text))
	return result
}

func exit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func help() error {
	registry := getRegistry()
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