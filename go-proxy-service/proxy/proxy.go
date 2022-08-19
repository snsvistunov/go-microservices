package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/snsvistunov/go-microservices/go-proxy-service/config"
)

func ReverseProxy(c *gin.Context) {
	userService, err := url.Parse("http://" + config.Cfg.UserServiceName)
	if err != nil {
		log.Fatal(err)
	}
	requestURL, err := url.Parse("/user/profile")
	if err != nil {
		log.Fatal(err)
	}
	if *c.Request.URL == *requestURL {
		fmt.Printf("Username from header %v", c.Request.Header["Username"][0])
		if !checkAuth(c.Request.Header["Username"][0]) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "request is unauthorized"})
			return
		}
	}
	proxy := httputil.NewSingleHostReverseProxy(userService)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = userService.Host
		req.URL.Scheme = userService.Scheme
		req.URL.Host = userService.Host
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func checkAuth(user string) bool {
	client := resty.New()
	request := client.NewRequest()
	request.Header.Add("Username", user)
	response, err := request.Get("http://" + config.Cfg.AuthServiceName + "/auth")
	if err != nil {
		log.Println("Proxy: unable to connect AuthService")
		return false
	}
	if response.StatusCode() == http.StatusOK {
		return true
	}
	return false
}
