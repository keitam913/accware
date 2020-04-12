package config

import (
	"encoding/json"
	"os"
	"strings"
)

type Config struct {
	Persons []Person `json:"persons"`
}

type Person struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Ratio float64 `json:"ratio"`
}

func Load() (*Config, error) {
	var ps []Person
	if err := json.NewDecoder(strings.NewReader(os.Getenv("PERSONS"))).Decode(&ps); err != nil {
		return nil, err
	}
	return &Config{
		Persons: ps,
	}, nil
}
