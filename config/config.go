package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Environment   string `json:"environment"`
	ServerAddress string `json:"server_address"`
	ServerPort    int    `json:"server_port"`
	PostgresURL   string `json:"postgres_url"`
}

func NewConfig() *Config {
	var config *Config

	content, err := os.ReadFile("./config/config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, &config)

	if err != nil {
		panic(err)
	}

	return config
}
