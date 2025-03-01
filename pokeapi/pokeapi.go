package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Config struct {
	Next     *string
	Previous *string
}

func CreateConfig() *Config {
	return &Config{Next: toPtr("https://pokeapi.co/api/v2/location-area/"), Previous: nil}
}

func toPtr(str string) *string {
	return &str
}

func GetArea(cfg *Config) (*LocationArea, error) {
	if cfg.Next == nil {
		return nil, fmt.Errorf("No Next URL")
	}
	req, err := http.Get(*cfg.Next)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var response LocationArea
	if err = json.Unmarshal(data, &response); err != nil {
		return nil, err
	}
	cfg.Previous = response.Previous
	cfg.Next = response.Next

	return &response, nil

}
