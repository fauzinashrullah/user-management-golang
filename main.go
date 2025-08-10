package main

import (
	"user-management-golang/router"
)

func main() {
	r := router.Route()
	r.Run(":8080")
}
