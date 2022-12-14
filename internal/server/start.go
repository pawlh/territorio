package server

import "github.com/gin-gonic/gin"

func Start() {
	router := gin.New()
	registerRoutes(router)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
