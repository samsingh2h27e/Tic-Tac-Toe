package game

import (
	model "Tic_Tac_toe/m/Model"
	rules "Tic_Tac_toe/m/Rules"
	"container/list"
)

type GameFactory struct{}

func (g *GameFactory) CreateGame(gameType GameType, size int) Game {
	game := Game{}
	game.B = model.Board{}
	game.B.ResetBoard()
	game.B.Size = 3
	game.B.EmptyCell = model.Symbol{Mask: "_"}
	game.Deque = *list.New()
	game.GameOver = false
	game.Rule = &rules.StandardRules{}
	return game
}