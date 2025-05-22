package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func explore(args []string, cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if args[0] == "" {
		return fmt.Errorf("No location area passed as an argument.")
	}
	url = url + args[0]

	var response struct {
		EncounterMethodRates []struct {
			EncounterMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"encounter_method"`
			VersionDetails []struct {
				Rate    int `json:"rate"`
				Version struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"version"`
			} `json:"version_details"`
		} `json:"encounter_method_rates"`
		GameIndex int `json:"game_index"`
		ID        int `json:"id"`
		Location  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"location"`
		Name  string `json:"name"`
		Names []struct {
			Language struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"language"`
			Name string `json:"name"`
		} `json:"names"`
		PokemonEncounters []struct {
			Pokemon        Pokemon `json:"pokemon"`
			VersionDetails []struct {
				EncounterDetails []struct {
					Chance          int   `json:"chance"`
					ConditionValues []any `json:"condition_values"`
					MaxLevel        int   `json:"max_level"`
					Method          struct {
						Name string `json:"name"`
						URL  string `json:"url"`
					} `json:"method"`
					MinLevel int `json:"min_level"`
				} `json:"encounter_details"`
				MaxChance int `json:"max_chance"`
				Version   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"version"`
			} `json:"version_details"`
		} `json:"pokemon_encounters"`
	}

	key := url
	cached, exist := cfg.Cache.Get(key)
	if exist {
		if err := json.Unmarshal(cached, &response); err != nil {
			return fmt.Errorf("Error unmarshalling JSON from cache. %w", err)
		}

		for _, encounter := range response.PokemonEncounters {
			fmt.Println(encounter.Pokemon.Name)
		}
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Couldn't reach the API. %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading response body. %w", err)
	}

	cfg.Cache.Add(key, body)

	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("Error unmarshalling JSON data. %w", err)
	}

	for _, encounter := range response.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
