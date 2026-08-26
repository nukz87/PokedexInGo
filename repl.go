package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$ Pokedex > ")
		if scanner.Scan() {
			text := cleanInput(scanner.Text())
			command := text[0]
			switch command {
			case "help":
				if err := commandHelp(cfg); err != nil {
					fmt.Printf("Error help command: %v\n", err)
				}
			case "exit":
				if err := commandExit(cfg); err != nil {
					fmt.Printf("Error exit command: %v\n", err)
				}
			case "map":
				if err := commandMap(cfg); err != nil {
					fmt.Printf("Error map command: %v\n", err)
				}
			case "mapb":
				if err := commandMapb(cfg); err != nil {
					fmt.Printf("Error mapb command: %v\n", err)
				}
			case "explore":
				if len(text) < 2 {
					fmt.Println("explore Usage: explore 'location-area-name'")
					continue
				}
				if err := commandExplore(cfg, text[1]); err != nil {
					fmt.Printf("Error explore command: %v\n", err)
				}
			case "catch":
				if len(text) < 2 {
					fmt.Println("catch Usage: catch 'pokemon-name'")
					continue
				}
				if err := commandCatch(cfg, text[1]); err != nil {
					fmt.Printf("Error catch command: %v\n", err)
				}
			case "inspect":
				if len(text) < 2 {
					fmt.Println("inspect Usage: inspect 'pokemon-name'")
					continue
				}
				if err := commandInspect(cfg, text[1]); err != nil {
					fmt.Printf("Error inspect command: %v\n", err)
				}
			case "pokedex":
				if err := commandPokedex(cfg); err != nil {
					fmt.Printf("Error pokedex command: %v\n", err)
				}
			default:
				fmt.Println("Unknown command")
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Data reading error: ", err)
		}
	}
}
