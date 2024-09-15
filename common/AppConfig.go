package common

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	Port int `json:"port"`
	Name string `json:"name"`
	Version string `json:"version"`
}

const FileName string = "configs.json"
var ApplicationConfig = AppConfig{
	Port: 8080,
	Name: "lecter",
	Version: "beta",
}

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