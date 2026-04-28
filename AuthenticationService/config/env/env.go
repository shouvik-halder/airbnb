package env

import (
	"fmt"
	"os"

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

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error Loading env file")
	}
}
