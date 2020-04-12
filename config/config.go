package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Persons       []Person `yaml:"persons"`
	OAuthClientID string
}

type Person struct {
	ID    string  `yaml:"id"`
	Name  string  `yaml:"name"`
	Ratio float64 `yaml:"ratio"`
}

func Load(filepath string) (*Config, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var conf Config
	if err := yaml.NewDecoder(f).Decode(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
