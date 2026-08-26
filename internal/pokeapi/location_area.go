package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (RespLocationAreasStruct, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if entry, ok := c.pokeCache.Get(url); ok {
		var finalStruct RespLocationAreasStruct
		if err := json.Unmarshal(entry, &finalStruct); err != nil {
			return RespLocationAreasStruct{}, fmt.Errorf("Error json umarshaling: %v", err)
		}
		return finalStruct, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationAreasStruct{}, fmt.Errorf("Error making request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationAreasStruct{}, fmt.Errorf("Error sending sending request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return RespLocationAreasStruct{}, fmt.Errorf("Error status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	var finalStruct RespLocationAreasStruct
	if err := json.Unmarshal(data, &finalStruct); err != nil {
		return RespLocationAreasStruct{}, fmt.Errorf("Error unmarshaling data: %v", err)
	}

	c.pokeCache.Add(url, data)

	return finalStruct, nil
}
