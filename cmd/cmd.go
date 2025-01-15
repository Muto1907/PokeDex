package cmd

import (
	"fmt"
	"os"

	internal "github.com/Muto1907/PokeDex/internal"
)

func CommandExit(conf *internal.Config) error {
	fmt.Println("Closing the PokeDex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(conf *internal.Config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}

func CommandMap(conf *internal.Config) error {
	location, err := conf.Client.Request_locations(conf.Next)
	if err != nil {
		return err
	}
	for _, loc := range location.Results {
		fmt.Printf("%s\n", loc.Name)
	}
	conf.Next = location.Next
	conf.Previous = location.Previous
	return nil
}

func CommandMapB(conf *internal.Config) error {
	if conf.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locations, err := conf.Client.Request_locations(conf.Previous)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Printf("%s\n", location.Name)
	}
	conf.Next = locations.Next
	conf.Previous = locations.Previous
	return nil
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(conf *internal.Config) error
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
		"map": {
			Name:        "map",
			Description: "Displays 20 location areas in the Pokemon world",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the 20 previous visited location areas in the Pokemon world",
			Callback:    CommandMapB,
		},
	}
}
