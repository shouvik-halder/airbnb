package app

import (
	"ReviewService/config"
	dbConfig "ReviewService/config/db"
	"ReviewService/controllers"
	"ReviewService/db"
	"ReviewService/routers"
	v1 "ReviewService/routers/v1"
	"ReviewService/services"
	"fmt"
	"net/http"
	"time"
)

type Application struct {
	Cfg   *config.Config
	Store *db.Storage
}

func NewApplication() *Application {
	cfg := config.Load()
	if err := dbConfig.SetupDB(cfg); err != nil {
		fmt.Println("Error Setting up db", err)
	}
	
	return &Application{
		Cfg: cfg,
		Store: db.InitStorage(),
	}
}

func (app *Application) Run() error {
	reviewRepo := app.Store.ReviewRepository
	reviewService := services.NewReviewService(reviewRepo);
	reviewController := controllers.NewReviewController(reviewService)
	reviewRouter := v1.NewReviewRouter(reviewController)
	server := &http.Server{
		Addr:         app.Cfg.Server.PORT,
		Handler:      routers.InitializeRouters(reviewRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server started on port 3004")
	return server.ListenAndServe()
}
