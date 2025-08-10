package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name" validate:"gte=3"`
	Age      int    `json:"age" validate:"gte=18,lte=60"`
	Username string `json:"username" validate:"gte=3"`
	Password string `json:"password" validate:"gte=8"`
}
