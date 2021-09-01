package model

type ResponseError struct {
	Status  int `json:""`
	Message string
}