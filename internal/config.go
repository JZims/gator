package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configFileName), nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encoder
	encoder := json.NewEncoder(file)

	// Writer
	if err := encoder.Encode(cfg); err != nil {
		return err
	}
	return nil
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {

	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	configFile, err := os.Open(configPath)
	if err != nil {
		return Config{}, err
	}
	defer configFile.Close()
	parsedJson := Config{}
	decoder := json.NewDecoder(configFile)

	if err := decoder.Decode(&parsedJson); err != nil {
		return Config{}, err
	}

	return parsedJson, nil
}

func (c *Config) SetUser(userName string) error {

	c.CurrentUserName = userName

	if err := write(*c); err != nil {
		return err
	}

	return nil
}
