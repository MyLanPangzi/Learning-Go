package config

import (
	"encoding/json"
	"os"
)

func Load(fileName string) (*DefaultConfig, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	m := map[string]any{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return &DefaultConfig{m}, nil
}
