package controller

import (
	"encoding/json"
	"myapp/model"
	"myapp/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GroupsGetAll(w http.ResponseWriter, r *http.Request) {

	res, err := service.GroupsGetAll(r.Context(), EntClient)
	if err != nil {
		json.NewEncoder(w).Encode(model.ResponseError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(res)
}

func GroupGetByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	groupID, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(model.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	res, err := service.GroupsGetByID(r.Context(), EntClient, groupID)
	if err != nil {
		json.NewEncoder(w).Encode(model.ResponseError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(res)
}
