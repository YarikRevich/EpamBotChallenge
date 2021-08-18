package wayestimator

import (
	"battlecity_test/action"
	"battlecity_test/direction"
	"battlecity_test/game"
	"battlecity_test/solver/target"
	"battlecity_test/solver/tickets"

	// "battlecity_test/solver/tickets"
	"battlecity_test/third-party/bfs"
	"battlecity_test/third-party/graph"

	"github.com/adam-lavrik/go-imath/ix"
)

type Ways struct {
	Top   game.Point
	Down  game.Point
	Right game.Point
	Left  game.Point
}

type WayEstimator struct {
	g *graph.Graph
	b *game.Board

	ways *Ways

	topWayTicket   tickets.Ticket
	downWayTicket  tickets.Ticket
	rightWayTicket tickets.Ticket
	leftWayTicket  tickets.Ticket
}

func (e *WayEstimator) initWays() {
	me := e.b.GetMe()

	e.ways = &Ways{
		Top:   game.Point{X: me.X, Y: me.Y + 1},
		Down:  game.Point{X: me.X + 1, Y: me.Y},
		Right: game.Point{X: me.X - 1, Y: me.Y},
		Left:  game.Point{X: me.X, Y: me.Y - 1},
	}
}

func (e *WayEstimator) initDefaultWayTicket(path []game.Point) {
	if len(path) > 0 {
		e.initWays()

		defaultWay := path[0]

		defaultTicket := tickets.New(e.b)
		defaultTicket.EstimateMan(tickets.DEFAULT)

		if e.ways.Top == defaultWay {
			e.topWayTicket = *defaultTicket
		}

		if e.ways.Down == defaultWay {
			e.downWayTicket = *defaultTicket
		}

		if e.ways.Right == defaultWay {
			e.rightWayTicket = *defaultTicket
		}

		if e.ways.Left == defaultWay {
			e.leftWayTicket = *defaultTicket
		}
	}
}

func (e *WayEstimator) createDefaultWayTicket() {
	b := bfs.New(e.g, e.b.GetMe())

	e.initDefaultWayTicket(b.Path(
		target.New(e.b).GetTarget()))
}

func (e *WayEstimator) estimateTopWay() int {
	e.topWayTicket.EstimateAuto(e.ways.Top)
	return e.topWayTicket.GetAmount()
}

func (e *WayEstimator) estimateRightWay() int {
	e.rightWayTicket.EstimateAuto(e.ways.Right)
	return e.rightWayTicket.GetAmount()
}

func (e *WayEstimator) estimateLeftWay() int {
	e.leftWayTicket.EstimateAuto(e.ways.Left)
	return e.leftWayTicket.GetAmount()
}

func (e *WayEstimator) estimateBottomWay() int {
	e.downWayTicket.EstimateAuto(e.ways.Down)
	return e.downWayTicket.GetAmount()
}

func (e *WayEstimator) GetWay() action.Action {
	e.createDefaultWayTicket()

	tw := e.estimateTopWay()
	bw := e.estimateBottomWay()
	rw := e.estimateRightWay()
	lw := e.estimateLeftWay()

	switch ix.MinSlice([]int{
		tw, bw, rw, lw}) {
	case tw:
		return action.Move(direction.UP)
	case bw:
		return action.Move(direction.DOWN)
	case rw:
		return action.Move(direction.RIGHT)
	default:
		return action.Move(direction.LEFT)
	}
}

func New(g *graph.Graph, b *game.Board) *WayEstimator {
	return &WayEstimator{
		g: g,
		b: b,
	}
}
