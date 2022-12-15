package game

type Stage struct {
	Height  int
	Width   int
	Board   [][]*Pixel
	Players map[string]*Player
}

func NewStage(height int, width int) Stage {

	players := make(map[string]*Player)
	dummyPlayer := NewPlayer("dummy", White)
	players["dummy"] = dummyPlayer

	board := make([][]*Pixel, height)

	for i := range board {
		board[i] = make([]*Pixel, width)
		for j := range board[i] {
			board[i][j] = &Pixel{Owner: dummyPlayer, Row: i, Col: j}
			dummyPlayer.TerritorySize++
		}
	}

	return Stage{
		Height:  height,
		Width:   width,
		Board:   board,
		Players: players,
	}
}

func (s *Stage) GetOwner(row int, col int) *Player {
	return s.Board[row][col].Owner
}

func (s *Stage) SetOwner(row int, col int, owner *Player) {
	s.Board[row][col].Owner.TerritorySize--

	s.Board[row][col].Owner = owner
	s.Board[row][col].Owner.TerritorySize++
}

func (s *Stage) AddPlayer(player *Player) {
	s.Players[player.Name] = player
}

type Pixel struct {
	Owner *Player

	Row int
	Col int
}

type Player struct {
	Name          string
	Color         Color
	Credits       int
	TerritorySize int
}

func NewPlayer(name string, color Color) *Player {
	return &Player{Name: name, Color: color}
}

type Color int64

const (
	Black Color = iota
	Red
	Green
	Blue
	White
)

func (c Color) String() string {
	switch c {
	case Black:
		return "black"
	case Red:
		return "red"
	case Green:
		return "green"
	case Blue:
		return "blue"
	case White:
		return "white"
	default:
		return "unknown"
	}
}
