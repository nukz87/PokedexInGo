package main

import (
	"time"

	"github.com/nukz87/PokedexInGo/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeClient:          pokeapi.NewClient(5 * time.Second),
		nextLocationURL:     new(string),
		previousLocationURL: new(string),
		commandList:         getCommand(),
	}
	startREPL(cfg)
}
