package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naolson2025/go-rest-api/models"
	"github.com/naolson2025/go-rest-api/utils"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
		return
	}

	// generate JWT token
	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate token",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"token": token,
	})
}