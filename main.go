package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := os.Stdin
	scanner := bufio.NewScanner(reader)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			fmt.Printf("Your command was: %s\n", cleanInput(userInput)[0])
		}
	}

}

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	words := strings.Fields(trimmedText)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return words
}
