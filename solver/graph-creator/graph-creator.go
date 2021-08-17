package graphcreator

import (
	"battlecity_test/game"
	"battlecity_test/solver/utils"
	"battlecity_test/third-party/graph"
	"github.com/YarikRevich/GSL/pkg/array"
)

type GraphCreator struct {
	graph *graph.Graph
	board *game.Board
}

func (gc *GraphCreator) CreateGraph() error {
	g := graph.New(gc.board.BoardSize())

	availablePoints := gc.board.GetAllPoints(utils.GetAvailableElements(gc.board)...)
	for _, v := range availablePoints {

		top := game.Point{X: v.X, Y: v.Y + 1}
		right := game.Point{X: v.X + 1, Y: v.Y}
		left := game.Point{X: v.X - 1, Y: v.Y}
		bottom := game.Point{X: v.X, Y: v.Y - 1}

		ok, err := array.InArray(availablePoints, top)
		if err != nil {
			return err
		}
		if ok {
			g.Connect(v, top)
		}

		ok, err = array.InArray(availablePoints, right)
		if err != nil {
			return err
		}
		if ok {
			g.Connect(v, right)
		}

		ok, err = array.InArray(availablePoints, left)
		if err != nil {
			return err
		}
		if ok {
			g.Connect(v, left)
		}

		ok, err = array.InArray(availablePoints, bottom)
		if err != nil {
			return err
		}
		if ok {
			g.Connect(v, bottom)
		}
	}

	gc.graph = g
	return nil
}

func (gc *GraphCreator) GetGraph() *graph.Graph {
	return gc.graph
}

func New(board *game.Board) *GraphCreator {
	return &GraphCreator{
		board: board,
	}
}
