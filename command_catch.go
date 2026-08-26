package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, arg ...string) error {
	pokemonName := arg[0]
	_, ok := cfg.pokedex[pokemonName]
	if !ok {
		URL := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
		value, err := cfg.pokeClient.GettingPokemonInfo(&URL)
		if err != nil {
			return fmt.Errorf("Error getting pokemon info api: %v", err)
		}
		cfg.pokedex[pokemonName] = value
	}
	pokemonInfo := cfg.pokedex[pokemonName]

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)

	catchBase := rand.IntN(100)
	baseExp := pokemonInfo.Base_Experience
	if rand.IntN(baseExp) > catchBase {
		fmt.Printf("%v escaped!\n", pokemonName)
		return nil
	}
	fmt.Printf("%v was caught!\n", pokemonName)
	return nil
}
