package config

import "ReviewService/config/env"

type serverConfig struct {
	PORT string
}

type Config struct {
	Server serverConfig
}

var AppConfig *Config

func Load() *Config {
	AppConfig = &Config{
		Server: serverConfig{
			PORT: env.GetString("PORT", ":8080"),
		},
	}

	return AppConfig
}
