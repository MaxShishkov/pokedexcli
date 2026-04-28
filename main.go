package main

import (
	"time"

	"github.com/MaxShishkov/pokedexcli/internal/pokeapi"
	"github.com/MaxShishkov/pokedexcli/internal/pokecache"
)

func main() {
	const baseTimeout = 10 * time.Second
	cache := pokecache.NewCache(7 * time.Second)
	pokeClient := pokeapi.NewClient(baseTimeout, cache)
	cfg := &requestConfig{
		caughtPokemons: make(map[string]pokeapi.Pokemon),
		pokeapiClient:  pokeClient,
	}
	startRepl(cfg)
}
