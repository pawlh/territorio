package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"territorio/internal/game"
	"territorio/internal/render"
)

type AddPlayerRequest struct {
	Name  string     `json:"name"`
	Color game.Color `json:"color"`
	Row   int        `json:"row"`
	Col   int        `json:"col"`
}

type AttackRequest struct {
	Attacker string `json:"attacker"`
	Defender string `json:"defender"`
}

func registerRoutes(router *gin.Engine) {
	router.GET("/addplayer", func(c *gin.Context) {
		session, ok := getSession(c)
		if !ok {
			c.JSON(400, gin.H{"error": "no session provided or session invalid"})
			return
		}

		var request AddPlayerRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}

		if session.GetOwner(request.Row, request.Col) != session.Players["dummy"] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "tile already occupied"})
			return
		}

		session.AddPlayer(game.NewPlayer(request.Name, request.Color))
		session.SetOwner(request.Row, request.Col, session.Players[request.Name])

		c.String(200, fmt.Sprintf("adding %s", request.Name))
	})

	router.GET("/tick", func(c *gin.Context) {
		session, ok := getSession(c)
		if !ok {
			c.JSON(400, gin.H{"error": "no session provided or session invalid"})
			return
		}

		session.Tick()
	})

	router.GET("/attack", func(c *gin.Context) {
		session, ok := getSession(c)
		if !ok {
			c.JSON(400, gin.H{"error": "no session provided or session invalid"})
			return
		}

		var request AttackRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}

		if err := session.SpreadPlayer(session.Players[request.Attacker], session.Players[request.Defender]); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		fmt.Println()
		render.Render(session)
	})
}

func getSession(c *gin.Context) (game.Stage, bool) {
	session, ok := games[c.Request.Header.Get("session")]
	return session, ok
}
