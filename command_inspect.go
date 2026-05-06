package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *requestConfig, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("please provide a pokemon name")
	}

	pokemon, caught := cfg.caughtPokemons[params[0]]
	if !caught {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}
