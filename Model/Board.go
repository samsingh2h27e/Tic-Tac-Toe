package model

import (
	"fmt"
	"log/slog"
)

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
	slog.Info("Displaying the board length","length",len(b.Grid))
	for _, row := range b.Grid {
		for _, val := range row {
			fmt.Printf("%v ",val.GetMask())
		}
		fmt.Println()
	}
}

func (b *Board) ResetBoard() {
	b.Grid = make([][]Symbol, b.Size)
    for i := 0; i < b.Size; i++ {
        b.Grid[i] = make([]Symbol, b.Size)

        for j := 0; j < b.Size; j++ {
            b.Grid[i][j] = Symbol{Mask: "_"}
        }
    }
}