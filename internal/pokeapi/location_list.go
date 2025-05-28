package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	locationAreas := RespLocations{}

	if pageURL != nil {
		url = *pageURL
	}

	// check if we already fetched locations for that url
	// if so, return locations from the cache
	items, ok := c.cache.Get(url)

	if ok {
		err := json.Unmarshal(items, &locationAreas)
		if err != nil {
			return RespLocations{}, err
		}
		return locationAreas, nil
	}


	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}

	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocations{}, err
	}

	// cache response
	c.cache.Add(url, dat)

	err = json.Unmarshal(dat, &locationAreas)
	if err != nil {
		return RespLocations{}, err
	}

	return locationAreas, nil
}
