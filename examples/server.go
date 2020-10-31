package main

import (
	openapi "github.com/cdimascio/gin-openapi/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(openapi.ValidateRequests("spec.yaml"))

	router.GET("/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"id": "pong"})
	})

	router.POST("/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"id": "pong"})
	})

	router.Run(":8080")
}
