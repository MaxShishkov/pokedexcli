package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *requestConfig, params ...string) error {
	if len(cfg.caughtPokemons) == 0 {
		return errors.New("you have not caught any pokemons yet")
	}

	fmt.Println("Your Pokedex:")
	for name := range cfg.caughtPokemons {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
