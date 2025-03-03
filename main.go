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
	cache := pokecache.NewCache(10 * time.Second)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			fmt.Println("Oops! Error has Occured")

		}
		text := scanner.Text()
		cleaned := cleanInput(text)
		param := ""
		if len(cleaned) > 1 {
			param = cleaned[1]
		}
		commandName := cleaned[0]
		command, exists := commands[commandName]
		if exists {
			err := command.callback(cfg, cache, param)
			if err != nil {
				fmt.Println("Error running command:", err)
			}
		}
	}
}
