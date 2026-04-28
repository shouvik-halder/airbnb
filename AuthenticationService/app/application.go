package app

import (
	"AuthenticationService/config"
	"AuthenticationService/controllers"
	db "AuthenticationService/db/repositories"
	"AuthenticationService/router"
	// v1router "AuthenticationService/router/v1"
	"AuthenticationService/services"
	"fmt"
	"net/http"
	"time"
)

type Application struct {
	Config *config.Config
	Store  *db.Storage
}

func NewApplication() *Application {
	return &Application{
		Config: config.Load(),
		Store:  db.InitStorage(),
	}
}

func (app *Application) Run() error {

	ur:= db.NewUserRepository();
	us:=services.NewUserService(ur);
	uc:= controllers.NewUserController(us);
	uRouter:=router.NewUserRouter(uc);
	server := &http.Server{
		Addr:         app.Config.Server.PORT,
		Handler:      router.InitializeRouter(uRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on Port ", server.Addr)

	return server.ListenAndServe()
}
