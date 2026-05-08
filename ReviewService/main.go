package main

import (
	"ReviewService/app"
	"log"
)

func main() {
	application := app.NewApplication()

	if err:= application.Run(); err!=nil{
		log.Fatal("Server failed to start", err)
	}
}

