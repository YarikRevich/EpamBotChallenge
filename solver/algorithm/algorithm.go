package algorithm

import (

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
)

func isFreeAt(c game.Point, a []game.Point) bool {
	for _, v := range a {
		if c == v {
			return true
		}
	}
	return false
}

func updateMyCoords(c game.Point) {
	MY_COORDS = c
}

//Just creates the graph ;) ...
func createGraph(c game.Point, a []game.Point) *graph.Graph {

	g := graph.New(600)

	for _, v := range a {

		top := game.Point{X: v.X, Y: v.Y + 1}
		right := game.Point{X: v.X + 1, Y: v.Y}
		left := game.Point{X: v.X - 1, Y: v.Y}
		bottom := game.Point{X: v.X, Y: v.Y - 1}

		if isFreeAt(top, a) {
			g.Connect(v, top)
		}

		if isFreeAt(right, a) {
			g.Connect(v, right)
		}

		if isFreeAt(left, a) {
			g.Connect(v, left)
		}

		if isFreeAt(bottom, a) {
			g.Connect(v, bottom)
		}
	}

	return g
}

//Analises the graph and checks if the hero in the trap ...
func analiseGraph(g *graph.Graph, myCoords game.Point, destination game.Point, b *game.Board)([]game.Point, bool) {
	var trap bool

	r := bfs.New(g, myCoords)

	path := r.Path(destination)
	if path == nil {

		trap = true

		a := b.GetBarriers()
		a = append(a, myCoords)

		g = createGraph(myCoords, a)
		r = bfs.New(g, myCoords)
	
		path = r.Path(utils.GetTheNearestElement(myCoords, utils.GetWalls(b)))
	}

	return path, trap
}

func GetBestTactic(myCoords game.Point, destination game.Point, b *game.Board) (string, bool) {

	if myCoords == EMPTY_COORDS {
		myCoords = MY_COORDS
	}

	a := b.GetAllPoints(game.NONE, game.TREE, game.ICE, game.PRIZE, game.PRIZE_IMMORTALITY, game.PRIZE_BREAKING_WALLS, game.PRIZE_VISIBILITY, game.PRIZE_NO_SLIDING, game.PRIZE_WALKING_ON_WATER, game.OTHER_TANK_DOWN, game.OTHER_TANK_LEFT, game.OTHER_TANK_RIGHT, game.OTHER_TANK_UP, game.AI_TANK_DOWN, game.AI_TANK_LEFT, game.AI_TANK_RIGHT, game.AI_TANK_UP, game.AI_TANK_PRIZE)
	a = append(a, myCoords)

	g := createGraph(myCoords, a)
	
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
