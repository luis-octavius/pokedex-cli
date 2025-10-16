package main 

import (
	"fmt"
	"strings"
) 

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)

	if len(trimmed) == 0 {
		return []string{}
	}

	splittedWords := strings.Split(trimmed, " ")

	return splittedWords 
}
