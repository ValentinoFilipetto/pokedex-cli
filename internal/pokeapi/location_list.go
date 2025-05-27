package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
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

	locationAreas := RespLocations{}
	err = json.Unmarshal(dat, &locationAreas)
	if err != nil {
		return RespLocations{}, err
	}

	return locationAreas, nil
}
