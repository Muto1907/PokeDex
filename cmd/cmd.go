package cmd

import (
	"fmt"
	"math/rand"
	"os"

	internal "github.com/Muto1907/PokeDex/internal"
)

func CommandExit(conf *internal.Config, args ...string) error {
	fmt.Println("Closing the PokeDex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(conf *internal.Config, args ...string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}

func CommandMap(conf *internal.Config, args ...string) error {
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

func CommandMapB(conf *internal.Config, args ...string) error {
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

func CommandExplore(conf *internal.Config, args ...string) error {
	if args[0] == "" {
		return fmt.Errorf("please enter a correct location")
	}
	fmt.Printf("Exploring %s...\n", args[0])
	location, err := conf.Client.Request_location_area(args[0])
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
		fmt.Printf("No Pokemon found for %s\n", args[0])
	}
	return nil
}

func CommandCatch(conf *internal.Config, args ...string) error {
	pokemon, err := conf.Client.Request_pokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	catch_chance := rand.Float64()
	if catch_chance > 1-(100/float64(pokemon.BaseExperience)) {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		conf.PokeDex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(conf *internal.Config, args ...string) error
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
		"catch": {
			Name:        "catch <Pokemon>",
			Description: "Try to catch a Pokemon",
			Callback:    CommandCatch,
		},
	}
}
