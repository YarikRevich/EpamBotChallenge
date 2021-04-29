package utils

import (
	"battlecity_test/game"
	"fmt"
	"math"
	"math/rand"
)

const (
	BULLET_SPEED = 2
)

var (
	Bulets [][]game.Point
)

func GetAIEnemies(b *game.Board) []game.Point {
	return b.GetAllPoints(
		game.AI_TANK_DOWN,
		game.AI_TANK_LEFT,
		game.AI_TANK_RIGHT,
		game.AI_TANK_UP,
		game.AI_TANK_PRIZE,

		game.PRIZE,
		game.PRIZE_BREAKING_WALLS,
		game.PRIZE_IMMORTALITY,
		game.PRIZE_NO_SLIDING,
		game.PRIZE_VISIBILITY,
		game.PRIZE_WALKING_ON_WATER,
	)
}

func GetWalls(b *game.Board) []game.Point {
	return b.GetAllPoints(
		game.WALL,
		game.WALL_DESTROYED_DOWN,
		game.WALL_DESTROYED_UP,
		game.WALL_DESTROYED_LEFT,
		game.WALL_DESTROYED_RIGHT,

		game.WALL_DESTROYED_DOWN_TWICE,
		game.WALL_DESTROYED_UP_TWICE,
		game.WALL_DESTROYED_LEFT_TWICE,
		game.WALL_DESTROYED_RIGHT_TWICE,

		game.WALL_DESTROYED_LEFT_RIGHT,
		game.WALL_DESTROYED_UP_DOWN,

		game.WALL_DESTROYED_UP_LEFT,
		game.WALL_DESTROYED_RIGHT_UP,
		game.WALL_DESTROYED_DOWN_LEFT,
		game.WALL_DESTROYED_DOWN_RIGHT,
	)
}

func CheckEqual(a, b []game.Point) bool {
	for _, av := range a {
		for _, bv := range b {
			if av != bv {
				return false
			}
		}
	}
	return true
}

func IsBattleWall(a game.Point, b []game.Point) bool {
	for _, v := range b {
		if a == v {
			return true
		}
	}
	return false
}

func IsWithin(a game.Point, b []game.Point) bool {
	//Checks whether a within b

	for _, v := range b {
		if a == v {
			return true
		}
	}
	return false

}

func IsWithinPrecision(a game.Point, b []game.Point, p int) bool {
	//Checks whether a within b, but uses the precision

	for _, v := range b {
		if a == v {
			return true
		}
	}
	return false
}

func GetTheNearestElement(o game.Point, c []game.Point) game.Point {
	theBest := c[rand.Intn(len(c)-1)]
	theBestLength := math.Sqrt((math.Pow(float64(theBest.X-o.X), 2) + math.Pow(float64(theBest.Y-o.Y), 2)))

	for _, v := range c {
		if n := math.Sqrt((math.Pow(float64(v.X-o.X), 2) + math.Pow(float64(v.Y-o.Y), 2))); n < theBestLength {
			theBest = v
			theBestLength = n
		}
	}
	return theBest
}

func IsClearWayOut(o game.Point, c game.Point, b *game.Board)bool {
	return true
}

func GetTheNearestWayOutWall(o game.Point, c []game.Point, b *game.Board) game.Point {
	theBest := c[rand.Intn(len(c)-1)]
	theBestLength := math.Sqrt((math.Pow(float64(theBest.X-o.X), 2) + math.Pow(float64(theBest.Y-o.Y), 2)))

	for _, v := range c {
		if n := math.Sqrt((math.Pow(float64(v.X-o.X), 2) + math.Pow(float64(v.Y-o.Y), 2))); (n < theBestLength) &&  IsClearWayOut(o, v, b){
			theBest = v
			theBestLength = n
		}
	}
	return theBest
}

func IsUpdatingProcess(e []game.Point) bool {

	if (IsWithin(game.Point{X: 1, Y: 1}, e) &&
		IsWithin(game.Point{X: 25, Y: 1}, e)) ||
		(IsWithin(game.Point{X: 2, Y: 32}, e) &&
			IsWithin(game.Point{X: 31, Y: 32}, e)) ||
		(IsWithin(game.Point{X: 4, Y: 1}, e) &&
			IsWithin(game.Point{X: 20, Y: 1}, e)) {
		return true
	}
	return false
}

func IsBulletAlive(specle game.Point, b []game.Point) (game.Point, bool) {
	for _, v := range b {
		fmt.Println(v)
		if math.Sqrt((math.Pow(float64(specle.X-v.X), 2) + math.Pow(float64(specle.Y-v.Y), 2))) <= BULLET_SPEED+1 {
			return v, true
		}
	}
	return game.Point{}, false
}

func IsElementEnemy(a game.Element) bool {
	return a == game.AI_TANK_DOWN ||
		a == game.AI_TANK_LEFT ||
		a == game.AI_TANK_PRIZE ||
		a == game.AI_TANK_RIGHT ||
		a == game.AI_TANK_UP ||
		a == game.OTHER_TANK_DOWN ||
		a == game.OTHER_TANK_LEFT ||
		a == game.OTHER_TANK_RIGHT ||
		a == game.OTHER_TANK_UP
}

func GetAvailableZoneToGo(b *game.Board) []game.Point {
	return b.GetAllPoints(game.NONE, game.TREE, game.ICE, game.PRIZE, game.PRIZE_BREAKING_WALLS, game.PRIZE_IMMORTALITY, game.PRIZE_NO_SLIDING, game.PRIZE_VISIBILITY, game.PRIZE_WALKING_ON_WATER)
}

func ElementIs(a game.Point, b game.Element, c *game.Board) bool {
	return c.GetAt(a) == b
}