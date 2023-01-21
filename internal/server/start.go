package server

import (
	"github.com/gin-gonic/gin"
	"territorio/internal/game"
)

var games = make(map[string]game.Stage)

func Start() {

	games["test"] = game.NewStage(10, 10)

	router := gin.New()
	registerRoutes(router)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
