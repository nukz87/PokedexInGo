package main

import "fmt"

func commandExplore(c *config, arg ...string) error {
	locationAreaName := arg[0]
	URL := "https://pokeapi.co/api/v2/location-area/" + locationAreaName

	pokemonList, err := c.pokeClient.ListPokemon(&URL)
	if err != nil {
		return fmt.Errorf("Error listing pokemon: %v", err)
	}

	fmt.Printf("Exploring %v...\n", locationAreaName)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonList.Pokemon_Encounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}
	return nil
}
