package middlewares

import (

	"battlecity_test/direction"
	"battlecity_test/game"
	"battlecity_test/solver/algorithm"
	"battlecity_test/solver/utils"
)

type Default struct {
	b *game.Board
}

type Middleware struct {
	Default //field for servicing elements ...

	Updation bool

	Trap bool

	Way      direction.Direction
	Shoot    bool
	MoveFire bool
}

func (m *Middleware) GetBestWayMiddleware() {
	t, trap := algorithm.GetBestTactic(m.Default.b.GetMe(), utils.GetTheNearestElement(m.Default.b.GetMe(), m.Default.b.GetEnemies()), m.Default.b)
	m.Trap = trap
	switch t {
	case algorithm.UP:
		m.Way = direction.UP
	case algorithm.RIGHT:
		m.Way = direction.RIGHT
	case algorithm.LEFT:
		m.Way = direction.LEFT
	case algorithm.DOWN:
		m.Way = direction.DOWN
	default:
		m.Way = direction.NONE
	}
}

func (m *Middleware) UpdatingProcessMiddleware() {
	if utils.IsUpdatingProcess(m.Default.b.GetBullets()) {
		m.Updation = true
		m.Way = direction.NONE
	}
}

func (m *Middleware) CanShootMiddleware() {

	if m.Way == direction.NONE {
		return
	}

	s := m.Default.b.GetMe()
	r := []game.Point{}

	for {

		if utils.IsBattleWall(s, m.Default.b.GetAllPoints(game.BATTLE_WALL)) {
			break
		}

		switch m.Way {
		case direction.UP:
			s = game.Point{X: s.X, Y: s.Y + 1}
		case direction.RIGHT:
			s = game.Point{X: s.X + 1, Y: s.Y}
		case direction.LEFT:
			s = game.Point{X: s.X - 1, Y: s.Y}
		case direction.DOWN:
			s = game.Point{X: s.X, Y: s.Y - 1}
		}

		r = append(r, s)
	}

	n := []game.Point{}
	n = append(n, m.Default.b.GetEnemies()...)
	n = append(n, m.Default.b.GetAllPoints(game.RIVER, game.ICE, game.TREE, game.PRIZE, game.PRIZE_BREAKING_WALLS, game.PRIZE_IMMORTALITY, game.PRIZE_NO_SLIDING, game.PRIZE_VISIBILITY, game.PRIZE_WALKING_ON_WATER)...)

	for _, a := range r {
		if utils.IsWithin(a, m.Default.b.GetBarriers()) {
			break
		}
		if utils.IsWithin(a, n) {
			m.Shoot = true
		}
	}
}

func (m *Middleware) ShouldMoveFireOrFireMoveMiddleware() {
	a := m.Default.b.GetMe()
	f := []game.Element{
		game.AI_TANK_DOWN,
		game.AI_TANK_LEFT,
		game.AI_TANK_PRIZE,
		game.AI_TANK_RIGHT,
		game.AI_TANK_UP,
		game.OTHER_TANK_DOWN,
		game.OTHER_TANK_LEFT,
		game.OTHER_TANK_RIGHT,
		game.OTHER_TANK_UP,
		game.PRIZE,
		game.PRIZE_BREAKING_WALLS,
		game.PRIZE_IMMORTALITY,
		game.PRIZE_NO_SLIDING,
		game.PRIZE_VISIBILITY,
		game.PRIZE_WALKING_ON_WATER,
	}

	for _, v := range f {
		if m.Default.b.IsNear(a, v) {
			m.MoveFire = true
		}
	}
}

func (m *Middleware) AllowToShootForcibly() {
	m.Shoot = true
}

func Run(b *game.Board) *Middleware {
	m := new(Middleware)
	m.Default.b = b

	m.GetBestWayMiddleware()
	m.UpdatingProcessMiddleware()

	switch {
	case m.Trap:
		m.AllowToShootForcibly()
	case !m.Updation:
		m.CanShootMiddleware()
		m.ShouldMoveFireOrFireMoveMiddleware()
	}
	return m
}
