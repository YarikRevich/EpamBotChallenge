package tickets

import (
	"battlecity_test/game"
)

const (
	//Death tickets

	ENEMY  = 100
	BULLET = 200

	DEFAULT = 200
)

type Ticket struct {
	board *game.Board

	amount int
}

func (t *Ticket) EstimateMan(modifiers ...int){
	for _, mod := range modifiers{
		t.amount += mod
	}
}

func (t *Ticket) EstimateAuto(point game.Point) {
	if t.board.IsBulletAt(point) {
		t.amount += BULLET
	}
	
}

func (t *Ticket) GetAmount() int {
	return t.amount
}

func New(board *game.Board) *Ticket {
	return &Ticket{
		board: board}
}
