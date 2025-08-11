package db

import (
	"user-management-golang/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {

	dsn := "host=localhost user=postgres password=postgres dbname=user_management_golang port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed connect to database")
	}

	db.AutoMigrate(&model.User{})
	return db
}
