package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetLocationAreas(url string) (string, string) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	res, err := http.Get(url)

	type LocationAreas struct {
		Count    int    `json:"count"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}

	locationAreas := LocationAreas{}

	dat, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, dat)
	}

	err = json.Unmarshal(dat, &locationAreas)

	if err != nil {
		return "", ""
	}

	// if locationAreas.Previous == "" {
	// 	fmt.Printf("you're on the first page\n")
	// 	return locationAreas.Previous, locationAreas.Next
	// }

	for _, area := range locationAreas.Results {
		fmt.Println(area.Name)
	}

	return locationAreas.Previous, locationAreas.Next
}
