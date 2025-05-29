package main

import "fmt"

func commandPokedex(config *config, args []string) error {
	if len(config.pokedex) == 0 {
		fmt.Println("you haven't caught any Pokemons.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
