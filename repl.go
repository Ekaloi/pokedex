package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/Ekaloi/pokedex/internal/pokeapi"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, words)
			if err != nil {

				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}


type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},"explore": {
			name:        "explore",
			description: "Explore area from map",
			callback:    commandExplore,
		},"catch": {
			name:        "catch",
			description: "catch a wild pokemon",
			callback:    commandCatch,
		},"inspect": {
			name:        "inspect",
			description: "inspect caught pokemon",
			callback:    commandInspect,
		},"pokedex": {
			name:        "pokedex",
			description: "show all caught pokemon",
			callback:    commandPokedex,
		},

	}
}