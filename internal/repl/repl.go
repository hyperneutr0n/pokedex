package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/hyperneutr0n/pokedex/internal/commands"
	"github.com/hyperneutr0n/pokedex/internal/pokecache"
)

func Start() {
	cfg := &commands.Config{
		Cache: pokecache.NewCache(5 * time.Second),
	}
	
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanned := scanner.Scan()
		if !scanned {
			break
		}

		input := scanner.Text()

		cliArguments := cleanInput(input)

		if len(cliArguments) == 0 {
			continue
		}

		err := commands.Execute(cliArguments, cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	result := strings.Fields(strings.ToLower(text))
	return result
}
