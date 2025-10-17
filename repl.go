package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	
	var conf config
	cliCommands := getCommands(&conf)

	for {
		fmt.Print("Pokedex > ")
			
		scanner.Scan()
		words := cleanInput(scanner.Text())

		commandWord := words[0]
		args := words[:1]

		cmd, ok := cliCommands[commandWord]
		if !ok {
			fmt.Println("Uknown command")
			continue 
		}

		if err := cmd.callback(args); err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}
