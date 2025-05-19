package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanned := scanner.Scan()
		if !scanned {
			break
		}

		input := scanner.Text()

		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}

		fmt.Printf("Your command was: %v\n", words[0])
	}
}