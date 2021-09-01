package router

import (
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func registerUserRouter(r *mux.Router) {

	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/get/{username}", controller.UserGetController).Methods(http.MethodGet)
	userRouter.HandleFunc("/create", controller.UserCreateController).Methods(http.MethodPost)
	
}