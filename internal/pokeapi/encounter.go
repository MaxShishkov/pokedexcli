package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"errors"
)

func (c *Client) Encounters(locationName string) (ShallowPokemonEncounters, error) {
	if locationName == "" {
		return ShallowPokemonEncounters{}, errors.New("location name must be specified")
	}

	url := baseURL + "location-area/" + locationName

	var body []byte

	if entry, exists := c.cache.Get(url); exists {
		body = entry
	} else {
		res, err := http.Get(url)
		if err != nil {
			return ShallowPokemonEncounters{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return ShallowPokemonEncounters{}, err
		}
		c.cache.Add(url, body)
	}

	encounterRes := ShallowPokemonEncounters{}
	err := json.Unmarshal(body, &encounterRes)
	if err != nil {
		return ShallowPokemonEncounters{}, err
	}

	return encounterRes, nil
}