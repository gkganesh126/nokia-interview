package routers

import (
	"net/http"
)

func InitRoutes() *http.ServeMux {
	//router := mux.NewRouter().StrictSlash(false)
	h := http.NewServeMux()
	// Routes for the User entity
	router := SetNokiaRouters(h)
	return router
}
