package main

import "fmt"

func commandInspect(config *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Inspect command requires argument")
		return nil
	}

	pokemon, ok := config.pokedex[args[0]]

	if ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			if stat.Stat.Name == "hp" {
				fmt.Printf(" -hp: %d\n", stat.BaseStat)
			} else if stat.Stat.Name == "attack" {
				fmt.Printf(" -attack: %d\n", stat.BaseStat)
			} else if stat.Stat.Name == "defense" {
				fmt.Printf(" -defense: %d\n", stat.BaseStat)
			} else if stat.Stat.Name == "special-attack" {
				fmt.Printf(" -special-attack: %d\n", stat.BaseStat)
			} else if stat.Stat.Name == "special-defense" {
				fmt.Printf(" -special-defense: %d\n", stat.BaseStat)
			} else if stat.Stat.Name == "speed" {
				fmt.Printf(" -speed: %d\n", stat.BaseStat)
			}
		}
		fmt.Printf("Types:\n")
		for _, tp := range pokemon.Types {
			fmt.Printf(" - %s\n", tp.Type.Name)
		}
	}

	fmt.Println("you have not caught that pokemon")
	return nil
}
