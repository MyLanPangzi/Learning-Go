package config

import (
	"encoding/json"
	"os"
)

type ExportConfig struct {
	Arn    string `json:"arn"`
	Db     string `json:"db"`
	Tables []struct {
		Regexp string `json:"regexp"`
		Table  string `json:"table"`
	} `json:"tables"`
}

func LoadExportConfigs() []ExportConfig {
	bytes, err := os.ReadFile("application.json")
	if err != nil {
		panic(err)
	}
	var exportConfig []ExportConfig
	err = json.Unmarshal(bytes, &exportConfig)
	if err != nil {
		panic(err)
	}
	return exportConfig
}
