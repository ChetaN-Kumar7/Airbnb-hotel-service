package app

import (
	dbConfig "ReviewService/config/db"
	config "ReviewService/config/env"
	"ReviewService/controllers"
	repo "ReviewService/db/repositories"
	router "ReviewService/routes"
	"ReviewService/services"
	"fmt"
	"net/http"
	"time"
)


type Config struct{
	Addr string
}

type Application struct{
	Config Config
}

func NewConfig() Config{
	port := config.GetString("PORT",":8080")
	return Config{
		Addr : port,
	}
}

func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,

	}
}


func (app *Application) Run() error{

	db,err:=dbConfig.SetupDb()

	if err != nil {
		fmt.Println("Error setup database:",err)
		return err
	}

	rr := repo.NewReviewRepository(db)
	rs := services.NewReviewService(rr)
	rc := controllers.NewReviewController(rs)
	rRouter := router.NewReviewRouter(rc)

	server := &http.Server{
		Addr: app.Config.Addr,
		Handler : router.SetupRouter(rRouter),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server on",app.Config.Addr)
	return server.ListenAndServe()
}