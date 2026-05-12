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
	if err := dbconfig.SeedDB(); err != nil {
		logger.Log.Error().Msg(err.Error())
	}
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

	rr := dbrepo.NewRolesRepository(dbconfig.GetDB())
	rs := services.NewRolesService(rr)
	rc := controllers.NewRoleController(rs)
	rRouter := v1router.NewRolesRouter(rc)

	pr := dbrepo.NewPermissionRepository(dbconfig.GetDB())
	ps := services.NewPermissionsService(pr)
	pc := controllers.NewPermissionController(ps)
	pRouter := v1router.NewPermissionsRouter(pc)

	urr := dbrepo.NewUserRolesRepository(dbconfig.GetDB())
	urs := services.NewUserRolesService(urr)
	urc := controllers.NewUserRoleController(urs)
	urRouter := v1router.NewUserRolesRouter(urc)

	server := &http.Server{
		Addr:         app.Config.Server.PORT,
		Handler:      router.InitializeRouter(uRouter, rRouter, pRouter, urRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on Port ", server.Addr)
	logger.Log.Info().Msg(fmt.Sprintf("Starting server on Port %s", server.Addr))

	return server.ListenAndServe()
}
