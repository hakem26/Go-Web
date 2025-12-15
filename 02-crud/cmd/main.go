package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/api/v1/users", GetUsers)

	server.Run(":8080")
}

func GetUsers(context *gin.Context)  {
	context.JSON(http.StatusOK, gin.H{"message": "Hello"})
}