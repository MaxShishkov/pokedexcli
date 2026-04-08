package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func startRepl() {
	cmdMap := make(map[string]cliCommand)

	cmdMap["exit"] = cliCommand{
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    commandExit,
	}

	cmdMap["help"] = cliCommand{
		Name:        "help",
		Description: "Show this help message",
		Callback:    func() error { return commandHelp(cmdMap) },
	}

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

		cmd, ok := cmdMap[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		cmd.Callback()
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	return strings.Fields(output)
}
