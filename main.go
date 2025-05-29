package main

import (
	"bufio"
	"fmt"
	"strings"
	"time"

	"os"

	"github.com/ValentinoFilipetto/pokedex-cli/internal/pokeapi"
)

var commands map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	pokeapiClient   *pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
}

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	words := strings.Fields(trimmedText)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return words
}

func main() {
	pokeClient := pokeapi.NewClient(30 * time.Second)
	config := &config{
		pokeapiClient: &pokeClient,
	}
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapBack,
		},
		"explore": {
			name: "explore",
			description: "Shows pokemon in a specific location",
			callback: commandExplore,
		},
	}
	reader := os.Stdin
	scanner := bufio.NewScanner(reader)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			inputWords := cleanInput(userInput)
			if len(inputWords) > 0 {
				if cmd, ok := commands[inputWords[0]]; ok {
					err := cmd.callback(config, inputWords[1:])
					if err != nil {
						fmt.Println(err)
					}
					continue
				} else {
					fmt.Printf("Unknown command: %s\n", inputWords[0])
				}
			}
		}
	}
}
