package handlers

import (
	"net/http"
	"eamplle.com/test/02-crud/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context)  {
	events := models.GetAllUsers()
	context.JSON(http.StatusOK, events)
}

func CreateUser(context *gin.Context)  {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = 1
	user.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}