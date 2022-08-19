package main

import (
	"github.com/gin-gonic/gin"
	"github.com/snsvistunov/go-microservices/go-proxy-service/config"
	"github.com/snsvistunov/go-microservices/go-proxy-service/proxy"
)

const configPath = "config/envs"

func main() {
	config.LoadConfig(configPath)
	service := gin.Default()

	service.GET("/microservice/name", proxy.ReverseProxy)
	service.GET("/user/profile", proxy.ReverseProxy)

	service.Run(config.Cfg.Port)
}
