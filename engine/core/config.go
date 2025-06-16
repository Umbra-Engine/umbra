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

type GameConfig struct {
	TargetFPS int `yaml:"target_fps"`
}

type Config struct {
	Window WindowConfig `yaml:"window"`
	Game   GameConfig   `yaml:"game"`
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

	// Default to 60 FPS if no specification
	if cfg.Game.TargetFPS <= 0 {
		cfg.Game.TargetFPS = 60
	}

	return &cfg, nil
}
