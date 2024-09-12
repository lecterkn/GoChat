package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	Port int `json:"port"`
}

const FileName string = "configs.json"

func LoadConfig() *AppConfig {
	file, err := os.ReadFile(FileName)
	if err != nil {
		return nil
	}
	var appConfig AppConfig
	err = json.Unmarshal(file, &appConfig)
	if (err != nil) {
		return nil
	}
	return &appConfig
}