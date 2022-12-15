package game

import (
	"testing"
)

func TestGetSpreadPixels(t *testing.T) {
	height := 10
	width := 10

	stage := NewStage(height, width)
	stage.AddPlayer(NewPlayer("testPlayer", Blue))

	// Set 0,0 pixel to a player
	stage.SetOwner(0, 0, stage.Players["testPlayer"])

	// Get the spread pixels
	if spreadCount := len(stage.getSpreadPixels(stage.Players["testPlayer"], stage.Players["dummy"])); spreadCount != 3 {
		t.Errorf("Spread pixel size should have been %d but was %d", 3, spreadCount)
	}

	// Set 0,1 pixel to a player
	stage.SetOwner(0, 1, stage.Players["testPlayer"])
	if spreadCount := len(stage.getSpreadPixels(stage.Players["testPlayer"], stage.Players["dummy"])); spreadCount != 4 {
		t.Errorf("Spread pixel size should have been %d but was %d", 4, spreadCount)
	}

	// set 9,9 pixel to a player
	stage.SetOwner(9, 9, stage.Players["testPlayer"])
	if spreadCount := len(stage.getSpreadPixels(stage.Players["testPlayer"], stage.Players["dummy"])); spreadCount != 7 {
		t.Errorf("Spread pixel size should have been %d but was %d", 7, spreadCount)
	}

	// set 5,5 pixel to a player
	stage.SetOwner(5, 5, stage.Players["testPlayer"])
	if spreadCount := len(stage.getSpreadPixels(stage.Players["testPlayer"], stage.Players["dummy"])); spreadCount != 15 {
		t.Errorf("Spread pixel size should have been %d but was %d", 15, spreadCount)
	}
}

func TestSpreadCost(t *testing.T) {
	height := 10
	width := 10

	stage := NewStage(height, width)
	stage.AddPlayer(NewPlayer("testPlayer", Blue))
	stage.Players["testPlayer"].Credits = 1000
	stage.Players["dummy"].Credits = 1000

	{
		spreadableCells := stage.getSpreadPixels(stage.Players["testPlayer"], stage.Players["dummy"])

		if cost := stage.getSpreadCost(spreadableCells, stage.Players["dummy"]); cost != 0 {
			t.Errorf("Spread cost should have been %d but was %d", 0, cost)
		}

	}
	{
		// set top left square to be owned by testPlayer
		stage.SetOwner(0, 0, stage.Players["testPlayer"])
		stage.SetOwner(0, 1, stage.Players["testPlayer"])
		stage.SetOwner(1, 0, stage.Players["testPlayer"])
		stage.SetOwner(1, 1, stage.Players["testPlayer"])

		spreadableCells := stage.getSpreadPixels(stage.Players["dummy"], stage.Players["testPlayer"])

		if cost := stage.getSpreadCost(spreadableCells, stage.Players["testPlayer"]); cost != 750 {
			t.Errorf("Spread cost should have been %d but was %d", 750, cost)
		}
	}
}

func TestSpreadPlayer(t *testing.T) {
	height := 10
	width := 10
	{
		stage := NewStage(height, width)
		stage.Players["dummy"].Credits = 1000

		stage.AddPlayer(NewPlayer("testPlayer", Blue))
		stage.Players["testPlayer"].Credits = 100
		// give testPlayer some pixels
		stage.SetOwner(0, 0, stage.Players["testPlayer"])
		stage.SetOwner(0, 1, stage.Players["testPlayer"])
		stage.SetOwner(0, 2, stage.Players["testPlayer"])
		stage.SetOwner(1, 0, stage.Players["testPlayer"])

		if err := stage.SpreadPlayer(stage.Players["testPlayer"], stage.Players["dummy"]); err == nil {
			t.Errorf("Spread should have failed, too expensive for test player")
		}

		if stage.Players["testPlayer"].Credits != 100 {
			t.Errorf("Test player should have still had %d credits but had %d", 100, stage.Players["testPlayer"].Credits)
		}

		if stage.Players["dummy"].Credits != 1000 {
			t.Errorf("Dummy player should have still had %d credits but had %d", 1000, stage.Players["dummy"].Credits)
		}
	}
	{
		stage := NewStage(height, width)
		stage.Players["dummy"].Credits = 1000

		stage.AddPlayer(NewPlayer("testPlayer", Blue))
		stage.Players["testPlayer"].Credits = 124
		// give testPlayer some pixels
		stage.SetOwner(0, 0, stage.Players["testPlayer"])
		stage.SetOwner(0, 1, stage.Players["testPlayer"])
		stage.SetOwner(0, 2, stage.Players["testPlayer"])
		stage.SetOwner(1, 0, stage.Players["testPlayer"])

		if err := stage.SpreadPlayer(stage.Players["testPlayer"], stage.Players["dummy"]); err != nil {
			t.Errorf("Spread should have succeeded, but failed with error: %s", err)
		}

		if stage.Players["testPlayer"].Credits != 0 {
			t.Errorf("Test player should have had %d credits but had %d", 0, stage.Players["testPlayer"].Credits)
		}

		if stage.Players["dummy"].Credits != 938 {
			t.Errorf("Dummy player should have had %d credits but had %d", 938, stage.Players["dummy"].Credits)
		}

		if stage.Players["testPlayer"].TerritorySize != 10 {
			t.Errorf("Test player should have had %d pixels but had %d", 10, stage.Players["testPlayer"].TerritorySize)
		}

		if stage.Players["dummy"].TerritorySize != 90 {
			t.Errorf("Dummy player should have had %d pixels but had %d", 90, stage.Players["dummy"].TerritorySize)
		}
	}
}

func TestTick(t *testing.T) {
	height := 10
	width := 10
	stage := NewStage(height, width)
	stage.AddPlayer(NewPlayer("testPlayer", Blue))
	stage.AddPlayer(NewPlayer("testPlayer2", Red))

	// give testPlayer some pixels
	stage.SetOwner(0, 0, stage.Players["testPlayer"])
	stage.SetOwner(0, 1, stage.Players["testPlayer"])
	stage.SetOwner(0, 2, stage.Players["testPlayer"])
	stage.SetOwner(0, 3, stage.Players["testPlayer"])

	// give testPlayer2 some pixels
	stage.SetOwner(9, 5, stage.Players["testPlayer2"])
	stage.SetOwner(9, 6, stage.Players["testPlayer2"])
	stage.SetOwner(9, 7, stage.Players["testPlayer2"])
	stage.SetOwner(9, 8, stage.Players["testPlayer2"])
	stage.SetOwner(9, 9, stage.Players["testPlayer2"])

	if stage.Players["testPlayer"].Credits != 0 {
		t.Errorf("Test player should have had %d credits but had %d", 0, stage.Players["testPlayer"].Credits)
	}

	if stage.Players["testPlayer2"].Credits != 0 {
		t.Errorf("Test player should have had %d credits but had %d", 0, stage.Players["testPlayer2"].Credits)
	}

	stage.Tick()

	if stage.Players["testPlayer"].Credits != 4 {
		t.Errorf("Test player should have had %d credits but had %d", 4, stage.Players["testPlayer"].Credits)
	}

	if stage.Players["testPlayer2"].Credits != 5 {
		t.Errorf("Test player2 should have had %d credits but had %d", 5, stage.Players["testPlayer2"].Credits)
	}
}
