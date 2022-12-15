package render

import (
	"fmt"
	"territorio/internal/game"
)

func Render(stage game.Stage) {
	// print board out to terminal
	for i := range stage.Board {
		for j := range stage.Board[i] {
			// print out pixel
			// set terminal color to black
			fmt.Print(colorToTerm(stage.Board[i][j].Owner.Color))
			fmt.Print(" □ ")
			fmt.Print(colorToTerm(game.Black))
		}
		fmt.Println()
	}
}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func colorToTerm(color game.Color) string {
	switch color {
	case game.Black:
		return "\033[30m"
	case game.Red:
		return "\033[31m"
	case game.Green:
		return "\033[32m"
	case game.Blue:
		return "\033[34m"
	case game.White:
		return "\033[97m"
	default:
		return "\\033[0m"

	}
}
