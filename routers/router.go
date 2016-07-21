package routers

import (
	"github.com/gorilla/mux"
	"github.com/whitesmith/powered-plants-web/controllers"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetDeviceRoutes(router)
	return router
}

func SetDeviceRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.ConnectDevice).Methods("GET")
	return router
}