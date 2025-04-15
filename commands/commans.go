package commands

import (
	"fmt"
	"os"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func commandExitCallback() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelpCommand() error {
	helpString := `Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`

	fmt.Println(helpString)

	return nil
}

func GetCommandMap() map[string]CliCommand {
	commandMap := map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exits the program",
			Callback:    commandExitCallback,
		},

		"help": {
			Name:        "help",
			Description: "Provides the user with guidance",
			Callback:    commandHelpCommand,
		},
	}

	return commandMap
}
