package config

import "os"

type Config struct {
	Debug string
}

var (
	cfg *Config
)

func Setconfig() *Config {

	if debugData := os.Getenv("DEBUG"); debugData == "" {
		cfg.Debug = "prod"
	} else {
		cfg.Debug = debugData
	}

	return cfg
}

func GetConfig() *Config {
	return cfg
}
