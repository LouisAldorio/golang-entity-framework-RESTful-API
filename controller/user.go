package controller

import (
	"encoding/json"
	"myapp/model"
	"myapp/service"
	"net/http"

	"github.com/gorilla/mux"
)

func UserCarGetController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	res, err := service.QueryUserCars(r.Context(), EntClient, vars["username"])
	if err != nil {
		json.NewEncoder(w).Encode(model.ResponseError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(res)
}

func UserGetController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	res, err := service.QueryUser(r.Context(), EntClient, vars["username"])
	if err != nil {
		json.NewEncoder(w).Encode(model.ResponseError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(res)
}

func UserCreateController(w http.ResponseWriter, r *http.Request) {

	var param model.UserCreateInput

	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		json.NewEncoder(w).Encode(model.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	res, err := service.CreateUser(r.Context(), EntClient, param)
	if err != nil {
		json.NewEncoder(w).Encode(model.ResponseError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(res)
}
