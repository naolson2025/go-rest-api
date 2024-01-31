package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naolson2025/go-rest-api/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		// need to abort because we don't want to call the next function
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized, no token provided",
		})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		fmt.Println(err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized, invalid token",
		})
		return
	}

	// add the userId to the gin context
	context.Set("userId", userId)
	// call the next middleware or function in the chain
	context.Next()
}
