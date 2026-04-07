package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MaxShishkov/pokedexcli/repl"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		clenadText := repl.CleanInput(text)
		fmt.Printf("Your command was: %v\n", clenadText[0])
	}
}
