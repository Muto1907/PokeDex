package cmd

import (
	"fmt"
	"os"

	internal "github.com/Muto1907/PokeDex/internal"
)

func CommandExit(conf *internal.Config, area string) error {
	fmt.Println("Closing the PokeDex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(conf *internal.Config, area string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}

func CommandMap(conf *internal.Config, area string) error {
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

func CommandMapB(conf *internal.Config, area string) error {
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

func CommandExplore(conf *internal.Config, area string) error {
	if area == "" {
		return fmt.Errorf("please enter a correct location")
	}
	fmt.Printf("Exploring %s...\n", area)
	location, err := conf.Client.Request_location_area(area)
	if err != nil {
		return err
	}
	pokemons := internal.Get_pokemon_names_from_location_area(location)
	if len(pokemons) > 0 {
		fmt.Println("Found Pokemon:")
		for _, name := range pokemons {
			fmt.Printf("- %s\n", name)
		}
	} else {
		fmt.Printf("No Pokemon found for %s\n", area)
	}
	return nil
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(conf *internal.Config, area string) error
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
		"explore": {
			Name:        "explore",
			Description: "Displays the available Pokemon inside a location Usage: explore <insert location  here>",
			Callback:    CommandExplore,
		},
	}
}
