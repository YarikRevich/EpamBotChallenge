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

	KD *int

	MyBullet *game.Point

	Updation bool

	Trap bool

	Way      direction.Direction
	Shoot    bool
	MoveFire bool
}

func (m *Middleware) GetBestWayMiddleware() {
	t, trap := algorithm.GetBestTactic(
		m.Default.b.GetMe(),
		*m.MyBullet,
		utils.GetTheNearestElement(m.Default.b.GetMe(), m.Default.b.GetEnemies()),
		m.Default.b,
	)
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
	if utils.IsUpdatingProcess(utils.GetAIEnemies(m.Default.b)) {
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

	for _, a := range r {
		if utils.IsWithin(a, m.Default.b.GetBarriers()) {
			break
		}
		if utils.IsWithin(a, m.Default.b.GetEnemies()) {

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

func (m *Middleware) StopKDTickerMiddleware() {

	*m.KD = 4
}

func (m *Middleware) ResetKDMiddleware() {
	if *m.KD == 4 {
		*m.KD = 0
	}
}

func (m *Middleware) RegKDMiddleware() {
	if *m.KD != 4 {
		*m.KD++
	}
}

func (m *Middleware) RegBulletMiddleware() {
	if *m.KD == 0 {

		myCoords := m.Default.b.GetMe()

		switch m.Way {
		case direction.UP:
			*m.MyBullet = game.Point{X: myCoords.X, Y: myCoords.Y + 1}
		case direction.RIGHT:
			*m.MyBullet = game.Point{X: myCoords.X + 1, Y: myCoords.Y}
		case direction.LEFT:
			*m.MyBullet = game.Point{X: myCoords.X - 1, Y: myCoords.Y}
		case direction.DOWN:
			*m.MyBullet = game.Point{X: myCoords.X, Y: myCoords.Y - 1}
		}
	}
}

func (m *Middleware) UpdateBulletMiddleware() {
	if *m.MyBullet != algorithm.EMPTY_COORDS {
		if n, ok := utils.IsBulletAlive(*m.MyBullet, m.Default.b.GetBullets()); ok {
			*m.MyBullet = n
		} else {
			*m.MyBullet = algorithm.EMPTY_COORDS
		}
	}
}

func Run(b *game.Board, KD *int, MyBullet *game.Point) *Middleware {
	m := new(Middleware)
	m.Default.b = b
	m.KD = KD
	m.MyBullet = MyBullet

	m.GetBestWayMiddleware()
	m.UpdatingProcessMiddleware()

	m.UpdateBulletMiddleware()

	switch {
	case m.Trap:
		m.ResetKDMiddleware()

		m.AllowToShootForcibly()
	case m.Updation:

		m.StopKDTickerMiddleware()
	case !m.Updation:

		m.CanShootMiddleware()

		if m.Shoot {
			m.ResetKDMiddleware() // Resets KD to its beginning position ...

			m.ShouldMoveFireOrFireMoveMiddleware()

			m.RegBulletMiddleware() // Regs outcoming bullet ...

		}
	}
	return m
}
