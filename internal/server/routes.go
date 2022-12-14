package server

import "github.com/gin-gonic/gin"

func registerRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}
