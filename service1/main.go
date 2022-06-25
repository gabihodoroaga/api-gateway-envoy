package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// This is required in order to serve swagger file
	_ "github.com/gabihodoroaga/api-gateway-envoy/service1/docs"
)

// @title           Service1 API
// @version         1.0
// @description     This is a sample server

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/service1/ping", ping)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

// @Summary      Get organization keys statistics
// @schemes      https
// @Description  Gets packet statistics for all public keys
// @Success      200      string  string
// @Router       /service1/ping [get]
func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

