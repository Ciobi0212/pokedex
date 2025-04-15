package main

import (
	"fmt"
	"strings"
)

func cleanInput(textInput string) []string {
	textInput = strings.ToLower(textInput)
	split := strings.Fields(textInput)

	return split
}

func main() {
	fmt.Println("Hello, World!")
}
