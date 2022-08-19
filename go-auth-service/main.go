package main

import (
	"github.com/gin-gonic/gin"
	"github.com/snsvistunov/go-microservices/go-auth-service/config"
	"github.com/snsvistunov/go-microservices/go-auth-service/controllers"
)

const configPath = "config/envs"

func main() {
	config.LoadConfig(configPath)

	service := gin.Default()

	service.GET("/auth", controllers.GetAuthStatus)

	service.Run(config.Cfg.Port)
}
