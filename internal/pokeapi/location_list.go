package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (ShallowLocationAreaResponse, error) {
	url := baseURL + "location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	res, err := http.Get(url)
	if err != nil {
		return ShallowLocationAreaResponse{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ShallowLocationAreaResponse{}, err
	}

	locationResp := ShallowLocationAreaResponse{}
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return ShallowLocationAreaResponse{}, err
	}

	return locationResp, nil

}
