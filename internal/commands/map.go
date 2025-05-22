package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func getMap(args []string, cfg *Config) error {
	url := cfg.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	var response struct {
		Count    int            `json:"count"`
		Next     string         `json:"next"`
		Previous string         `json:"previous"`
		Result   []LocationArea `json:"results"`
	}

	// Check for cache
	key := url
	cached, exist := cfg.Cache.Get(key)
	if exist {
		if err := json.Unmarshal(cached, &response); err != nil {
			return fmt.Errorf("Error unmarshalling JSON from cache. %w", err)
		}

		return printLocationNames(response.Result, cfg, response.Next, response.Previous)
	}

	// Cache doesn't exist, call the API
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Couldn't reach the API. %w", err)
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading response body. %w", err)
	}

	// Create a cache
	cfg.Cache.Add(key, body)

	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("Error unmarshalling JSON data. %w", err)
	}

	cfg.Next = response.Next
	cfg.Previous = response.Previous

	return printLocationNames(response.Result, cfg, response.Next, response.Previous)
}

func getMapBack(args []string, cfg *Config) error {
	url := cfg.Previous
	if url == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	var response struct {
		Count    int            `json:"count"`
		Next     string         `json:"next"`
		Previous string         `json:"previous"`
		Result   []LocationArea `json:"results"`
	}

	// Check for cache
	key := url
	cached, exist := cfg.Cache.Get(key)
	if exist {
		if err := json.Unmarshal(cached, &response); err != nil {
			return fmt.Errorf("Error unmarshalling JSON from cache. %w", err)
		}

		return printLocationNames(response.Result, cfg, response.Next, response.Previous)
	}

	// Cache doesn't exist, call the API
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Couldn't reach the API. %w", err)
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading response body. %w", err)
	}

	// Create a cache
	cfg.Cache.Add(key, body)

	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("Error unmarshalling JSON data. %w", err)
	}

	return printLocationNames(response.Result, cfg, response.Next, response.Previous)
}

func printLocationNames(locationAreas []LocationArea, cfg *Config, next, previous string) error {
	for _, area := range locationAreas {
		fmt.Println(area.Name)
	}
	cfg.Next = next
	cfg.Previous = previous
	return nil
}
