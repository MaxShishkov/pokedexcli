package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "pokemon/" + name
	fmt.Println(url)

	var body []byte

	if entry, exists := c.cache.Get(url); exists {
		body = entry
	} else {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}
		c.cache.Add(url, body)
	}

	pokemonResp := Pokemon{}
	err := json.Unmarshal(body, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonResp, nil
}
