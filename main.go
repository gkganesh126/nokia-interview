package main

import (
	"flag"
	"net/http"

	"github.com/gkganesh126/nokia-interview/common"
	"github.com/gkganesh126/nokia-interview/controllers"
	"github.com/gkganesh126/nokia-interview/controllers/cache"
	"github.com/gkganesh126/nokia-interview/routers"
	"github.com/golang/glog"
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

	controllers.StorageCache = *cache.NewStorage()
	controllers.ReloadCacheFromDb()

	glog.Info("Listening...")
	server.ListenAndServe()
}
