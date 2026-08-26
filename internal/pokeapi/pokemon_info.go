package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GettingPokemonInfo(pageURL *string) (Pokemon, error) {
	URL := *pageURL

	if entry, ok := c.pokeCache.Get(URL); ok {
		var finalStruct Pokemon
		if err := json.Unmarshal(entry, &finalStruct); err != nil {
			return Pokemon{}, fmt.Errorf("Error json unmarshaling: %v", err)
		}
		return finalStruct, nil
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error making request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error calling api: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Error convert io.ReadCloser to []byte: %v", err)
	}
	c.pokeCache.Add(URL, data)

	var finalStruct Pokemon
	if err := json.Unmarshal(data, &finalStruct); err != nil {
		return Pokemon{}, fmt.Errorf("Error json unmarshaling: %v", err)
	}

	return finalStruct, nil
}
