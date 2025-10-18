package main

import (
	"fmt"
	"os"
	"github.com/luis-octavius/pokedex-cli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(args []string) error
	config      *config
}

type config struct {
	Next     string
	Previous string
}

func commandExit(c *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	commands := getCommands(c)
	for _, c := range commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}

	fmt.Println()
	return nil
}

func commandMap(c *config, args []string) error {
	url := c.Next 
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area?limit=20"
	}

	page := pokeapi.GetLocations(url)
	printLocations(page)

	c.Next = page.Next 
	c.Previous = page.Previous 
	return nil
}

func commandMapBack(c *config, args []string) error {
	if c.Previous == "" {
		fmt.Println("you're on the first page")
		return nil 
	}

	page := pokeapi.GetLocations(c.Previous)
	printLocations(page)
	c.Next = page.Next 
	c.Previous = page.Previous
	return nil
}

func getCommands(c *config) map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Shows a help message",
			callback:    func (args []string) error {
				return commandHelp(c, args)
			},
			config:      c,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func (args []string) error {
				return commandExit(c, args)
			},
			config:      c,
		},
		"map": {
			name:        "map",
			description: "Displays the next map locations",
			config:      c,
			callback:    func (args []string) error {
				return commandMap(c, args)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous map locations",
			callback:    func (args []string) error {
				return commandMapBack(c, args)
			},
			config:      c,
		},
	}
}

func printLocations(locations pokeapi.LocationOffset) {
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
}
