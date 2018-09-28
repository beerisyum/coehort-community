package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Twitch struct {
		client_id string `json:"twitch"`
	} `json:"twitch"`
	DB struct {
		DSN string `json:"db"`
	} `json:"db"`
	Bot struct {
		Prefix string `json:"bot"`
	} `json:"bot"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config *Config

	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
