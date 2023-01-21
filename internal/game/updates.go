package game

import (
	"fmt"
)

type CreditError struct {
	Short int
}

func (e CreditError) Error() string {
	return fmt.Sprintf("not enough credits: %d short", e.Short)
}

func (s *Stage) SpreadPlayer(attacker *Player, defender *Player) error {

	spreadableCells := s.getSpreadPixels(attacker, defender)
	spreadCost := s.getSpreadCost(spreadableCells, defender)
	if remaining := attacker.Credits - (spreadCost * 2); remaining < 0 {
		return CreditError{Short: remaining * -1}
	}

	for _, pixel := range spreadableCells {
		s.SetOwner(pixel.Row, pixel.Col, attacker)
	}

	defender.Credits -= spreadCost
	//if defender.Credits > 0 {
	attacker.Credits -= spreadCost * 2
	//} else {
	//	attacker.Credits -= (spreadCost - (defender.Credits * -1)) * 2
	//	defender.Credits = 0
	//}

	return nil
}

func (s *Stage) getSpreadPixels(attacker *Player, defender *Player) []*Pixel {
	var spreadableCells []*Pixel

	for i := range s.Board {
		for j := range s.Board[i] {
			if s.GetOwner(i, j) == defender && s.adjacentPixelIsPlayer(i, j, attacker) {
				spreadableCells = append(spreadableCells, s.Board[i][j])
			}
		}
	}

	return spreadableCells
}

func (s *Stage) getSpreadCost(spreadableCells []*Pixel, defender *Player) int {
	return int(float64(len(spreadableCells)) / float64(defender.TerritorySize) * float64(defender.Credits))
}

func (s *Stage) adjacentPixelIsPlayer(row int, col int, player *Player) bool {
	// check all eight surrounding pixels, return true if any are the player
	if row > 0 && s.Board[row-1][col].Owner == player {
		return true
	} else if row < s.Height-1 && s.Board[row+1][col].Owner == player {
		return true
	} else if col > 0 && s.Board[row][col-1].Owner == player {
		return true
	} else if col < s.Width-1 && s.Board[row][col+1].Owner == player {
		return true
	} else if row > 0 && col > 0 && s.Board[row-1][col-1].Owner == player {
		return true
	} else if row > 0 && col < s.Width-1 && s.Board[row-1][col+1].Owner == player {
		return true
	} else if row < s.Height-1 && col > 0 && s.Board[row+1][col-1].Owner == player {
		return true
	} else if row < s.Height-1 && col < s.Width-1 && s.Board[row+1][col+1].Owner == player {
		return true
	}
	return false
}

func (s *Stage) Tick() {
	for _, player := range s.Players {
		if player != s.Players["dummy"] {
			player.Credits += player.TerritorySize
		}
	}
}
