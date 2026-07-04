package rules

import model "Tic_Tac_toe/m/Model"

type Rules interface {
	IsValidMove(board model.Board,row int,col int) bool
	CheckWin(b model.Board, s model.Symbol) bool
	CheckDraw(b model.Board) bool
}