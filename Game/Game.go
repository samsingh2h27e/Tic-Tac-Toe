package game

import (
	model "Tic_Tac_toe/m/Model"
	notification "Tic_Tac_toe/m/Notification"
	rules "Tic_Tac_toe/m/Rules"
	"container/list"
)

type Game struct {
	B model.Board
	Deque list.List
	Rule rules.Rules
	Observers []notification.Observer
	GameOver bool
}

func (g *Game) AddPlayers(p model.Player) {
	g.Deque.PushBack(p)
}

func (g *Game) MakeGameOver() {
	g.GameOver = true
}

func (g *Game) GetCurrentPlayer() model.Player {
	player := g.Deque.Front()
	g.Deque.Remove(player)
	return player.Value.(model.Player)
	
}