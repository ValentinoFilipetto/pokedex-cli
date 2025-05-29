package main

import "fmt"

func commandExplore(config *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Explore command requires argument")
		return nil
	}

	fmt.Printf("Exploring %s...\n", args[0])

	encounters, err := config.pokeapiClient.GetPokemonByLocation(args[0])
	if err != nil {
		return fmt.Errorf("error while fetching Pokemons for %s", args[0])
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range encounters.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
