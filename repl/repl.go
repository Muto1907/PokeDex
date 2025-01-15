package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Muto1907/PokeDex/cmd"
	"github.com/Muto1907/PokeDex/internal"
)

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

func StartREPL(config *internal.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) == 0 {
			continue
		}
		word := cleanedInput[0]
		if cmd, ok := cmd.GetCommands()[word]; ok {
			err := cmd.Callback(config)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
