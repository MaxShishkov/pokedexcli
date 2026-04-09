package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MaxShishkov/pokedexcli/internal/pokeapi"
)

type requestConfig struct {
	pokeapiClient pokeapi.Client
	next          *string
	previouse     *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*requestConfig) error
}

func startRepl(cfg *requestConfig) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		cmdName := words[0]

		cmd, ok := getCommands()[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(cfg)
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
