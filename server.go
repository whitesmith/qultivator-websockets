package main

import (
	"log"
	"net/http"
	"github.com/codegangsta/negroni"
	"github.com/whitesmith/powered-plants-web/config"
	"github.com/whitesmith/powered-plants-web/routers"
)

func main() {
	config.Init()
	router := routers.InitRoutes()
	server := negroni.Classic()
	server.UseHandler(router)

	log.Printf("Listening on port: " + config.Get().HostPort)
	http.ListenAndServe(":" + config.Get().HostPort, server)
}