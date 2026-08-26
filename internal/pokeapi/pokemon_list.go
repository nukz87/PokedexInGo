package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(pageURL *string) (RespPokemonInLocationAreas, error) {
	URL := *pageURL

	if entry, ok := c.pokeCache.Get(URL); ok {
		var finalStruct RespPokemonInLocationAreas
		if err := json.Unmarshal(entry, &finalStruct); err != nil {
			return RespPokemonInLocationAreas{}, fmt.Errorf("Error json unmarshaling: %v", err)
		}
		return finalStruct, nil
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return RespPokemonInLocationAreas{}, fmt.Errorf("Error making new request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonInLocationAreas{}, fmt.Errorf("Error do request %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return RespPokemonInLocationAreas{}, fmt.Errorf("Error status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespPokemonInLocationAreas{}, fmt.Errorf("Error convert io.ReadCloser to []byte: %v", err)
	}
	c.pokeCache.Add(URL, data)

	var finalStruct RespPokemonInLocationAreas
	if err := json.Unmarshal(data, &finalStruct); err != nil {
		return RespPokemonInLocationAreas{}, fmt.Errorf("Error json unmarshaling: %v", err)
	}

	return finalStruct, nil
}
