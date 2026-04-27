package main

import "AuthenticationService/app"

func main(){
	cfg:= app.NewConfig(":####");
	application := app.NewApplication(cfg)

	application.Run()
}