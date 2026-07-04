package game

import (
	model "Tic_Tac_toe/m/Model"
	notification "Tic_Tac_toe/m/Notification"
	rules "Tic_Tac_toe/m/Rules"
	"container/list"
	"fmt"
	"log/slog"
)

type Game struct {
	B model.Board
	Deque *list.List
	Rule rules.Rules
	Observers []notification.Observer
	GameOver bool
}

func (g *Game) AddPlayers(p model.Player) {
	slog.Info("Received player",
		"id", p.Id,
		"name", p.Name,
		"player", p,
	)

	e := g.Deque.PushBack(p)
	slog.Info("Stored",
		"value", e.Value,
		"type", fmt.Sprintf("%T", e.Value),
	)
}

func (g *Game) MakeGameOver() {
	g.GameOver = true
}

func (g *Game) GetCurrentPlayer() model.Player {
	slog.Info("Entered in current player method")
	player := g.Deque.Front()
	g.Deque.Remove(player)
	return player.Value.(model.Player)
	
}