package app

import (
	"fmt"
	"net/http"
	"time"
)


type Application struct {
	Config Config
}

type Config struct {
	Addr string // PORT
}

func (app *Application) Run() error{
	server := &http.Server{
		Addr: app.Config.Addr,
		Handler: nil, //TODO: Setup a chi router and put it here
		// ReadTimeout: ##* time.Second,
		// WriteTimeout: ##* time.Second,
	}

	fmt.Println("Starting server on Port ", server.Addr)

	return  server.ListenAndServe()
}