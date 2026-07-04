	package game

	import (
		model "Tic_Tac_toe/m/Model"
		rules "Tic_Tac_toe/m/Rules"
		"container/list"
		"log/slog"
	)

	type GameFactory struct{}

	func (g *GameFactory) CreateGame(gameType GameType, size int) *Game {
		slog.Info("Entered createGame")
		game := Game{}
		game.B = model.Board{}
		game.B.Size = size
		game.B.ResetBoard()
		slog.Info("board is resetted")
		game.B.Display()
		game.B.EmptyCell = model.Symbol{Mask: "_"}
		game.Deque = list.New()
		game.GameOver = false
		game.Rule = &rules.StandardRules{}
		return &game
	}