package standart

import (
	"battlecity_test/game"
	"battlecity_test/solver/coords"
	"battlecity_test/solver/suicide"
	"battlecity_test/solver/target"
	"battlecity_test/solver/tickets"
	"battlecity_test/solver/wayestimator/bottom"
	"battlecity_test/solver/wayestimator/left"
	"battlecity_test/solver/wayestimator/right"
	"battlecity_test/solver/wayestimator/top"
	"battlecity_test/third-party/bfs"
	"battlecity_test/third-party/graph"
)

type StandartWayEstimator struct {
	g *graph.Graph
	b *game.Board

	coords *coords.Coords

	topWayEstimator    *top.TopWayEstimator
	bottomWayEstimator *bottom.BottomWayEstimator
	rightWayEstimator  *right.RightWayEstimator
	leftWayEstimator   *left.LeftWayEstimator
}

func (swe *StandartWayEstimator) Estimate() {
	b := bfs.New(swe.g, swe.b.GetMe())

	t := target.New(swe.b).GetTarget()
	path := b.Path(t)

	if len(path) > 0 {
		var defaultWay game.Point
		if suicide.IsBetterToSuicide(swe.b.GetMe(), t) {
			// 	return suicide.GetMoveToSoicide()
		}
		defaultWay = path[0]

		switch defaultWay {
		case swe.coords.Top:
			swe.topWayEstimator.GetTicket().EstimateMan(tickets.DEFAULT)
		case swe.coords.Down:
			swe.bottomWayEstimator.GetTicket().EstimateMan(tickets.DEFAULT)
		case swe.coords.Right:
			swe.rightWayEstimator.GetTicket().EstimateMan(tickets.DEFAULT)
		default:
			swe.leftWayEstimator.GetTicket().EstimateMan(tickets.DEFAULT)
		}
	}
}

func New(b *game.Board,
	topWayEstimator *top.TopWayEstimator,
	bottomWayEstimator *bottom.BottomWayEstimator,
	rightWayEstimator *right.RightWayEstimator,
	leftWayEstimator *left.LeftWayEstimator) *StandartWayEstimator {
	return &StandartWayEstimator{
		b:                  b,
		coords:             coords.New(b),
		topWayEstimator:    topWayEstimator,
		bottomWayEstimator: bottomWayEstimator,
		rightWayEstimator:  rightWayEstimator,
		leftWayEstimator:   leftWayEstimator,
	}
}
