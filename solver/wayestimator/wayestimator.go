package wayestimator

import (
	"battlecity_test/action"
	"battlecity_test/direction"
	"battlecity_test/game"
	"battlecity_test/solver/target"

	// "battlecity_test/solver/tickets"
	"battlecity_test/third-party/bfs"
	"battlecity_test/third-party/graph"

	"github.com/adam-lavrik/go-imath/ix"
)

type WayEstimator struct {
	g *graph.Graph
	b *game.Board

	d action.Action //Default way
	bestw action.Action //Best way
}

func (e *WayEstimator) ProcessGraph() {
	e.createDefaultWay()

	tw := e.estimateTopWay()
	bw := e.estimateBottomWay()
	rw := e.estimateRightWay()
	lw := e.estimateLeftWay()

	mv := ix.MinSlice([]int{
		tw, bw, rw, lw})
	
	if tw == mv{
		e.bestw = action.Move(direction.UP)
	}

	if bw == mv{
		e.bestw = action.Move(direction.DOWN)
	}

	if rw == mv{
		e.bestw = action.Move(direction.RIGHT)
	}

	if lw == mv{
		e.bestw = action.Move(direction.LEFT)
	}
}

func (e *WayEstimator) getDefaultWayAction(path []game.Point)action.Action{
	if len(path) > 0{
		me := e.b.GetMe()
		firstPoint := path[0]

		if firstPoint.X == me.X + 1 && firstPoint.Y == me.Y{
			return action.MoveFire(direction.LEFT)
		}

		if firstPoint.X == me.X - 1 && firstPoint.Y == me.Y{
			return action.MoveFire(direction.RIGHT)
		}

		if firstPoint.X == me.X && firstPoint.Y + 1 == me.Y{
			return action.MoveFire(direction.UP)
		}

		if firstPoint.X == me.X && firstPoint.Y - 1  == me.Y{
			return action.MoveFire(direction.DOWN)
		}
	}
	return action.DoNothing()
}

func (e *WayEstimator) createDefaultWay(){
	b := bfs.New(e.g, e.b.GetMe())

 	e.d = e.getDefaultWayAction(b.Path(
		target.New(e.b).GetTarget()))
}

func (e *WayEstimator) estimateTopWay() int {
	return 0
}

func (e *WayEstimator) estimateRightWay() int {
	return 0
}

func (e *WayEstimator) estimateLeftWay() int {
	return 0
}

func (e *WayEstimator) estimateBottomWay() int {
	return 0
}

func (e *WayEstimator) GetWay() action.Action {
	return e.bestw
}

func New(g *graph.Graph, b *game.Board) *WayEstimator {
	return &WayEstimator{
		g: g,
		b: b,
	}
}
