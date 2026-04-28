package config

import (
	"AuthenticationService/config/env"
)

type serverConfig struct {
	PORT string
}

type Config struct {
	Server serverConfig
}

func Load() *Config {

	return &Config{
		Server: serverConfig{
			PORT: env.GetString("PORT", ":8080"),
		},
	}
}
