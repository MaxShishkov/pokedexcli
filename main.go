package main

import (
	"time"

	"github.com/MaxShishkov/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &requestConfig{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
