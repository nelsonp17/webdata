package app

import (
	"github.com/Nelson2017-8/webdata/app/controllers"
	"github.com/gorilla/mux"
)

func Routes(route *mux.Router) {
	route.HandleFunc("/dollar", controllers.PromedioView).Methods("GET")
}
