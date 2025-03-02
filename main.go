package main

import (
	"bufio"
	"fmt"
	"github.com/shaneplunkett/godex/internal/pokecache"
	"github.com/shaneplunkett/godex/pokeapi"
	"os"
	"strings"
	"time"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func main() {
	commands := createCommands()
	cfg := pokeapi.CreateConfig()
	cache := pokecache.NewCache(5 * time.Millisecond)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {

		}
		text := scanner.Text()
		cleaned := strings.Fields(strings.ToLower(text))
		commandName := cleaned[0]
		command, exists := commands[commandName]
		if exists {
			err := command.callback(cfg, cache)
			if err != nil {
				fmt.Println("Error running command:", err)
			}
		}
	}
}
