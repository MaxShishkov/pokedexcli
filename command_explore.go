package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *requestConfig, params ...string) error {
	if len(params) == 0 {
		return errors.New("Location name was not specified.")
	}
	if len(params) != 1 {
		return errors.New("Exactly one location must be specified")
	}

	encountersRes, err := cfg.pokeapiClient.Encounters(params[0])
	if err != nil {
		return err
	}

	for _, encounter := range encountersRes.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}