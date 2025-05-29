package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonByName(name string) (PokemonResponse, error) {
	url := baseURL + "/pokemon/" + name
	pokemon := PokemonResponse{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}

	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return PokemonResponse{}, err
	}

	return pokemon, err
}