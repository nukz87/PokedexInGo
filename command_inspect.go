package main

import "fmt"

func commandInspect(cfg *config, arg ...string) error {
	pokemonName := arg[0]

	pokemonInfo, ok := cfg.pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name:", pokemonInfo.Name)
	fmt.Println("Height:", pokemonInfo.Height)
	fmt.Println("Weight:", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonInfo.Stats {
		fmt.Printf("  - %v: %v\n", stat.Stat.Name, stat.Base_Stat)
	}
	fmt.Println("Types:")
	for _, t := range pokemonInfo.Types {
		fmt.Printf("  - %v\n", t.Type.Name)
	}

	return nil
}
