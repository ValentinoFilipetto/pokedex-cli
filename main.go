package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	words := strings.Fields(trimmedText)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return words
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")

	for commandName, command := range commands {
		fmt.Printf("%s: %s\n", commandName, command.description)
	}

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
					cmd.callback()
				} else {
					fmt.Printf("Unknown command: %s\n", inputWords[0])
				}
			}
		}
	}
}
