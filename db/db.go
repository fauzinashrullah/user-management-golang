package db

import (
	"user-management-golang/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {

	dsn := "host=localhost user=postgres password=postgres dbname=learn port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("gagal koneksi ke database")
	}

	db.AutoMigrate(&model.User{})
	return db
}
