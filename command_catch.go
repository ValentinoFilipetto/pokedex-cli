package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
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
	max := big.NewInt(int64(pokemon.BaseExperience + 1))
	randomBigInt, err := rand.Int(rand.Reader, max)
	if err != nil {
		return fmt.Errorf("error generating random number: %w", err)
	}
	random_number := int(randomBigInt.Int64()) // Convert *big.Int to int
	if random_number < threshold {
		fmt.Printf("%s was caught!\n", args[0])
		caughtPokemon := pokemon

		config.pokedex[pokemon.Name] = caughtPokemon
		return nil
	}

	fmt.Printf("%s escaped!\n", args[0])
	return nil
}
