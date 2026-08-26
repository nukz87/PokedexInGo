package main

import (
	"github.com/nukz87/PokedexInGo/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callBack    func(*config, ...string) error
}

type config struct {
	pokeClient          pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
	commandList         map[string]cliCommand
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callBack:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Print manual",
			callBack:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Print a list of location areas/ Next page of the list",
			callBack:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Print a list of location areas/ Previous page of the list",
			callBack:    commandMapb,
		},
		"explore": {
			name:        "explore 'location-area-name'",
			description: "Print a list of pokemons found in 'location-area-name'",
			callBack:    commandExplore,
		},
	}
}
