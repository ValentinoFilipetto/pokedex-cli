package pokeapi

// If Next and Previous were just regular string types,
// we'd have to represent the absence of a URL with an empty string, which could be ambiguous.
type RespLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
