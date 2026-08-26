package main

import (
	"time"

	"github.com/nukz87/PokedexInGo/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokedex:             make(map[string]pokeapi.Pokemon),
		pokeClient:          pokeapi.NewClient(5 * time.Second),
		nextLocationURL:     new(string),
		previousLocationURL: new(string),
		commandList:         getCommand(),
	}
	startREPL(cfg)
}
