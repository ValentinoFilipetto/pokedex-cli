package main

import (
	"fmt"
)

// commandMap retrieves and displays the next 20 location areas from the PokeAPI.
func commandMap(config *config, args []string) error {
	locations, err := config.pokeapiClient.GetLocationAreas(config.nextLocationUrl)
	if err != nil {
		return err
	}

	config.prevLocationUrl = locations.Previous
	config.nextLocationUrl = locations.Next

	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	return nil
}

// commandMapBack retrieves and displays the previous 20 location areas from the PokeAPI.
func commandMapBack(config *config, args []string) error {
	if config.prevLocationUrl == nil {
		fmt.Printf("you're on the first page\n")
		return nil
	}
	locations, err := config.pokeapiClient.GetLocationAreas(config.prevLocationUrl)
	if err != nil {
		return err
	}

	config.prevLocationUrl = locations.Previous
	config.nextLocationUrl = locations.Next

	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	return nil
}
