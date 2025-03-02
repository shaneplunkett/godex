package main

import (
	"fmt"
	"github.com/shaneplunkett/godex/internal/pokecache"
	"github.com/shaneplunkett/godex/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *pokeapi.Config, c *pokecache.Cache) error
}

func commandExit(cfg *pokeapi.Config, c *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *pokeapi.Config, c *pokecache.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, c := range createCommands() {
		fmt.Printf("%s: %s\n", name, c.description)
	}
	return nil
}

func commandMap(cfg *pokeapi.Config, c *pokecache.Cache) error {
	res, err := pokeapi.GetArea(cfg)
	if err != nil {
		fmt.Printf("No more pages!")
	}
	for _, area := range res.Results {
		fmt.Printf("%s\n", area.Name)
	}
	return nil

}

func commandMapb(cfg *pokeapi.Config, c *pokecache.Cache) error {
	if cfg.Previous == nil {
		fmt.Printf("You're on the first page")
	}
	cfg.Next = cfg.Previous
	res, err := pokeapi.GetArea(cfg)
	if err != nil {
		fmt.Printf("No more pages!")
	}
	for _, area := range res.Results {
		fmt.Printf("%s\n", area.Name)
	}
	return nil
}

func createCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get next 20 areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "Get previous 20 areas",
			callback:    commandMapb,
		},
	}
}
