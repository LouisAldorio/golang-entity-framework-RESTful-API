package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"myapp/controller"
)

func registerCarRouter(r *mux.Router) {

	carRouter := r.PathPrefix("/cars").Subrouter()
	carRouter.HandleFunc("/get/{username}", controller.UserCarGetController).Methods(http.MethodGet)
	
}

