package main

import (
	"fmt"
	"github.com/shaneplunkett/godex/internal/pokecache"
	"github.com/shaneplunkett/godex/pokeapi"
	"math/rand"
	"os"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *pokeapi.Config, c *pokecache.Cache, p string) error
}

func commandExit(cfg *pokeapi.Config, c *pokecache.Cache, p string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *pokeapi.Config, c *pokecache.Cache, p string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, c := range createCommands() {
		fmt.Printf("%s: %s\n", name, c.description)
	}
	return nil
}

func commandMap(cfg *pokeapi.Config, c *pokecache.Cache, p string) error {
	res, err := pokeapi.GetArea(cfg, c)
	if err != nil {
		fmt.Printf("No more pages!")
	}
	for _, area := range res.Results {
		fmt.Printf("%s\n", area.Name)
	}
	return nil

}

func commandMapb(cfg *pokeapi.Config, c *pokecache.Cache, p string) error {
	if cfg.Previous == nil {
		fmt.Printf("You're on the first page")
	}
	cfg.Next = cfg.Previous
	res, err := pokeapi.GetArea(cfg, c)
	if err != nil {
		fmt.Printf("No more pages!")
	}
	for _, area := range res.Results {
		fmt.Printf("%s\n", area.Name)
	}
	return nil
}

func commandExplore(cfg *pokeapi.Config, c *pokecache.Cache, p string) error {
	res, err := pokeapi.GetAreaId(cfg, c, p)
	if err != nil {
		fmt.Printf("Oops Please Try Again!")
	}
	fmt.Printf("Exploring: %v\n", p)
	fmt.Printf("Found Pokemon:\n")
	for _, poke := range res.PokemonEncounters {
		fmt.Printf("- %s\n", poke.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *pokeapi.Config, c *pokecache.Cache, p string) error {
	res, err := pokeapi.GetPokemonSpecies(cfg, c, p)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", res.Name)
	time.Sleep(1 * time.Second)
	fmt.Print("*shake*...\n")
	time.Sleep(1 * time.Second)
	fmt.Print("*shake*...\n")
	time.Sleep(1 * time.Second)
	fmt.Print("*shake*...\n")
	time.Sleep(1 * time.Second)
	if rand.Intn(256) < res.CaptureRate {
		fmt.Printf("%v was caught!\n", res.Name)
	} else {
		fmt.Printf("Oh No! %v broke free!\n", res.Name)
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
		"explore": {
			name:        "explore <area_name>",
			description: "List all Pokemon in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Attempt to Catch a Pokemon",
			callback:    commandCatch,
		},
	}
}
