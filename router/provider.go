package router

import "github.com/gorilla/mux"

func Register(r *mux.Router) {
	registerCarRouter(r)
	registerUserRouter(r)
	registerGroupRouter(r)
}
