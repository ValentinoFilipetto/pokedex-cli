package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(config *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Catch command requires argument")
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	pokemon, err := config.pokeapiClient.GetPokemonByName(args[0])
	if err != nil {
		return fmt.Errorf("error while fetching Pokemon for")
	}

	// the higher the Pokemon's base experience, the harder it is to catch it.
	threshold := 50
    random_number := rand.IntN(pokemon.BaseExperience + 1)
	if random_number < threshold {
		fmt.Printf("%s was caught!\n", args[0])
		caughtPokemon := Pokemon{
			ID:     pokemon.ID,
			Name:   pokemon.Name,
			Species: pokemon.Species,
		}

		config.pokedex[pokemon.Name] = caughtPokemon
		return nil
	}

    fmt.Printf("%s escaped!\n", args[0])
	return nil
}