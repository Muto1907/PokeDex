package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Muto1907/PokeDex/cmd"
	"github.com/Muto1907/PokeDex/internal"
	"github.com/Muto1907/PokeDex/internal/pokecache"
)

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

func StartREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &internal.Config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: nil,
		Cache:    pokecache.NewCache(5 * time.Second),
	}
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
