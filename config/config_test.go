package config

import (
	"os"
	"testing"
)

func TestEnvConfig(t *testing.T) {
	// Set environment variables
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("DATABASE_PATH", "/tmp/test.db")
	os.Setenv("API_KEY", "test-key")

	// Load config
	cfg := Load()

	// Check values
	if cfg.Server.Port != "8080" {
		t.Errorf("Expected port 8080, got %s", cfg.Server.Port)
	}

	if cfg.Database.Path != "/tmp/test.db" {
		t.Errorf("Expected database path /tmp/test.db, got %s", cfg.Database.Path)
	}

	if cfg.APIKey != "test-key" {
		t.Errorf("Expected API key test-key, got %s", cfg.APIKey)
	}

	// Clean up
	os.Unsetenv("SERVER_PORT")
	os.Unsetenv("DATABASE_PATH")
	os.Unsetenv("API_KEY")
}
