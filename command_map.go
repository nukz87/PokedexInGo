package main

import "fmt"

func commandMap(cfg *config, arg ...string) error {
	var URL *string
	if *cfg.nextLocationURL != "" {
		URL = cfg.nextLocationURL
	}

	locationAreaStruct, err := cfg.pokeClient.ListLocationAreas(URL)
	if err != nil {
		return fmt.Errorf("Error processing api: %v", err)
	}

	for _, result := range locationAreaStruct.Results {
		fmt.Println(result.Name)
	}
	*cfg.nextLocationURL = locationAreaStruct.Next
	*cfg.previousLocationURL = locationAreaStruct.Previous

	return nil
}
