package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Ciobi0212/pokedex/commands"
)

func cleanInput(textInput string) []string {
	textInput = strings.ToLower(textInput)
	split := strings.Fields(textInput)

	return split
}

func main() {
	commandMap := commands.GetCommandMap()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		// Wait for user input
		scanner.Scan()

		// Read user input
		textInput := scanner.Text()

		// Clean input
		inputSlice := cleanInput(textInput)

		// Get command
		commandWord := inputSlice[0]

		// Execute command if supported
		command, wasFound := commandMap[commandWord]
		if !wasFound {
			fmt.Println("Unknown command")
			continue
		}

		command.Callback()
	}
}
