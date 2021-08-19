package coords

import "battlecity_test/game"

type Coords struct {
	b *game.Board

	Top   game.Point
	Down  game.Point
	Right game.Point
	Left  game.Point
}

func (c *Coords) EstimateCoords() {
	me := c.b.GetMe()

	c.Top = game.Point{X: me.X, Y: me.Y + 1}
	c.Down = game.Point{X: me.X + 1, Y: me.Y}
	c.Right = game.Point{X: me.X - 1, Y: me.Y}
	c.Left = game.Point{X: me.X, Y: me.Y - 1}
}

func New(b *game.Board) *Coords {
	return &Coords{
		b: b,
	}
}
