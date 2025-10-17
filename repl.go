package main 

import (
	"strings"
	"fmt"
	"bufio"
	"os"
) 

type cliCommand struct {
	name        string 
	description string 
	callback		func() error
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		
		scanner.Scan()			
		words := cleanInput(scanner.Text())

		commandWord := words[0]

		cliCommands := getCommands()

		if commandWord == cliCommands[commandWord].name {
			cliCommands[commandWord].callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}


