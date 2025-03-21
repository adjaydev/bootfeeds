package config

import (
	"encoding/json"
	"log"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(u string) {
	c.CurrentUserName = u

	content, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Printf("Error writing json: %s", err)
	}

	path, err := getConfigFilePath()
	if err != nil {
		log.Printf("Error reading path: %s", err)
	}

	os.WriteFile(path+"/"+configFileName, content, 0644)
}

func Read() *Config {
	path, err := getConfigFilePath()
	if err != nil {
		log.Printf("Error reading path: %s", err)
	}

	content, err := os.ReadFile(path + "/" + configFileName)
	if err != nil {
		log.Printf("Error reading file: %s", err)
	}

	var cfg Config
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		log.Printf("error unmarshaling JSON: %s", err)
	}

	return &cfg
}

func getConfigFilePath() (string, error) {
	pwd, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return pwd, nil
}
