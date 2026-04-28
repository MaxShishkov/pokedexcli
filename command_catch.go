package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *requestConfig, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("Please provide a pokemon name")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", params[0])

	pokemon, err := cfg.pokeapiClient.GetPokemon(params[0])
	if err != nil {
		fmt.Println("Failed to get pokemon data:", err)
		return err
	}

	minChance := 0.1
	maxchance := 0.9
	scale := 150.0

	chance := minChance + (maxchance-minChance)*(float64(pokemon.BaseExperience)/scale)

	if rand.Float64() < chance {
		fmt.Printf("Gotcha! %s was caught!\n", params[0])
		cfg.caughtPokemons[params[0]] = pokemon
	} else {
		fmt.Printf("Oh no, %s escaped!\n", params[0])
	}

	return nil
}
