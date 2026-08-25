package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, value := range cfg.commandList {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	return nil
}
