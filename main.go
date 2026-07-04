package main

import (
	game "Tic_Tac_toe/m/Game"
	model "Tic_Tac_toe/m/Model"
	"fmt"
	"log/slog"
)

func main() {
	slog.Info("Starting game")
	gameFactory := game.GameFactory{}
	game := gameFactory.CreateGame(game.STANDARD, 3)
	player1 := model.Player{
		Id:   1,
		Name: "Sambhav",
		S: model.Symbol{
			Mask: "X",
		},
	}
	player2 := model.Player{
		Id:   2,
		Name: "Shekhar",
		S: model.Symbol{
			Mask: "O",
		},
	}
	slog.Info("Adding players")
	game.AddPlayers(player1)
	game.AddPlayers(player2)
	for !game.GameOver {
		player := game.GetCurrentPlayer()
		slog.Info("Current ","player",player.Name)
		x,y := player.Move()
		slog.Info("Player move","x",x,"y",y)
		if game.Rule.IsValidMove(game.B,x,y) {
			game.B.PlaceMask(x,y,player.S)
			game.B.Display()
			if game.Rule.CheckDraw(game.B) {
				fmt.Println("game is draw")
				game.B.Display()
				game.MakeGameOver()
				continue
			}	
			if game.Rule.CheckWin(game.B,player.S) {
				fmt.Println("Player: " + player.Name + " wins")
				game.B.Display()
				game.MakeGameOver()
				continue
			}
			game.AddPlayers(player)
		}else {
			fmt.Println("Move is invalid")
			player2 := game.GetCurrentPlayer()
			game.AddPlayers(player)
			game.AddPlayers(player2)
		}
	}
}
