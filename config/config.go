package config

import (
	"github.com/BurntSushi/toml"
)

// Config this app config
type TwitterConfig struct {
	ApiKey            string
	ApiKeySecret      string
	AccessToken       string
	AccessTokenSecret string
	TimeoutSecond     int
}

func ReadConfig(path string) (*TwitterConfig, bool) {
	var config TwitterConfig
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, false
	}
	return &config, true
}
