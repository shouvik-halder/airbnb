package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetString(key, dafaultValue string) string {
	load()
	if res, ok := os.LookupEnv(key); ok {
		return res
	}
	return dafaultValue
}

func load() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("unable to load env file")
	}
}
