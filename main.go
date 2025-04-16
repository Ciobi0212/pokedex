package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Ciobi0212/pokedex/internal/commands"
	"github.com/Ciobi0212/pokedex/internal/utils"
)

func main() {
	appState := commands.GetInitAppState()

	commandMap := commands.GetCommandMap()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		// Wait for user input
		scanner.Scan()

		// Read user input
		textInput := scanner.Text()

		if len(textInput) <= 0 {
			fmt.Println("Type something you silly :)")
			continue
		}

		// Clean input
		inputSlice := utils.CleanInput(textInput)

		// Get command
		commandWord := inputSlice[0]

		// Execute command if supported
		command, wasFound := commandMap[commandWord]
		if !wasFound {
			fmt.Println("Unknown command, type help to see what you can do")
			continue
		}

		err := command.Callback(appState, inputSlice[1:])
		if err != nil {
			fmt.Println(err)
		}
	}
}
