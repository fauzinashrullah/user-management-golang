package main

import (
	"user-management-golang/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.Route(r)

	r.Run()
}
