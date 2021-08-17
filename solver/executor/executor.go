package executor

import (
	"battlecity_test/action"
	"battlecity_test/game"
	"battlecity_test/solver/graph-creator"
	"battlecity_test/solver/wayestimator"
)

type Executor struct {
	board *game.Board
}

func (e *Executor) ProcessMap() (action.Action, error) {
	gc := graphcreator.New(e.board)

	if err := gc.CreateGraph(); err != nil{
		return action.DoNothing(), err
	}
	g := gc.GetGraph()

	we := wayestimator.New(g, e.board)
	we.ProcessGraph()

	return we.GetWay(), nil
}

func (e *Executor) GetAction() action.Action {
	if ac, err := e.ProcessMap(); err == nil {
		return ac
	}
	return action.DoNothing()
}

func New(board *game.Board) *Executor {
	return &Executor{
		board: board,
	}
}
