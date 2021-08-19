package tickets

import (
	"battlecity_test/game"
)

const (
	/**
		Table of tickets
		# Negative
			1) BULLET: -200

		# Positive
			1) DEFAULT: 200 (Points for default ways determined by algorithm)
			2) LIVE_ENEMY: 100
			3) AI_ENEMY: 50

		#Helpers
			1) SUICIDE = 1000
	**/

	BULLET = -200

	DEFAULT    = 200
	LIVE_ENEMY = 100
	AI_ENEMY   = 50
)

type Ticket struct {
	board *game.Board

	amount int
}

func (t *Ticket) EstimateMan(modifiers ...int) {
	for _, mod := range modifiers {
		t.amount += mod
	}
}

func (t *Ticket) EstimateAuto(point game.Point) {
	if t.board.IsBulletAt(point) {
		t.amount += BULLET
	}

	if t.board.IsAt(point, game.AI_TANK_DOWN) ||
		t.board.IsAt(point, game.AI_TANK_LEFT) ||
		t.board.IsAt(point, game.AI_TANK_PRIZE) ||
		t.board.IsAt(point, game.AI_TANK_RIGHT) ||
		t.board.IsAt(point, game.AI_TANK_UP) {
		t.amount += AI_ENEMY
	}

	if t.board.IsAt(point, game.TANK_DOWN) ||
		t.board.IsAt(point, game.TANK_LEFT) ||
		t.board.IsAt(point, game.TANK_RIGHT) ||
		t.board.IsAt(point, game.TANK_UP) {
		t.amount += LIVE_ENEMY
	}
}

func (t *Ticket) EstimateSuicideWay() {

}

func (t *Ticket) GetAmount() int {
	return t.amount
}

func New(board *game.Board) *Ticket {
	return &Ticket{
		board: board}
}
