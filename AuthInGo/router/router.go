package router

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
	"AuthInGo/utils"
)

type Router interface{
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router, RoleRouter Router) *chi.Mux{
	chiRouter := chi.NewRouter()

	chiRouter.Get("/ping",controllers.PingHandler)
	UserRouter.Register(chiRouter)
	
	RoleRouter.Register(chiRouter)

	chiRouter.HandleFunc("/api/fakestoreservice/*",utils.ProxyToService("https://fakestoreapi.com","/api/fakestoreservice"))
	
	return chiRouter
}