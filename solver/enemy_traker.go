package solver

import (
	"battlecity_test/game"
	"math"
)

const (
	specleScanRadius float64 = 5
)

var (
	specle game.Point

	ZERO_SPECLE = game.Point{}
)

func getEnemies(b *game.Board) []game.Point {
	return b.GetAllPoints(game.AI_TANK_DOWN, game.AI_TANK_LEFT, game.AI_TANK_RIGHT, game.AI_TANK_UP, game.OTHER_TANK_DOWN, game.OTHER_TANK_LEFT, game.OTHER_TANK_RIGHT, game.OTHER_TANK_UP)
}

func specleIsStart()bool{
	return specle == ZERO_SPECLE
}

func getSpecle()game.Point{
	return specle
}

func setSpecle(c game.Point) {
	specle =  c
}

func clearSpecle() {
	specle = game.Point{}
}

func isEnemyAlive(e []game.Point)(game.Point, bool) {
	for _, v := range e{
		if math.Sqrt((math.Pow(float64(specle.X - v.X), 2) + math.Pow(float64(specle.Y - v.Y), 2))) <= specleScanRadius{
			return v, true
		}
	}
	return game.Point{}, false
}

func getTheNearestEnemy(c []game.Point, o game.Point) game.Point {
	var theBest game.Point
	var theBestLength float64
	
	for _, v := range c {
		if theBest.X == 0 && theBest.Y == 0 {
			theBest = v
			theBestLength = math.Sqrt((math.Pow(float64(v.X - o.X), 2) + math.Pow(float64(v.Y - o.Y), 2)))
			continue
		}
		if n := math.Sqrt((math.Pow(float64(v.X - o.X), 2) + math.Pow(float64(v.Y - o.Y), 2))); n < theBestLength{
			theBest = v
			theBestLength = n
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
