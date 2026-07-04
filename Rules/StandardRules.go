package rules

import model "Tic_Tac_toe/m/Model"

type StandardRules struct {
}

func (sr *StandardRules) IsValidMove(board model.Board, row int, col int) bool {
	if len(board.Grid) <= row || len(board.Grid[0]) <= col || row < 0 || col < 0 || !board.IsCellEmpty(row,col) {
		return false
	}
	return true
}
func (sr *StandardRules) CheckWin(b model.Board, s model.Symbol) bool {
	length := len(b.Grid)
	for i := 0; i < length; i++ {
		if b.Grid[i][i] != s {
			return false
		}
	}
	for i := 0; i < length; i++ {
		if b.Grid[i][0] != s {
			return false
		} 
	}
	for i := 0;i < length;i++ {
		if b.Grid[0][i] != s {
			return false
		}
	}
	for i := 0;i < length;i++ {
		if b.Grid[2][i] != s {
			return false
		}
	}
	for i := 0;i < length;i++ {
		if b.Grid[i][2] != s {
			return false
		} 
	}
	for i := 0;i<length;i++ {
		if b.Grid[2-i][i] != s {
			return false
		}
	}
	return true
}
func (sr *StandardRules) CheckDraw(b model.Board) bool {
	length := len(b.Grid)
	for i := 0;i < length ;i++ {
		for j := 0;j < length;j++ {
			if b.Grid[i][j] == b.EmptyCell {
				return false
			}
		}
	} 
	return true
}