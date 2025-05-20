package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hyperneutr0n/pokedex/internal/commands"
)

func Start() {
	cfg := &commands.Config{}
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

		err := commands.Execute(words[0], cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	result := strings.Fields(strings.ToLower(text))
	return result
}
