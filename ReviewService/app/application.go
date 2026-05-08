package app

import (
	"ReviewService/config"
	"ReviewService/routers"
	v1 "ReviewService/routers/v1"
	"fmt"
	"net/http"
	"time"
)

type Application struct {
	cfg *config.Config
}

func NewApplication() *Application {
	cfg := config.Load()
	return &Application{
		cfg: cfg,
	}
}

func (app *Application) Run() error {

	reviewRouter := v1.NewReviewRouter()
	server := &http.Server{
		Addr:         app.cfg.Server.PORT,
		Handler:      routers.InitializeRouters(reviewRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server started on port 3004")
	return server.ListenAndServe()
}
