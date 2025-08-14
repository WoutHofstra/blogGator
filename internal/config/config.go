package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"fmt"
)

type Config struct {
	DbURL			string		`json:"db_url"`
	CurrentUserName		string		`json:"current_user_name"`
}

func Read() (Config, error) {
	homeDir, _ := os.UserHomeDir()

	filePath := filepath.Join(homeDir, "/.gatorconfig.json")
	file, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("Could not read file: %w", err)
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return Config{}, fmt.Errorf("Could not unmarshal file: %w", err)
	}
	return config, nil
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name

	bytes, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("Could not marshal data: %w", err)
	}

        homeDir, _ := os.UserHomeDir()

        filePath := filepath.Join(homeDir, "/.gatorconfig.json")

	err = os.WriteFile(filePath, bytes, 0644)
	if err != nil {
		return fmt.Errorf("Could not write file: %w", err)
	}
	return nil
}
