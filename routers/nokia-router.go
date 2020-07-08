package routers

import (
	//"github.com/gorilla/mux"
	"net/http"
	//_ "net/http/pprof"
	"github.com/gkganesh126/nokia-interview/controllers"
)

func SetNokiaRouters(router *http.ServeMux) *http.ServeMux {

	router.HandleFunc("/usersGet", controllers.GetUsers)
	router.HandleFunc("/usersCreate", controllers.CreateUser)
	router.HandleFunc("/usersDelete", controllers.DeleteUser)
	router.HandleFunc("/usersUpdate", controllers.UsersUpdate)

	return router
}
