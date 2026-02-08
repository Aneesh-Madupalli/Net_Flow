package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config holds application settings (reserved for future use)
type Config struct{}

// Default returns the default configuration
func Default() Config {
	return Config{}
}

// configPath returns the path to the config file
func configPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	appDir := filepath.Join(dir, "NetFlow")
	if err := os.MkdirAll(appDir, 0700); err != nil {
		return "", err
	}
	return filepath.Join(appDir, "config.json"), nil
}

// Load reads config from disk, or returns default if not found/invalid
func Load() (Config, error) {
	path, err := configPath()
	if err != nil {
		return Default(), err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Default(), nil
		}
		return Default(), err
	}
	var c Config
	if err := json.Unmarshal(data, &c); err != nil {
		return Default(), nil
	}
	return c, nil
}

// Save writes config to disk
func Save(c Config) error {
	path, err := configPath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}
