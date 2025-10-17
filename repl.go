package main 

import (
	"strings"
	"fmt"
	"bufio"
	"os"
) 

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		
		scanner.Scan()			
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue 
		}

		commandName := words[0]

		fmt.Println("Your command was:", commandName)
	}
}
func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}
