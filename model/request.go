package model

type LoginRequest struct {
	Username string
	Password string
}
type RegisterRequest struct {
	Name     string `json:"name" validate:"gte=3"`
	Age      int    `json:"age" validate:"gte=18,lte=60"`
	Username string `json:"username" validate:"gte=3"`
	Password string `json:"password" validate:"gte=8"`
}
