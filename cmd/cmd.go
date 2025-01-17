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
		conf.PokeDex[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

func CommandInspect(conf *internal.Config, args ...string) error {
	pokemon, ok := conf.PokeDex[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typee := range pokemon.Types {
		fmt.Printf("  - %s\n", typee.Type.Name)
	}
	return nil
}

func CommandPokedex(conf *internal.Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for key := range conf.PokeDex {
		fmt.Printf("  - %s\n", key)
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
		"inspect": {
			Name:        "inspect <Pokemon>",
			Description: "Lookup all your caught Pokemon in your Pokedex",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "List the names of all your caught Pokemon",
			Callback:    CommandPokedex,
		},
	}
}
