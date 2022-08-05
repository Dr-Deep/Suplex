package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/tripledoomer/Suplex/internals/logging"
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
			DefaultRole   string `json:"default"`
			ModeratorRole string `json:"moderator"`
		} `json:"roles"`
	} `json:"guild"`
	Logging struct {
		Level int `json:"level"`
	} `json:"logging"`
}

func ParseConfigFromJSONFile(configPath string) {
	if cf, err := ioutil.ReadFile(configPath); err != nil {
		logging.Logf(logging.Fatal, err.Error())
	} else {
		var c Config
		if err := json.Unmarshal(cf, &c); err != nil {
			logging.Logf(logging.Fatal, err.Error())
		}
		Cfg = &c
	}
}
