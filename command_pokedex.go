package main

import "fmt"

func commandPokedex(cfg *config, arg ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("Your Pokedex is empty.")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for key, _ := range cfg.pokedex {
		fmt.Printf(" - %v\n", key)
	}

	return nil
}
