package main

import (
	"fmt"
	"os"
)

// commandExit prints a message and exists from the CLI.
func commandExit(config *config, args []string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}
