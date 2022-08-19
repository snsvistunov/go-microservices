package main

import (
	"github.com/gin-gonic/gin"
	"github.com/snsvistunov/go-microservices/go-user-service/config"
	"github.com/snsvistunov/go-microservices/go-user-service/controllers"
)

const configPath = "config/envs"

func main() {
	config.LoadConfig(configPath)

	service := gin.Default()

	service.GET("/user/profile", controllers.GetUserProfile)
	service.GET("/microservice/name", controllers.GetMicroserviceName)

	service.Run(config.Cfg.Port)
}
