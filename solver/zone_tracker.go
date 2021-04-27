package solver

// import (

// 	"battlecity_test/game"
// 	"battlecity_test/solver/algorithm"
// 	"battlecity_test/solver/algorithm/graph"
// )

// const (
// 	//Available zone codes ...
// 	FIRST_ZONE_CODE = iota
// 	SECOND_ZONE_CODE
// 	THIRD_ZONE_CODE
// 	FORTH_ZONE_CODE

// 	//Zero value for not chosen action ...
// 	ZERO_VALUE

// 	//Zero value for tactic ...

// 	ZERO_TACTIC = ""

// 	//Available actions for set of them ...
// 	UP    = "UP"
// 	RIGHT = "RIGHT"
// 	LEFT  = "LEFT"
// 	DOWN  = "DOWN"
// )

// var (
// 	FIRST_ZONE_EDGES  = []game.Point{{X: 16, Y: 0}, {X: 32, Y: 16}}
// 	SECOND_ZONE_EDGES = []game.Point{{X: 16, Y: 32}, {X: 0, Y: 16}}
// 	THIRD_ZONE_EDGES  = []game.Point{{X: 16, Y: 32}, {X: 32, Y: 16}}
// 	FORTH_ZONE_EDGES  = []game.Point{{X: 16, Y: 0}, {X: 0, Y: 16}}

// 	//Control points of each zone ...
// 	FIRST_ZONE_CONTROL_POINT  = game.Point{X: 32, Y: 0}
// 	SECOND_ZONE_CONTROL_POINT = game.Point{X: 0, Y: 32}
// 	THIRD_ZONE_CONTROL_POINT  = game.Point{X: 32, Y: 32}
// 	FORTH_ZONE_CONTROL_POINT  = game.Point{X: 0, Y: 0}

// 	//Current tactic of movement

// 	CURRENT_TACTIC = ZERO_VALUE

// 	//Not available zone

// 	NOT_AVAILABLE_ZONE           = game.Point{X: 0, Y: 14}
// 	IS_NOT_AVAILABLE_ZONE_ACTIVE bool

// 	GRAPH = map[game.Point][]game.Point{}

// 	EMPTY_COORDS = game.Point{}

// 	MY_COORDS = game.Point{}
// )

// //Shows the zone where the hero is placed ...
// func getCurrentZone(c game.Point) int {
// 	switch {
// 	case c.X >= FIRST_ZONE_EDGES[0].X && c.Y <= FIRST_ZONE_EDGES[1].Y:
// 		return FIRST_ZONE_CODE
// 	case c.X <= SECOND_ZONE_EDGES[0].X && c.Y >= SECOND_ZONE_EDGES[1].Y:
// 		return SECOND_ZONE_CODE
// 	case c.X >= THIRD_ZONE_EDGES[0].X && c.Y >= THIRD_ZONE_EDGES[1].Y:
// 		return THIRD_ZONE_CODE
// 	case c.X <= FORTH_ZONE_EDGES[0].X && c.Y <= FORTH_ZONE_EDGES[1].Y:
// 		return FORTH_ZONE_CODE
// 	}
// 	return ZERO_VALUE
// }

// // func isFreeZoneToStay(e game.Element) bool {
// // 	return e == game.NONE || e == game.ICE || e == game.AI_TANK_PRIZE
// // }

// // func isFreeToShot(e game.Element) bool {
// // 	return e == game.AI_TANK_DOWN || e == game.AI_TANK_LEFT || e == game.AI_TANK_RIGHT || e == game.AI_TANK_UP || e == game.OTHER_TANK_DOWN || e == game.OTHER_TANK_LEFT || e == game.OTHER_TANK_RIGHT || e == game.OTHER_TANK_UP
// // }

// // func isEnemyAt(c game.Point, e []game.Point) bool {
// // 	for _, v := range e {
// // 		if c == v {
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }

// func isFreeAt(c game.Point, a []game.Point) bool {
// 	for _, v := range a {
// 		if c == v {
// 			return true
// 		}
// 	}
// 	return false
// }

// func updateMyCoords(c game.Point) {
// 	MY_COORDS = c
// }

// func createGraph(c game.Point, a []game.Point) *graph.Graph {

// 	g := graph.New(600)

// 	for _, v := range a {

// 		top := game.Point{X: v.X, Y: v.Y + 1}
// 		right := game.Point{X: v.X + 1, Y: v.Y}
// 		left := game.Point{X: v.X - 1, Y: v.Y}
// 		bottom := game.Point{X: v.X, Y: v.Y - 1}

// 		if isFreeAt(top, a) {
// 			g.Connect(v, top)
// 		}

// 		if isFreeAt(right, a) {
// 			g.Connect(v, right)
// 		}

// 		if isFreeAt(left, a) {
// 			g.Connect(v, left)
// 		}

// 		if isFreeAt(bottom, a) {
// 			g.Connect(v, bottom)
// 		}
// 	}

// 	return g
// }

// func getTheBestTactic(myCoords game.Point, destination game.Point, b *game.Board) string {

// 	if myCoords == EMPTY_COORDS{
// 		myCoords = MY_COORDS
// 	}

// 	a := b.GetAllPoints(game.NONE, game.TREE, game.ICE, game.PRIZE_IMMORTALITY, game.PRIZE_BREAKING_WALLS, game.PRIZE_VISIBILITY, game.PRIZE_NO_SLIDING, game.PRIZE_WALKING_ON_WATER, game.OTHER_TANK_DOWN, game.OTHER_TANK_LEFT, game.OTHER_TANK_RIGHT, game.OTHER_TANK_UP, game.AI_TANK_DOWN, game.AI_TANK_LEFT, game.AI_TANK_RIGHT, game.AI_TANK_UP, game.AI_TANK_PRIZE)
// 	a = append(a, myCoords)

// 	g := createGraph(myCoords, a)
// 	r := algorithm.New(g, myCoords)

// 	path := r.Path(destination)
// 	if len(path) <= 1 {
// 		return ZERO_TACTIC
// 	}

// 	top := game.Point{X: myCoords.X, Y: myCoords.Y + 1}
// 	right := game.Point{X: myCoords.X + 1, Y: myCoords.Y}
// 	left := game.Point{X: myCoords.X - 1, Y: myCoords.Y}
// 	bottom := game.Point{X: myCoords.X, Y: myCoords.Y - 1}
	

// 	switch path[1]{
// 	case top:
// 		updateMyCoords(top)
// 		return UP
// 	case right:
// 		updateMyCoords(right)
// 		return RIGHT
// 	case left:
// 		updateMyCoords(left)
// 		return LEFT
// 	case bottom:
// 		updateMyCoords(bottom)
// 		return DOWN
// 	default:
// 		return ZERO_TACTIC
// 	}
// }