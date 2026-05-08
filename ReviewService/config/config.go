package config

import "ReviewService/config/env"

type serverConfig struct {
	PORT string
}

type dbConfig struct {
	DBUSER string
	DBPASS string
	DBNET  string
	DBNAME string
	DBADDR string
}

type Config struct {
	Server serverConfig
	DB     dbConfig
}

var AppConfig *Config

func Load() *Config {
	AppConfig = &Config{
		Server: serverConfig{
			PORT: env.GetString("PORT", ":8080"),
		},
		DB: dbConfig{
			DBUSER: env.GetString("DBUSER", "user"),
			DBPASS: env.GetString("DBPASS", ""),
			DBNET:  env.GetString("DBNET", "tcp"),
			DBNAME: env.GetString("DBNAME", ""),
			DBADDR: env.GetString("DBADDR", ""),
		},
	}

	return AppConfig
}
