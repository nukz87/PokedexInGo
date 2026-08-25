package main

import "fmt"

func commandMap(cfg *config) error {
	var URL *string
	if *cfg.nextLocationURL != "" {
		URL = cfg.nextLocationURL
	}

	apiStruct, err := cfg.pokeClient.ListLocationAreas(URL)
	if err != nil {
		return fmt.Errorf("Error processing api: %v", err)
	}

	for _, result := range apiStruct.Results {
		fmt.Println(result.Name)
	}
	*cfg.nextLocationURL = apiStruct.Next
	*cfg.previousLocationURL = apiStruct.Previous

	return nil
}
