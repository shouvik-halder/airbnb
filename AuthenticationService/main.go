package main

import "AuthenticationService/app"

func main(){
	cfg:= app.Config{
		Addr: ":####",
	}
	application := app.Application{
		Config: cfg,
	}

	application.Run()
}