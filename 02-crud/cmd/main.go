package main

import (

	"eamplle.com/test/02-crud/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/api/v1/users", handlers.GetUsers)
	server.POST("/api/v1/users", handlers.CreateUser)

	server.Run(":8080")
}