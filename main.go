package main

import (
	"bufio"
	"fmt"
	"strings"

	"os"

	"github.com/ValentinoFilipetto/pokedex-cli/internal/pokeapi"
)

var commands map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	prev string
	next string
}

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	words := strings.Fields(trimmedText)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return words
}

func commandExit(config *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}

func commandHelp(config *config) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")

	for commandName, command := range commands {
		fmt.Printf("%s: %s\n", commandName, command.description)
	}

	return nil
}

func commandMap(config *config) error {
	prev, next := pokeapi.GetLocationAreas(config.next)
	config.prev = prev
	config.next = next
	return nil
}

func commandMapBack(config *config) error {
	if config.prev == "" {
		fmt.Printf("you're on the first page\n")
		return nil
	}
	prev, next := pokeapi.GetLocationAreas(config.prev)
	config.prev = prev
	config.next = next
	return nil
}

func main() {
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
	}
	reader := os.Stdin
	scanner := bufio.NewScanner(reader)
	config := config{}

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			inputWords := cleanInput(userInput)
			if len(inputWords) > 0 {
				if cmd, ok := commands[inputWords[0]]; ok {
					cmd.callback(&config)
				} else {
					fmt.Printf("Unknown command: %s\n", inputWords[0])
				}
			}
		}
	}
}
