package commands

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/Ciobi0212/pokedex/internal/apipokeinteraction"
	"github.com/Ciobi0212/pokedex/internal/models"
	"github.com/Ciobi0212/pokedex/internal/pokecache"
	"github.com/Ciobi0212/pokedex/internal/utils"
	"gopkg.in/yaml.v3"
)

type AppState struct {
	mapPagOffset  int
	cache         *pokecache.Cache
	caughtPokemon map[string]*models.Pokemon
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*AppState, []string) error
}

func commandExitCallback(state *AppState, params []string) error {
	if len(params) != 0 {
		return fmt.Errorf("exit command doesn't support any params")
	}

	fmt.Println("Closing the Pokedex... Goodbye!")

	state.cache.StopReaping()

	time.Sleep(time.Millisecond * 100)

	os.Exit(0)

	return nil
}

func commandHelpCommand(state *AppState, params []string) error {

	if len(params) != 0 {
		return fmt.Errorf("help command doesn't support any params")
	}

	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Here are the available commands:")
	fmt.Println("")

	availableCommands := GetCommandMap()

	commandNames := make([]string, 0, len(availableCommands))
	for name := range availableCommands {
		commandNames = append(commandNames, name)
	}
	sort.Strings(commandNames)

	for _, name := range commandNames {
		command := availableCommands[name]

		fmt.Printf("%-10s : %s\n", command.Name, command.Description)

	}
	fmt.Println("")

	return nil
}

func commandMapCallback(state *AppState, params []string) error {
	if len(params) != 0 {
		return fmt.Errorf("map command doesn't support any params")
	}

	state.mapPagOffset += 20

	mapNames, err := apipokeinteraction.GetLocationAreas(state.mapPagOffset, state.cache)
	if err != nil {
		return fmt.Errorf("error getting location areas: %w", err)
	}

	for _, name := range mapNames {
		fmt.Println(name)
	}

	return nil
}

func commandMapbCallback(state *AppState, params []string) error {
	if len(params) != 0 {
		return fmt.Errorf("mapb command doesn't support any params")
	}
	state.mapPagOffset -= 20

	if state.mapPagOffset < -20 {
		fmt.Println("you can't go back to any pages, run map command first")
		return nil
	}

	if state.mapPagOffset == -20 {
		fmt.Println("you are on the first page")
		return nil
	}

	mapNames, err := apipokeinteraction.GetLocationAreas(state.mapPagOffset, state.cache)

	if err != nil {
		return fmt.Errorf("error getting areas %w", err)
	}

	for _, name := range mapNames {
		fmt.Println(name)
	}

	return nil
}

func commandExploreCallback(state *AppState, params []string) error {
	if len(params) != 1 {
		return fmt.Errorf("explore command only supports 1 param")
	}

	areaName := params[0]

	pokemonsName, err := apipokeinteraction.GetPokemonsFromArea(areaName, state.cache)
	if err != nil {
		return fmt.Errorf("error getting pokemons: %w", err)
	}

	for _, name := range pokemonsName {
		fmt.Println(name)
	}

	return nil
}

func commandCatchCallback(state *AppState, params []string) error {
	if len(params) != 1 {
		return fmt.Errorf("catch command only supports 1 param")
	}

	pokemonName := params[0]

	pokemonInfo, err := apipokeinteraction.GetPokemonInfo(pokemonName, state.cache)

	if err != nil {
		return fmt.Errorf("error getting pokemon base exp: %w", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	catchChance := utils.CalculateCatchChanceTiered(float64(pokemonInfo.BaseExp))

	fmt.Printf("Your chance of catching it is: %v percent\n", catchChance*100)

	if rand.Float64() < catchChance {
		fmt.Printf("YOU CAUGHT %s !\n", pokemonName)
		state.caughtPokemon[pokemonName] = pokemonInfo
	} else {
		fmt.Printf("Upsy, %s escaped !\n", pokemonName)
	}

	return nil
}

func commandInspectCallback(state *AppState, params []string) error {
	if len(params) != 1 {
		return fmt.Errorf("inspect command only supports 1 param")
	}

	pokemonName := params[0]

	pokemonInfo, ok := state.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("you didn't caught this pokemon")
	}

	yamlBytes, err := yaml.Marshal(pokemonInfo)

	if err != nil {
		return fmt.Errorf("error converting pokemon info to yaml %w", err)
	}

	fmt.Println(string(yamlBytes))

	return nil
}

func commandPokedexCallback(state *AppState, params []string) error {
	if len(params) != 0 {
		return fmt.Errorf("pokedex command doesn't support any params")
	}

	if len(state.caughtPokemon) == 0 {
		fmt.Println("No pokemons in pokedex, better start catching!")
		return nil
	}

	fmt.Println("Your pokedex: ")
	for _, pokemonInfo := range state.caughtPokemon {
		fmt.Printf("- %s\n", pokemonInfo.Name)
	}

	return nil
}

func GetCommandMap() map[string]CliCommand {
	commandMap := map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exits the Pokedex application",
			Callback:    commandExitCallback,
		},
		"help": {
			Name:        "help",
			Description: "Displays this help message",
			Callback:    commandHelpCommand,
		},
		"map": {
			Name:        "map",
			Description: "Displays the next page of location areas",
			Callback:    commandMapCallback,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous page of location areas",
			Callback:    commandMapbCallback,
		},
		"explore": {
			Name: "explore",

			Description: "Lists the Pokémon available in a specific location area. Usage: explore <location_area_name>",
			Callback:    commandExploreCallback,
		},
		"catch": {
			Name:        "catch",
			Description: "Attempts to catch a specified Pokémon found in an area. Usage: catch <pokemon_name>",
			Callback:    commandCatchCallback,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Shows details of a Pokémon you have already caught. Usage: inspect <pokemon_name>",
			Callback:    commandInspectCallback,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Lists all the Pokémon currently in your Pokedex (that you have caught)",
			Callback:    commandPokedexCallback,
		},
	}

	return commandMap
}

func GetInitAppState() *AppState {
	state := AppState{
		mapPagOffset:  -20,
		cache:         pokecache.NewCache(time.Millisecond * 10000),
		caughtPokemon: make(map[string]*models.Pokemon),
	}

	return &state
}
