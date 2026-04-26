package main

import (
	"fmt"
)

func commandMapf(cfg *requestConfig, params ...string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = locationResp.Next
	cfg.previouse = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *requestConfig, params ...string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previouse)
	if err != nil {
		return err
	}

	cfg.next = locationResp.Next
	cfg.previouse = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
