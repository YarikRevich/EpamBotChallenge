package solver

import (
	"fmt"
	"battlecity_test/game"
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

	//Available tactics of movements to get the enemies ...

	FIRST_TACTIC_ENEMY
	SECOND_TACTIC_ENEMY
	THIRD_TACTIC_ENEMY
	FORTH_TACTIC_ENEMY
	FIFTH_TACTIC_ENEMY
	SIXTH_TACTIC_ENEMY

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

	//Not available zone

	NOT_AVAILABLE_ZONE           = game.Point{X: 0, Y: 14}
	IS_NOT_AVAILABLE_ZONE_ACTIVE bool
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

//Checks whether hero is stuck ...
func gotStuck(c game.Point) bool {
	fmt.Println(Steps)
	if len(Steps) > 2 {
		return Steps[len(Steps)-1].X == Steps[len(Steps)-2].X+1 ||
			Steps[len(Steps)-1].X == Steps[len(Steps)-2].X-1 ||
			Steps[len(Steps)-1].X == Steps[len(Steps)-2].Y+1 ||
			Steps[len(Steps)-1].X == Steps[len(Steps)-2].Y-1
	}
	return false
}

func notDangerous(e game.Element) bool {
	return (e == game.NONE || e == game.ICE || e == game.OTHER_TANK_DOWN || e == game.OTHER_TANK_LEFT || e == game.OTHER_TANK_RIGHT || e == game.OTHER_TANK_UP) &&
		e != game.TREE &&
		e != game.RIVER &&
		e != game.BULLET
}

// //Checks if hero at the control point ...
// func gotTheLastPointOfZone(c game.Point) bool {
// 	return c == FIRST_ZONE_CONTROL_POINT ||
// 		c == SECOND_ZONE_CONTROL_POINT ||
// 		c == THIRD_ZONE_CONTROL_POINT ||
// 		c == FORTH_ZONE_CONTROL_POINT
// }
