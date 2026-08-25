package main

import "fmt"

func commandMapb(cfg *config) error {
	var URL *string
	if *cfg.previousLocationURL != "" {
		URL = cfg.previousLocationURL
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
