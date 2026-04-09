package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/MaxShishkov/pokedexcli/internal/map_helper"
)

func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var locationAreasResponse map_helper.LocationAreaResponse
	err = json.Unmarshal(body, &locationAreasResponse)
	if err != nil {
		log.Fatal(err)
	}

	for _, locationArea := range locationAreasResponse.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}
