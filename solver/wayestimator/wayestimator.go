package wayestimator

import (
	"battlecity_test/action"
	"battlecity_test/direction"
	"battlecity_test/game"

	"battlecity_test/solver/tickets"
	"battlecity_test/third-party/graph"

	"github.com/adam-lavrik/go-imath/ix"

	"battlecity_test/solver/wayestimator/bottom"
	"battlecity_test/solver/wayestimator/left"
	"battlecity_test/solver/wayestimator/right"
	"battlecity_test/solver/wayestimator/standart"
	"battlecity_test/solver/wayestimator/top"
)

type WayEstimatorManager struct {
	g *graph.Graph
	b *game.Board

	topWayEstimator    *top.TopWayEstimator
	bottomWayEstimator *bottom.BottomWayEstimator
	rightWayEstimator  *right.RightWayEstimator
	leftWayEstimator   *left.LeftWayEstimator
}

func (e *WayEstimatorManager) Estimate() action.Action {

	tw := e.topWayEstimator.Estimate()
	bw := e.bottomWayEstimator.Estimate()
	rw := e.rightWayEstimator.Estimate()
	lw := e.leftWayEstimator.Estimate()

	standart.New(e.b,
		e.topWayEstimator,
		e.bottomWayEstimator,
		e.rightWayEstimator,
		e.leftWayEstimator).Estimate()

	switch ix.MaxSlice([]int{
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

func New(g *graph.Graph, b *game.Board) *WayEstimatorManager {
	return &WayEstimatorManager{
		g:                  g,
		b:                  b,
		topWayEstimator:    top.New(tickets.New(b)),
		bottomWayEstimator: bottom.New(tickets.New(b)),
		rightWayEstimator:  right.New(tickets.New(b)),
		leftWayEstimator:   left.New(tickets.New(b)),
	}
}
