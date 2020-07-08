package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/gkganesh126/nokia-interview/common"
	"github.com/gkganesh126/nokia-interview/routers"
)

// Entry point for the program
func main() {

	flag.Parse()
	flag.Set("logtostderr", "true")
	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}

	glog.Info("Listening...")
	server.ListenAndServe()
}
