package game

import (
	"testing"
)

func TestNewStageDimensions(t *testing.T) {
	height := 10
	width := 10

	stage := NewStage(height, width)

	if stage.Height != height {
		t.Errorf("Stage height is not %d", height)
	}

	if stage.Width != width {
		t.Errorf("Stage width is not %d", width)
	}

	if len(stage.Board) != height {
		t.Errorf("Stage board height is not %d", height)
	}

	for i := range stage.Board {
		if len(stage.Board[i]) != width {
			t.Errorf("Stage board width is not %d on row %d", width, i)
		}
	}

	for i := range stage.Board {
		for j := range stage.Board[i] {
			if stage.Board[i][j].Owner.Name != "dummy" {
				t.Errorf("Stage board pixel owner is not dummy on row %d, column %d", i, j)
			}
		}
	}
}

func TestNewStagePlayers(t *testing.T) {
	height := 10
	width := 10

	stage := NewStage(height, width)

	if len(stage.Players) != 1 {
		t.Errorf("Stage players is not 1")
	}

	if _, ok := stage.Players["dummy"]; !ok {
		t.Errorf("Stage players does not contain dummy")
	}

	if stage.Players["dummy"].Name != "dummy" {
		t.Errorf("Stage players dummy name is not dummy")
	}

	if stage.Players["dummy"].Color != Black {
		t.Errorf("Stage players dummy color is not Black")
	}
}

func TestSetOwner(t *testing.T) {
	height := 10
	width := 10

	stage := NewStage(height, width)

	stage.SetOwner(0, 0, NewPlayer("test", Red))

	if stage.GetOwner(0, 0).Name != "test" {
		t.Errorf("Stage owner is not test")
	}
}
