package cmd

import (
	"fmt"
	"os"
)

func CommandExit() error {
	fmt.Println("Closing the PokeDex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the PokeDex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
	}
}
