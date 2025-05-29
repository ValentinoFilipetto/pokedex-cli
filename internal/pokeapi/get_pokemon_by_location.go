package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonByLocation(location string) (Encounters, error) {
	url := baseURL + "/location-area/" + location
	encounters := Encounters{}

    // check if we already fetched encounters for that url
	// if so, return encounters from the cache
	items, ok := c.cache.Get(url)

	if ok {
		err := json.Unmarshal(items, &encounters)
		if err != nil {
			return Encounters{}, err
		}
		return encounters, nil
	}


	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Encounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Encounters{}, err
	}

	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Encounters{}, err
	}

	// cache response
	c.cache.Add(url, dat)

	err = json.Unmarshal(dat, &encounters)
	if err != nil {
		return Encounters{}, err
	}

	return encounters, err
}