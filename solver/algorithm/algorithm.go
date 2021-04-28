package algorithm

import (
	"math/rand"

	"battlecity_test/game"
	"battlecity_test/solver/algorithm/bfs"
	"battlecity_test/solver/algorithm/graph"
	"battlecity_test/solver/utils"
)

const (
	ZERO_TACTIC = ""

	UP    = "UP"
	RIGHT = "RIGHT"
	LEFT  = "LEFT"
	DOWN  = "DOWN"
)

var (
	EMPTY_COORDS = game.Point{}

	MY_COORDS = game.Point{}

	MY_BULLET = game.Point{}
)

func isFreeAt(c game.Point, a []game.Point) bool {
	for _, v := range a {
		if c == v {
			return true
		}
	}
	return false
}

func getPointsOfBulletsAround(b []game.Point) []game.Point{
	r := []game.Point{}
	for _, v := range b{
		r = append(
			r, 
			game.Point{X: v.X, Y: v.Y+2}, 
			game.Point{X: v.X+2, Y: v.Y}, 
			game.Point{X: v.X-2, Y: v.Y}, 
			game.Point{X: v.X, Y: v.Y-2}, 
		)
	}
	return r
}

func updateMyCoords(c game.Point) {
	MY_COORDS = c
}

func getTheAroundCoordsOfDestination(d game.Point) []game.Point {
	return []game.Point{
		{X: d.X, Y: d.Y + 1},
		{X: d.X + 1, Y: d.Y},
		{X: d.X - 1, Y: d.Y},
		{X: d.X, Y: d.Y - 1},
		d,
	}
}

//Just creates the graph ;) ...
func createGraph(c game.Point, a []game.Point, b []game.Point) *graph.Graph {

	g := graph.New(2000)
	allBullets := getPointsOfBulletsAround(b)

	for _, v := range a {

		top := game.Point{X: v.X, Y: v.Y + 1}
		right := game.Point{X: v.X + 1, Y: v.Y}
		left := game.Point{X: v.X - 1, Y: v.Y}
		bottom := game.Point{X: v.X, Y: v.Y - 1}

		if isFreeAt(top, a) && !utils.IsWithin(top, allBullets) {
			g.Connect(v, top)
		}

		if isFreeAt(right, a) && !utils.IsWithin(right, allBullets) {
			g.Connect(v, right)
		}

		if isFreeAt(left, a) && !utils.IsWithin(left, allBullets) {
			g.Connect(v, left)
		}

		if isFreeAt(bottom, a) && !utils.IsWithin(bottom, allBullets) {
			g.Connect(v, bottom)
		}
	}

	return g
}

//Analises the graph and checks if the hero in the trap ...
func analiseGraph(g *graph.Graph, myCoords game.Point, destination game.Point, b *game.Board) ([]game.Point, bool) {
	var trap bool

	r := bfs.New(g, myCoords)

	// d := getTheAroundCoordsOfDestination(destination)
	path := r.Path(destination)

	if path == nil {
		trap = true

		a := b.GetBarriers()
		a = append(a, myCoords)

		g = createGraph(myCoords, a, b.GetBullets())
		r = bfs.New(g, myCoords)

		path = r.Path(utils.GetTheNearestElement(myCoords, utils.GetWalls(b)))
	}

	return path, trap
}

func GetBestTactic(myCoords game.Point, myBullet game.Point, destination game.Point, b *game.Board) (string, bool) {

	if myCoords == EMPTY_COORDS {
		if MY_COORDS != EMPTY_COORDS {
			myCoords = MY_COORDS
		} else {
			myCoords = game.Point{X: rand.Intn(b.BoardSize()), Y: rand.Intn(b.BoardSize())}
		}
	}

	a := b.GetAllPoints(game.NONE, game.TREE, game.ICE, game.PRIZE, game.PRIZE_IMMORTALITY, game.PRIZE_BREAKING_WALLS, game.PRIZE_VISIBILITY, game.PRIZE_NO_SLIDING, game.PRIZE_WALKING_ON_WATER, game.OTHER_TANK_DOWN, game.OTHER_TANK_LEFT, game.OTHER_TANK_RIGHT, game.OTHER_TANK_UP, game.AI_TANK_DOWN, game.AI_TANK_LEFT, game.AI_TANK_RIGHT, game.AI_TANK_UP, game.AI_TANK_PRIZE)
	a = append(a, myCoords, myBullet)

	g := createGraph(myCoords, a, b.GetBullets())

	path, trap := analiseGraph(g, myCoords, destination, b)

	if len(path) <= 1 {
		return ZERO_TACTIC, trap
	}

	top := game.Point{X: myCoords.X, Y: myCoords.Y + 1}
	right := game.Point{X: myCoords.X + 1, Y: myCoords.Y}
	left := game.Point{X: myCoords.X - 1, Y: myCoords.Y}
	bottom := game.Point{X: myCoords.X, Y: myCoords.Y - 1}

	switch path[1] {
	case top:
		updateMyCoords(top)
		return UP, trap
	case right:
		updateMyCoords(right)
		return RIGHT, trap
	case left:
		updateMyCoords(left)
		return LEFT, trap
	case bottom:
		updateMyCoords(bottom)
		return DOWN, trap
	default:
		return ZERO_TACTIC, trap
	}
}
