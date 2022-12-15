package main

import (
	"fmt"
	"territorio/internal/game"
	"territorio/internal/render"
)

func main() {
	// start server
	//server.Start()
	stage := game.NewStage(10, 10)
	stage.AddPlayer(game.NewPlayer("testPlayer", game.Blue))
	stage.SetOwner(0, 0, stage.Players["testPlayer"])
	stage.SetOwner(0, 1, stage.Players["testPlayer"])
	stage.SetOwner(1, 0, stage.Players["testPlayer"])

	stage.Players["testPlayer"].Credits = 1000

	render.Render(stage)

	stage.Tick()
	fmt.Println()
	stage.SpreadPlayer(stage.Players["testPlayer"], stage.Players["dummy"])
	render.Render(stage)

	fmt.Println()
	stage.SpreadPlayer(stage.Players["testPlayer"], stage.Players["dummy"])
	render.Render(stage)

	fmt.Println()
	stage.SpreadPlayer(stage.Players["testPlayer"], stage.Players["dummy"])
	render.Render(stage)

}
