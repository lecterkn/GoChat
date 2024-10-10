package common

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	Port    int    `json:"port"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

// configファイル
const FileName string = "build/package/configs.json"

// デフォルトのAppConfig
var ApplicationConfig = AppConfig{
	Port:    8080,
	Name:    "lecter",
	Version: "beta",
}

/*
 * configファイルから設定を読み込む
 */
func LoadConfig() *AppConfig {
	file, err := os.ReadFile(FileName)
	if err != nil {
		return nil
	}
	var appConfig AppConfig
	err = json.Unmarshal(file, &appConfig)
	if err != nil {
		return nil
	}
	return &appConfig
}
