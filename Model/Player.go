package model

import "fmt"

type Player struct {
	Id    int
	Name  string
	S     Symbol
	Score int
}

func (p *Player) Move() (int, int) {
	var x int
	var y int
	fmt.Println("Player: "+ p.Name + "enter row and column:")
	fmt.Scan(&x,&y)
	return x,y
}