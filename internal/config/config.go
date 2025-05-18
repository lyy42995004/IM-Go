package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type LogConfig struct {
	LogPath string `toml:"logPath"`
}

type Config struct {
	LogConfig `toml:"logConfig"`
}

var config *Config

func LoadConfig() error {
	if _, err := toml.DecodeFile("/home/Tian/im-go/configs/config_local.toml", config); err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func GetConfig() *Config {
	if config == nil {
		config = new(Config)
		_ = LoadConfig()
	}
	return config
}
