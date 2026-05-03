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

type authConfig struct {
	TokenSecret string
}

type loggerConfig struct {
	MAXSIZEMB      int
	MAXBACKUPCOUNT int
	MAXAGEDAYS     int
}

type Config struct {
	Server serverConfig
	DB     dbConfig
	Auth   authConfig
	Logger loggerConfig
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
			DBADDR: env.GetString("DBADDR", ""),
			DBNAME: env.GetString("DBNAME", ""),
			DBNET:  env.GetString("DBNET", "tcp"),
		},
		Auth: authConfig{
			TokenSecret: env.GetString("AUTH_TOKEN_SECRET", "dev-auth-token-secret-change-me"),
		},
		Logger: loggerConfig{
			MAXSIZEMB:      env.GetInt("MAXSIZEMB", 20),
			MAXAGEDAYS:     env.GetInt("MAXAGEDAYS", 14),
			MAXBACKUPCOUNT: env.GetInt("MAXBACKUPCOUNT", 2),
		},
	}
	return AppConfig
}

func GetConfig() *Config {
	return AppConfig
}
