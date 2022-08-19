package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snsvistunov/go-microservices/go-user-service/config"
)

func GetMicroserviceName(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"MicroserviceName": config.Cfg.ServiceName})
}
