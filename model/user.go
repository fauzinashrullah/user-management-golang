package model

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Age      int
	Username string
	Password string
	Role     string
}
