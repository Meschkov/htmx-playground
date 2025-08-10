package config

import (
	"errors"
	"github.com/goccy/go-yaml"
	"os"
)

var ErrNoConfigFile = errors.New("no config file provided")

// Config represents the application configuration structure.
type Config struct {
	ServerConfig *serverConfig `yaml:"server"`
	LoggerConfig *loggerConfig `yaml:"logger"`
}

// LoadConfig loads the configuration from a YAML file specified by filePath.
func LoadConfig(filePath string) (*Config, error) {
	if filePath == "" {
		return nil, ErrNoConfigFile
	}
	// Read the YAML file content.
	f, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Load the YAML configuration from the file.
	var c Config
	if err := yaml.Unmarshal(f, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

type serverConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type loggerConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"` // "text" | "json"
}
