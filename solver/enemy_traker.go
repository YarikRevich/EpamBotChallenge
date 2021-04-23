package solver

import (
	"battlecity_test/game"
	"math"
)

func getTheNearestEnemy(c []game.Point, o game.Point) game.Point {
	var theBest game.Point
	for _, v := range c {
		if theBest.X == 0 && theBest.Y == 0 {
			theBest = v
			continue
		}
		if int(math.Abs(float64(o.X-v.X))) < theBest.X && int(math.Abs(float64(o.Y-v.Y))) < theBest.Y {
			theBest = v
		}
	}
	return theBest
}

func ifEnemyInAvailableZone(c game.Point) bool {
	if c.X != 0 && c.Y != 0 {
		z := getCurrentZone(c)
		return z == SECOND_ZONE_CODE || z == THIRD_ZONE_CODE
	}
	return false
}

func getTacticToGetTheEnemy(o, c game.Point) int {
	switch {
	case o.Y == c.Y && o.X < c.X:
		return FIRST_TACTIC_ENEMY
	case o.Y == c.Y && o.X > c.X:
		return SECOND_TACTIC_ENEMY
	case o.X <= c.X && o.Y <= c.Y:
		return THIRD_TACTIC_ENEMY
	case o.X >= c.X && o.Y <= c.Y:
		return FORTH_TACTIC_ENEMY
	case o.X <= c.X && o.Y >= c.Y:
		return FIFTH_TACTIC_ENEMY
	case o.X >= c.X && o.Y >= c.Y:
		return SIXTH_TACTIC_ENEMY
	}
	return ZERO_VALUE
}
