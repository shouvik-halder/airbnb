package main

import "AuthenticationService/app"

func main() {
	application := app.NewApplication()

	application.Run()
}
