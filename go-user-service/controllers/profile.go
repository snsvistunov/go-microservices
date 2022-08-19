package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snsvistunov/go-microservices/go-user-service/data"
)

func GetUserProfile(context *gin.Context) {
	//var user models.User

	username := context.Request.Header["Username"]
	if user, ok := data.Users[username[0]]; ok {
		context.SecureJSON(http.StatusOK, gin.H{"UserProfile": user})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Error": "user not found"})
}
