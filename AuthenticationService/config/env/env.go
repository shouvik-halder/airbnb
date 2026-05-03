package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetString(key, fallback string) string {
	loadEnv()

	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}

func GetInt(key string, defaultVal int) int {
	valStr := GetString(key, "")
	if valStr == "" {
		return defaultVal
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultVal
	}

	return val
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error Loading env file")
	}
}
