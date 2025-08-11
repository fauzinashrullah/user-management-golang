package model

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Username string `json:"username"`
}

type ApiResponse struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
	Data    any    `json:"data"`
}

type LoginResponse struct {
	Name  string `json:"name"`
	Role  string `json:"role"`
	Token string `json:"token"`
}
