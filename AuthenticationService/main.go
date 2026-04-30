package main

import (
	"AuthenticationService/app"
	"log"
)

func main() {
	application := app.NewApplication()

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
