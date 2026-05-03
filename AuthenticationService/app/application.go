package app

import (
	"AuthenticationService/config"
	dbconfig "AuthenticationService/config/db"
	"AuthenticationService/config/logger"
	"AuthenticationService/controllers"
	dbrepo "AuthenticationService/db/repositories"
	"AuthenticationService/router"
	v1router "AuthenticationService/router/v1"
	"AuthenticationService/services"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Application struct {
	Config *config.Config
	Store  *dbrepo.Storage
}

func NewApplication() *Application {
	cfg := config.Load()
	if err := dbconfig.SetupDB(cfg); err != nil {
		log.Fatal(err)
	}
	logger.InitLogger(cfg)
	return &Application{
		Config: cfg,
		Store:  dbrepo.InitStorage(),
	}
}

func (app *Application) Run() error {
	ur := dbrepo.NewUserRepository(dbconfig.GetDB())
	us := services.NewUserService(ur, app.Config.Auth.TokenSecret)
	uc := controllers.NewUserController(us)
	uRouter := v1router.NewUserRouter(uc)
	server := &http.Server{
		Addr:         app.Config.Server.PORT,
		Handler:      router.InitializeRouter(uRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on Port ", server.Addr)
	logger.Log.Info().Msg(fmt.Sprintf("Starting server on Port %s", server.Addr))

	return server.ListenAndServe()
}
