package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/shaneplunkett/godex/internal/pokecache"
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

func GetArea(cfg *Config, c *pokecache.Cache) (*LocationArea, error) {
	if cfg.Next == nil {
		return nil, fmt.Errorf("No Next URL")
	}

	//cache check logic
	val, ok := c.Get(*cfg.Next)
	if ok {
		var response LocationArea
		err := json.Unmarshal(val, &response)
		if err != nil {
			return &response, nil
		}
	}

	// none cache logic
	req, err := http.Get(*cfg.Next)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	// Add to cache
	c.Add(*cfg.Next, data)

	var response LocationArea
	if err = json.Unmarshal(data, &response); err != nil {
		return nil, err
	}
	cfg.Previous = response.Previous
	cfg.Next = response.Next

	return &response, nil

}
