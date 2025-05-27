package main

import (
	"fmt"
)

// commandHelp displays the available commands and their descriptions.
func commandHelp(config *config) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
