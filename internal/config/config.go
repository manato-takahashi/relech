package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Repository struct {
	Name  string `yaml:"name"`
	Owner string `yaml:"owner"`
	Base  string `yaml:"base"` // main or master
	Head  string `yaml:"head"` // develop
}

type PRTemplate struct {
	Title string `yaml:"title"` // e.g. "本番リリース({{.Date}})"
}

type Config struct {
	Repositories []Repository `yaml:"repositories"`
	PRTemplate   PRTemplate   `yaml:"pr_template"`
}

func Load(path string) (*Config, error) {
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
