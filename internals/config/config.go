package config

import (
	"encoding/json"
	"io/ioutil"
)

var Cfg *Config

type Config struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
	Guild  struct {
		ID string

		Channels struct {
			WelcomeChannel string `json:"welcome"`
			LoggingChannel string `json:"logs"`
		} `json:"channels"`

		Roles struct {
			DefaultRole string `json:"default"`
		} `json:"roles"`
	} `json:"guild"`
}

func ParseConfigFromJSONFile(configPath string) error {
	if cf, err := ioutil.ReadFile(configPath); err != nil {
		return err
	} else {
		var c Config
		if err := json.Unmarshal(cf, &c); err != nil {
			return err
		}
		Cfg = &c
		return nil
	}
}
