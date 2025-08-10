package config

import (
	"errors"
	"os"
	"testing"
)

func configFileSetup(t *testing.T, content string) string {
	t.Helper()
	tmpFile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer tmpFile.Close()

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}

	return tmpFile.Name()
}

func configFileCleanup(t *testing.T, filePath string) {
	t.Helper()
	if err := os.Remove(filePath); err != nil {
		t.Fatalf("failed to remove temp file: %v", err)
	}
}

func TestLoadConfigReturnsErrorWhenFilePathIsEmpty(t *testing.T) {
	_, err := LoadConfig("")
	if !errors.Is(err, ErrNoConfigFile) {
		t.Errorf("expected error %v, got %v", ErrNoConfigFile, err)
	}
}

func TestLoadConfigReturnsErrorWhenFileDoesNotExist(t *testing.T) {
	_, err := LoadConfig("nonexistent.yaml")
	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func TestLoadConfigParsesValidConfigFile(t *testing.T) {
	content := `
server:
  port: 8080
  host: "localhost"
logger:
  level: "info"
  format: "json"
`
	filePath := configFileSetup(t, content)
	defer configFileCleanup(t, filePath)

	config, err := LoadConfig(filePath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if config.ServerConfig.Port != 8080 {
		t.Errorf("expected port 8080, got %d", config.ServerConfig.Port)
	}
	if config.ServerConfig.Host != "localhost" {
		t.Errorf("expected host 'localhost', got %s", config.ServerConfig.Host)
	}
	if config.LoggerConfig.Level != "info" {
		t.Errorf("expected level 'info', got %s", config.LoggerConfig.Level)
	}
	if config.LoggerConfig.Format != "json" {
		t.Errorf("expected format 'json', got %s", config.LoggerConfig.Format)
	}
}

func TestLoadConfigReturnsErrorForInvalidYAML(t *testing.T) {
	content := `
server:
  port: "not-a-number"
`
	filePath := configFileSetup(t, content)
	defer configFileCleanup(t, filePath)

	_, err := LoadConfig(filePath)
	if err == nil {
		t.Error("expected an error, got nil")
	}
}
