package main

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

// Config holds the configuration details.
type Config struct {
	Port     int      `yaml:"port"`
	Backends []string `yaml:"backends"`
	Strategy string   `yaml:"strategy"`
}

// LBConfig reads and validates the load balancer configuration from a YAML file.
func LBConfig() (*Config, error) {
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err // Early return if file read fails.
	}

	var config Config
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		return nil, err // Early return if YAML unmarshalling fails.
	}

	// Combine validation checks for better readability and efficiency.
	if len(config.Backends) == 0 || config.Port == 0 {
		return nil, errors.New("configuration error: backend hosts expected, none provided; load balancer port not found")
	}

	return &config, nil
}
