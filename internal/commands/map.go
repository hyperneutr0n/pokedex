package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func getMap(cfg *Config) error {
	url := cfg.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Couldn't reach the API. %w", err)
	}
	defer res.Body.Close()

	var response struct {
		Count    int            `json:"count"`
		Next     string         `json:"next"`
		Previous string         `json:"previous"`
		Result   []LocationArea `json:"results"`
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&response); err != nil {
		return fmt.Errorf("Error parsing json. %w", err)
	}

	for _, area := range response.Result {
		fmt.Println(area.Name)
	}

	cfg.Next = response.Next
	cfg.Previous = response.Previous

	return nil
}

func getMapBack(cfg *Config) error {
	url := cfg.Previous
	if url == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Couldn't reach the API. %w", err)
	}
	defer res.Body.Close()

	var response struct {
		Count    int            `json:"count"`
		Next     string         `json:"next"`
		Previous string         `json:"previous"`
		Result   []LocationArea `json:"results"`
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&response); err != nil {
		return fmt.Errorf("Error parsing json. %w", err)
	}

	for _, area := range response.Result {
		fmt.Println(area.Name)
	}

	cfg.Next = response.Next
	cfg.Previous = response.Previous

	return nil
}