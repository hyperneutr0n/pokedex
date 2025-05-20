package commands

import (
	"fmt"
	"os"
)

func exit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
