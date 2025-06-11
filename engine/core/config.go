package core

import (
	"gopkg.in/yaml.v3"
	"os"
)

type WindowConfig struct {
	Title      string `yaml:"title"`
	Width      int    `yaml:"width"`
	Height     int    `yaml:"height"`
	Fullscreen bool   `yaml:"fullscreen"`
}

type Config struct {
	Window WindowConfig `yaml:"window"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
