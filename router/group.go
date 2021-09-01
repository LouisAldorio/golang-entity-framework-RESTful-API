package router

import (
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func registerGroupRouter(r *mux.Router) {

	groupRouter := r.PathPrefix("groups").Subrouter()
	groupRouter.HandleFunc("/", controller.GroupsGetAll).Methods(http.MethodGet)
	groupRouter.HandleFunc("/{id}", controller.GroupGetByID).Methods(http.MethodPost)
}
