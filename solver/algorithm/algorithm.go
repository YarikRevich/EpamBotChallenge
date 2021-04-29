package algorithm

import (
	"fmt"
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

func updateMyCoords(c game.Point) {
	MY_COORDS = c
}

//Just creates the graph ;) ...
func createGraph(c game.Point, a []game.Point, b []game.Point) *graph.Graph {

	g := graph.New(2000)
	// allBullets := getPointsOfBulletsAround(b)

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
// func getPath(g *graph.Graph, myCoords game.Point, destination game.Point, b *game.Board) []game.Point {

// 	return path
// }

func exceptBullets(a []game.Point, b []game.Point) []game.Point {
	for _, bullet := range b {
		top := game.Point{X: bullet.X, Y: bullet.Y + 2}
		right := game.Point{X: bullet.X + 2, Y: bullet.Y}
		left := game.Point{X: bullet.X - 2, Y: bullet.Y}
		bottom := game.Point{X: bullet.X, Y: bullet.Y - 2}

		topNear := game.Point{X: bullet.X, Y: bullet.Y + 1}
		rightNear := game.Point{X: bullet.X + 1, Y: bullet.Y}
		leftNear := game.Point{X: bullet.X - 1, Y: bullet.Y}
		bottomNear := game.Point{X: bullet.X, Y: bullet.Y - 1}

		for i, v := range a {
			if v == top ||
				v == right ||
				v == left ||
				v == bottom ||
				v == bullet ||
				v == topNear ||
				v == rightNear ||
				v == leftNear ||
				v == bottomNear {
				if len(a) > i {
					a = append(a[:i], a[i+1:]...)
				} else {
					a = append(a[:i], a[len(a)-1:]...)
				}
			}
		}
	}
	return a
}

func GetBestTactic(myCoords game.Point, myBullet game.Point, destination game.Point, b *game.Board) string {

	if myCoords == EMPTY_COORDS {
		if MY_COORDS != EMPTY_COORDS {
			myCoords = MY_COORDS
		} else {
			myCoords = game.Point{X: rand.Intn(b.BoardSize()), Y: rand.Intn(b.BoardSize())}
		}
	}

	a := b.GetAllPoints(utils.GetAvailableElements(b)...)
	// a = exceptBullets(a, b.GetBullets())
	a = append(a, myCoords, myBullet)

	g := createGraph(myCoords, a, b.GetBullets())

	r := bfs.New(g, myCoords)
	path := r.Path(destination)

	if path == nil {
		//IF IT IS A TRAP IT CREATES A NEW GRAPH AND FIND THE WAYOUT ...
		fmt.Println("TRAP")
		// a := b.GetAllPoints(append(utils.GetAvailableElements(b), utils.GetWallsElements(b)...)...)

		// g := createGraph(myCoords, a, b.GetBullets())

		// r := bfs.New(g, myCoords)
		// path = r.Path(utils.GetTheNearestElement(myCoords, utils.GetWalls(b), b))
	}

	if len(path) <= 1 {
		return ZERO_TACTIC
	}

	top := game.Point{X: myCoords.X, Y: myCoords.Y + 1}
	right := game.Point{X: myCoords.X + 1, Y: myCoords.Y}
	left := game.Point{X: myCoords.X - 1, Y: myCoords.Y}
	bottom := game.Point{X: myCoords.X, Y: myCoords.Y - 1}

	switch path[1] {
	case top:
		updateMyCoords(top)
		return UP
	case right:
		updateMyCoords(right)
		return RIGHT
	case left:
		updateMyCoords(left)
		return LEFT
	case bottom:
		updateMyCoords(bottom)
		return DOWN
	default:
		return ZERO_TACTIC
	}
}
