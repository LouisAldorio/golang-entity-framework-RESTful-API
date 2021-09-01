package model

type UserCreateInput struct {
	Username string   `json:"username"`
	Age      int      `json:"age"`
	Cars     []string `json:"cars"`
}
