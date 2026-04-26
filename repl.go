package main

import (
	"fmt"
	"strings"

	"github.com/MaxShishkov/pokedexcli/internal/pokeapi"
	"github.com/peterh/liner"
)

type requestConfig struct {
	pokeapiClient pokeapi.Client
	next          *string
	previouse     *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*requestConfig, ...string) error
}

func startRepl(cfg *requestConfig) {
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)

	for {
		input, err := line.Prompt("Pokedex > ")
		if err != nil {
			break
		}
		line.AppendHistory(input)

		words := cleanInput(input)
		if len(words) == 0 {
			continue
		}
		cmdName := words[0]
		params := words[1:]

		cmd, ok := getCommands()[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err = cmd.callback(cfg, params...)
		if err != nil {
			fmt.Println("Error occurred:", err)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	return strings.Fields(output)
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
		"explore": {
			name:			"explore",
			description:	"Get the list of pokemon in the area",
			callback:		commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
