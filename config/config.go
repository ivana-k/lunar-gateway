package config

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Services map[string]string                             `yaml:"services"`
	Gateway  Gateway                                       `yaml:"gateway"`
	Groups   map[string]map[string]map[string]MethodConfig `yaml:"groups"`
}

type Gateway struct {
	Route string `yaml:"route"`
	Port  string `yaml:"port"`
}

type MethodConfig struct {
	MethodRoute string `yaml:"method_route"`
	Type        string `yaml:"type"`
	Service     string `yaml:"service"`
}

func LoadConfig() (*Config, error) {
	path := "config.yml"

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var conf Config
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}

// loading methods that should pass withput token
func LoadNoAuthConfig() (*Config, error) {
	path := "no_auth_config.yml"

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var conf Config
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
