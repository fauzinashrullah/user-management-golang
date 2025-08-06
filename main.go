package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
    ID   uint   `gorm:"primaryKey"`
    Name string
    Age  int
}

func main() {
    dsn := "host=localhost user=postgres password=postgres dbname=learn port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("gagal koneksi ke database")
    }

    db.AutoMigrate(&User{})

    r := gin.Default()

    r.GET("/users", func(c *gin.Context) {
        name := c.Query("name")
        var users []User

        query := db
        if name != ""{
            query = query.Where("name LIKE ?", "%"+name+"%")
        }

        query.Find(&users)
        c.JSON(200,users)
    })

    r.POST("/users", func(c *gin.Context) {
        var user User
        if err := c.BindJSON(&user); err != nil {
            c.JSON(400, gin.H{"error": "invalid data"})
            return
        }

        db.Create(&user)
        c.JSON(200, user)
    })

    r.GET("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        var user User
        if err := db.First(&user, "id = ?", id).Error; err != nil{
            c.JSON(400, gin.H{"error": "not found"})
            return
        }
        
        c.JSON(200,user)
    })


    r.Run()
}
