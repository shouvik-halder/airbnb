package app

import (
	"AuthenticationService/config"
	"fmt"
	"net/http"
	"time"
)

type Application struct {
	Config *config.Config
}

func NewApplication() *Application {
	cfg:= config.Load()
	return &Application{
		Config: cfg,
	}
}


func (app *Application) Run() error {
	server := &http.Server{
		Addr:         app.Config.Server.PORT,
		Handler:      nil, //TODO: Setup a chi router and put it here
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on Port ", server.Addr)

	return server.ListenAndServe()
}
