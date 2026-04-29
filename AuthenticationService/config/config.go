package config

import (
	"AuthenticationService/config/env"
)

type serverConfig struct {
	PORT string
}

type dbConfig struct {
	DBUSER string
	DBPASS string
	DBADDR string
	DBNAME string
	DBNET  string
}

type Config struct {
	Server serverConfig
	DB     dbConfig
}

func Load() *Config {
	return &Config{
		Server: serverConfig{
			PORT: env.GetString("PORT", ":8080"),
		},
		DB: dbConfig{
			DBUSER: env.GetString("DBUSER", "user"),
			DBPASS: env.GetString("DBPASS", ""),
			DBADDR: env.GetString("DBADDR", ""),
			DBNAME: env.GetString("DBNAME", ""),
			DBNET:  env.GetString("DBNET", "tcp"),
		},
	}
}
