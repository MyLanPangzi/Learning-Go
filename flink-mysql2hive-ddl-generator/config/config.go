package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Endpoint string `json:"endpoint"`
	Configs  []ExportConfig
}

type ExportConfig struct {
	Arn      string `json:"arn"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Db       string `json:"db"`
	DestDb   string `json:"dest_db"`
	Tables   []struct {
		Regexp string `json:"regexp"`
		Table  string `json:"table"`
	} `json:"tables"`
}

func LoadConfiguration() Configuration {
	bytes, err := os.ReadFile("application.json")
	if err != nil {
		panic(err)
	}
	var cfg Configuration
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
