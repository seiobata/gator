package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl string `json:"db_url"`
	User  string `json:"current_user_name"`
}

func (cfg *Config) SetUser(user string) error {
	cfg.User = user
	dat, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	os.WriteFile(filePath, dat, 0644)
	return nil
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}
	var cfg Config
	if err = json.Unmarshal(dat, &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configFileName), nil
}
