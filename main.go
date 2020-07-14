package main

import (
	"flag"
	"net/http"

	"github.com/gkganesh126/nokia-interview/common"
	"github.com/gkganesh126/nokia-interview/controllers"
	"github.com/gkganesh126/nokia-interview/controllers/cache"
	"github.com/gkganesh126/nokia-interview/routers"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	glog.Info("Listening...")
	server.ListenAndServe()
}
