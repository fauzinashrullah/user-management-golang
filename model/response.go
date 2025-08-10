package model

type UserResponse struct {
	ID       uint
	Name     string
	Age      int
	Username string
}

type ApiResponse struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
	Data    any    `json:"data"`
}
