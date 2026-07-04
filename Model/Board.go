package model

import "fmt"

type Board struct {
	Grid      [][]Symbol
	Size      int
	EmptyCell Symbol
}

func (b *Board) IsCellEmpty(row int, col int) bool {
	return b.EmptyCell == b.Grid[row][col]
}

func (b *Board) PlaceMask(row int, col int, s Symbol) {
	b.Grid[row][col] = s
}

func (b *Board) GetCell(row int, col int) Symbol {
	return b.Grid[row][col]
}

func (b *Board) Display() {
	for _, row := range b.Grid {
		for _, val := range row {
			fmt.Printf("%v ",val)
		}
		fmt.Println()
	}
}

func (b *Board) ResetBoard() {
	for i,row := range b.Grid {
		for j := range row {
			b.Grid[i][j] = Symbol{Mask: "_"}
		}
	}
}