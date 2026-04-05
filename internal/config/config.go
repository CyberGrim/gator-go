package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, configFileName), nil
}

func Read() (Config, error) {
	configFile, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonFile, err := os.Open(configFile)
	if err != nil {
		return Config{}, err
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func write(cfg Config) error {
	configFile, err := getConfigFilePath()
	if err != nil {
		return err
	}

	createdfile, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer createdfile.Close()

	encoder := json.NewEncoder(createdfile)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	err := write(*cfg)
	return err
}
