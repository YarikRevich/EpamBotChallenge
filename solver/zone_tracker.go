package solver

import (
	"battlecity_test/game"
	"math/rand"
)

const (
	//Available zone codes ...
	FIRST_ZONE_CODE = iota
	SECOND_ZONE_CODE
	THIRD_ZONE_CODE
	FORTH_ZONE_CODE

	//Available tactics of movement ...
	FIRST_TACTIC
	SECOND_TACTIC
	THIRD_TACTIC
	FORTH_TACTIC

	//Zero value for not chosen action ...
	ZERO_VALUE

	//Available actions for set of them ...
	UP    = "UP"
	RIGHT = "RIGHT"
	LEFT  = "LEFT"
	DOWN  = "DOWN"
)

var (
	FIRST_ZONE_EDGES  = []game.Point{{X: 16, Y: 0}, {X: 32, Y: 16}}
	SECOND_ZONE_EDGES = []game.Point{{X: 16, Y: 32}, {X: 0, Y: 16}}
	THIRD_ZONE_EDGES  = []game.Point{{X: 16, Y: 32}, {X: 32, Y: 16}}
	FORTH_ZONE_EDGES  = []game.Point{{X: 16, Y: 0}, {X: 0, Y: 16}}

	//Control points of each zone ...
	FIRST_ZONE_CONTROL_POINT  = game.Point{X: 32, Y: 0}
	SECOND_ZONE_CONTROL_POINT = game.Point{X: 0, Y: 32}
	THIRD_ZONE_CONTROL_POINT  = game.Point{X: 32, Y: 32}
	FORTH_ZONE_CONTROL_POINT  = game.Point{X: 0, Y: 0}

	//Current tactic of movement

	CURRENT_TACTIC = ZERO_VALUE
)

//Shows the zone where the hero is placed ...
func getCurrentZone(c game.Point) int {
	switch {
	case c.X >= FIRST_ZONE_EDGES[0].X && c.Y <= FIRST_ZONE_EDGES[1].Y:
		return FIRST_ZONE_CODE
	case c.X <= SECOND_ZONE_EDGES[0].X && c.Y >= SECOND_ZONE_EDGES[1].Y:
		return SECOND_ZONE_CODE
	case c.X >= THIRD_ZONE_EDGES[0].X && c.Y >= THIRD_ZONE_EDGES[1].Y:
		return THIRD_ZONE_CODE
	case c.X <= FORTH_ZONE_EDGES[0].X && c.Y <= FORTH_ZONE_EDGES[1].Y:
		return FORTH_ZONE_CODE
	}
	return ZERO_VALUE
}

//Checks if hero at the control point ...
func gotTheLastPointOfZone(c game.Point) bool {
	return c == FIRST_ZONE_CONTROL_POINT ||
		c == SECOND_ZONE_CONTROL_POINT ||
		c == THIRD_ZONE_CONTROL_POINT ||
		c == FORTH_ZONE_CONTROL_POINT
}

func getTacticByCodeZone(zone int) int {
	switch zone {
	case FIRST_ZONE_CODE:
		return FIRST_TACTIC
	case SECOND_ZONE_CODE:
		return SECOND_TACTIC
	case THIRD_ZONE_CODE:
		return THIRD_TACTIC
	case FORTH_ZONE_CODE:
		return FORTH_TACTIC
	}
	return ZERO_VALUE
}

//Due to the previous tactic returns the next one ...
func getNextTactic() int {
	switch CURRENT_TACTIC {
	case FIRST_TACTIC:
		return SECOND_TACTIC
	case SECOND_TACTIC:
		return THIRD_TACTIC
	case THIRD_TACTIC:
		return FORTH_TACTIC
	}
	return ZERO_VALUE
}

func setNextTactic(tactic int) {
	CURRENT_TACTIC = tactic
}

//Due to the tactic name returns the action set
func getActionSetByTactic() []string {
	switch CURRENT_TACTIC {
	case FIRST_TACTIC:
		return []string{UP, LEFT, RIGHT, DOWN}
	case SECOND_TACTIC:
		return []string{UP, RIGHT, LEFT, DOWN}
	case THIRD_TACTIC:
		return []string{DOWN, LEFT, UP}
	case FORTH_TACTIC:
		return []string{RIGHT, UP, DOWN}
	}
	return nil
}

func getRandomTactic() int {
	a := [...]int{FIRST_TACTIC, SECOND_TACTIC, THIRD_TACTIC, FORTH_TACTIC}
	for {
		c := a[rand.Intn(3)]
		if c != CURRENT_TACTIC{
			return c
		}
	}
}
