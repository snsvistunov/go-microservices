package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snsvistunov/go-microservices/go-auth-service/data"
)

func GetAuthStatus(context *gin.Context) {
	username := context.Request.Header["Username"]
	if user, ok := data.Users[username[0]]; ok {
		if user.AuthStatus {
			context.Status(http.StatusOK)
			return
		}
		context.Status(http.StatusUnauthorized)
		return
	}
	context.Status(http.StatusUnauthorized)
}
