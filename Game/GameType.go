package game

type GameType int

const (
	STANDARD GameType = iota
)

func (g *GameType) String () string {
	return "standard"
}