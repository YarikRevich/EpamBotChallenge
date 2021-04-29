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

	Trap      bool
	StaticWay *direction.Direction

	Recession bool

	Way      direction.Direction
	Shoot    bool
	MoveFire bool
}

func (m *Middleware) GetBestWayMiddleware() {
	t := algorithm.GetBestTactic(
		m.Default.b.GetMe(),
		*m.MyBullet,
		utils.GetTheNearestElement(m.Default.b.GetMe(), utils.GetAvailableEnemies(m.Default.b.GetEnemies(), m.Default.b), m.Default.b),
		m.Default.b,
	)

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

func (m *Middleware) TrapMiddleware() {
	// if *m.StaticWay == direction.NONE{
	// 	*m.StaticWay = m.Way
	// }
}

func (m *Middleware) ResetTrapMiddleware() {
	*m.StaticWay = direction.NONE
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

		var counter int

		switch m.Way {
		case direction.UP, direction.DOWN:
			for counter != 2 {

				right := game.Point{X: s.X + counter, Y: s.Y}
				left := game.Point{X: s.X - counter, Y: s.Y}

				if utils.IsElementEnemy(m.Default.b.GetAt(right)) && utils.ElementIs(right, game.OTHER_TANK_LEFT, m.Default.b) ||
					utils.IsElementEnemy(m.Default.b.GetAt(right)) && utils.ElementIs(right, game.AI_TANK_LEFT, m.Default.b) {
					r = append(r, right)
				}
				if utils.IsElementEnemy(m.Default.b.GetAt(left)) && utils.ElementIs(left, game.OTHER_TANK_RIGHT, m.Default.b) ||
					utils.IsElementEnemy(m.Default.b.GetAt(left)) && utils.ElementIs(left, game.AI_TANK_RIGHT, m.Default.b) {
					r = append(r, left)
				}
				counter++
			}
		case direction.RIGHT, direction.LEFT:
			for counter != 2 {

				top := game.Point{X: s.X, Y: s.Y + counter}
				bottom := game.Point{X: s.X, Y: s.Y - counter}

				if utils.IsElementEnemy(m.Default.b.GetAt(top)) && utils.ElementIs(top, game.OTHER_TANK_DOWN, m.Default.b) ||
					utils.IsElementEnemy(m.Default.b.GetAt(top)) && utils.ElementIs(top, game.AI_TANK_DOWN, m.Default.b) {
					r = append(r, top)
				}
				if utils.IsElementEnemy(m.Default.b.GetAt(bottom)) && utils.ElementIs(bottom, game.OTHER_TANK_UP, m.Default.b) ||
					utils.IsElementEnemy(m.Default.b.GetAt(bottom)) && utils.ElementIs(bottom, game.AI_TANK_UP, m.Default.b) {
					r = append(r, bottom)
				}

				if utils.IsElementEnemy(m.Default.b.GetAt(top)) {
					r = append(r, top)
				}
				if utils.IsElementEnemy(m.Default.b.GetAt(bottom)) {
					r = append(r, bottom)
				}
				counter++
			}
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

	top := game.Point{X: a.X, Y: a.Y + 1}
	right := game.Point{X: a.X + 1, Y: a.Y}
	left := game.Point{X: a.X - 1, Y: a.Y}
	down := game.Point{X: a.X, Y: a.Y - 1}

	for i := 2; i >= 0; i-- {

		if m.MoveFire {
			break
		}

		switch {
		case (utils.ElementIs(a, game.TANK_UP, m.Default.b) && utils.ElementIs(down, game.OTHER_TANK_UP, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_UP, m.Default.b) && utils.ElementIs(down, game.AI_TANK_UP, m.Default.b)):
			m.MoveFire = true
		case (utils.ElementIs(a, game.TANK_DOWN, m.Default.b) && utils.ElementIs(top, game.OTHER_TANK_DOWN, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_DOWN, m.Default.b) && utils.ElementIs(top, game.AI_TANK_DOWN, m.Default.b)):
			m.MoveFire = true
		case (utils.ElementIs(a, game.TANK_LEFT, m.Default.b) && utils.ElementIs(right, game.OTHER_TANK_LEFT, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_LEFT, m.Default.b) && utils.ElementIs(right, game.AI_TANK_LEFT, m.Default.b)):
			m.MoveFire = true
		case (utils.ElementIs(a, game.TANK_RIGHT, m.Default.b) && utils.ElementIs(left, game.OTHER_TANK_RIGHT, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_RIGHT, m.Default.b) && utils.ElementIs(left, game.AI_TANK_RIGHT, m.Default.b)):
			m.MoveFire = true
		case (utils.ElementIs(a, game.TANK_UP, m.Default.b) && utils.ElementIs(right, game.OTHER_TANK_UP, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_UP, m.Default.b) && utils.ElementIs(right, game.AI_TANK_UP, m.Default.b)):
			m.MoveFire = true

		case (utils.ElementIs(a, game.TANK_UP, m.Default.b) && utils.ElementIs(right, game.OTHER_TANK_DOWN, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_UP, m.Default.b) && utils.ElementIs(right, game.AI_TANK_DOWN, m.Default.b)):
			m.MoveFire = true
		case (utils.ElementIs(a, game.TANK_UP, m.Default.b) && utils.ElementIs(left, game.OTHER_TANK_DOWN, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_UP, m.Default.b) && utils.ElementIs(left, game.AI_TANK_DOWN, m.Default.b)):
			m.MoveFire = true
		case (utils.ElementIs(a, game.TANK_DOWN, m.Default.b) && utils.ElementIs(right, game.OTHER_TANK_DOWN, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_DOWN, m.Default.b) && utils.ElementIs(right, game.AI_TANK_DOWN, m.Default.b)):
			m.MoveFire = true
		case (utils.ElementIs(a, game.TANK_DOWN, m.Default.b) && utils.ElementIs(left, game.OTHER_TANK_DOWN, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_DOWN, m.Default.b) && utils.ElementIs(left, game.AI_TANK_DOWN, m.Default.b)):
			m.MoveFire = true

		case (utils.ElementIs(a, game.TANK_DOWN, m.Default.b) && utils.ElementIs(right, game.OTHER_TANK_DOWN, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_DOWN, m.Default.b) && utils.ElementIs(right, game.AI_TANK_DOWN, m.Default.b)):
			m.MoveFire = true
		case (utils.ElementIs(a, game.TANK_DOWN, m.Default.b) && utils.ElementIs(left, game.OTHER_TANK_DOWN, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_DOWN, m.Default.b) && utils.ElementIs(left, game.AI_TANK_DOWN, m.Default.b)):
			m.MoveFire = true
		case (utils.ElementIs(a, game.TANK_UP, m.Default.b) && utils.ElementIs(right, game.OTHER_TANK_DOWN, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_UP, m.Default.b) && utils.ElementIs(right, game.AI_TANK_DOWN, m.Default.b)):
			m.MoveFire = true
		case (utils.ElementIs(a, game.TANK_DOWN, m.Default.b) && utils.ElementIs(left, game.OTHER_TANK_DOWN, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_DOWN, m.Default.b) && utils.ElementIs(left, game.AI_TANK_DOWN, m.Default.b)):
			m.MoveFire = true

		case (utils.ElementIs(a, game.TANK_RIGHT, m.Default.b) && utils.ElementIs(down, game.OTHER_TANK_RIGHT, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_RIGHT, m.Default.b) && utils.ElementIs(down, game.AI_TANK_RIGHT, m.Default.b)):
			m.MoveFire = true
		case (utils.ElementIs(a, game.TANK_RIGHT, m.Default.b) && utils.ElementIs(down, game.OTHER_TANK_LEFT, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_RIGHT, m.Default.b) && utils.ElementIs(down, game.AI_TANK_LEFT, m.Default.b)):
			m.MoveFire = true
		case (utils.ElementIs(a, game.TANK_LEFT, m.Default.b) && utils.ElementIs(down, game.OTHER_TANK_RIGHT, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_LEFT, m.Default.b) && utils.ElementIs(down, game.AI_TANK_RIGHT, m.Default.b)):
			m.MoveFire = true
		case (utils.ElementIs(a, game.TANK_LEFT, m.Default.b) && utils.ElementIs(down, game.OTHER_TANK_LEFT, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_LEFT, m.Default.b) && utils.ElementIs(down, game.AI_TANK_LEFT, m.Default.b)):
			m.MoveFire = true

		case (utils.ElementIs(a, game.TANK_RIGHT, m.Default.b) && utils.ElementIs(top, game.OTHER_TANK_RIGHT, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_RIGHT, m.Default.b) && utils.ElementIs(top, game.AI_TANK_RIGHT, m.Default.b)):
			m.MoveFire = true
		case (utils.ElementIs(a, game.TANK_RIGHT, m.Default.b) && utils.ElementIs(top, game.OTHER_TANK_LEFT, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_RIGHT, m.Default.b) && utils.ElementIs(top, game.AI_TANK_LEFT, m.Default.b)):
			m.MoveFire = true
		case (utils.ElementIs(a, game.TANK_LEFT, m.Default.b) && utils.ElementIs(top, game.OTHER_TANK_RIGHT, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_LEFT, m.Default.b) && utils.ElementIs(top, game.AI_TANK_RIGHT, m.Default.b)):
			m.MoveFire = true
		case (utils.ElementIs(a, game.TANK_LEFT, m.Default.b) && utils.ElementIs(top, game.OTHER_TANK_LEFT, m.Default.b)) ||
			(utils.ElementIs(a, game.TANK_LEFT, m.Default.b) && utils.ElementIs(top, game.AI_TANK_LEFT, m.Default.b)):
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

func (m *Middleware) RecessionMiddleware() {
	if *m.KD != 4 {

		myCoords := m.Default.b.GetMe()

		switch m.Way {
		case direction.UP:
			if utils.IsElementEnemy(m.Default.b.GetAt(game.Point{X: myCoords.X, Y: myCoords.Y + 1})) {
				m.Recession = true
			}
		case direction.RIGHT:
			if utils.IsElementEnemy(m.Default.b.GetAt(game.Point{X: myCoords.X + 1, Y: myCoords.Y})) {
				m.Recession = true
			}
		case direction.LEFT:
			if utils.IsElementEnemy(m.Default.b.GetAt(game.Point{X: myCoords.X - 1, Y: myCoords.Y})) {
				m.Recession = true
			}
		case direction.DOWN:
			if utils.IsElementEnemy(m.Default.b.GetAt(game.Point{X: myCoords.X, Y: myCoords.Y - 1})) {
				m.Recession = true
			}
		}
		if m.Recession {
			switch {
			case utils.IsWithin(game.Point{X: myCoords.X, Y: myCoords.Y + 1}, utils.GetAvailableZoneToGo(m.Default.b)):
				m.Way = direction.UP
			case utils.IsWithin(game.Point{X: myCoords.X + 1, Y: myCoords.Y}, utils.GetAvailableZoneToGo(m.Default.b)):
				m.Way = direction.RIGHT
			case utils.IsWithin(game.Point{X: myCoords.X - 1, Y: myCoords.Y}, utils.GetAvailableZoneToGo(m.Default.b)):
				m.Way = direction.LEFT
			case utils.IsWithin(game.Point{X: myCoords.X, Y: myCoords.Y - 1}, utils.GetAvailableZoneToGo(m.Default.b)):
				m.Way = direction.DOWN
			}
		}
	}
}

func Run(b *game.Board, KD *int, MyBullet *game.Point, StaticWay *direction.Direction) *Middleware {
	m := new(Middleware)
	m.Default.b = b
	m.KD = KD
	m.MyBullet = MyBullet
	m.StaticWay = StaticWay

	m.RegKDMiddleware()

	m.GetBestWayMiddleware()

	m.UpdatingProcessMiddleware()

	m.UpdateBulletMiddleware()

	m.RecessionMiddleware()

	if m.Recession {
		return m
	}

	// switch {
	// case m.Trap:
	// 	m.ResetKDMiddleware()

	// 	m.TrapMiddleware()
	// 	m.AllowToShootForcibly()
	// default:
	m.ResetTrapMiddleware()

	m.CanShootMiddleware()

	if m.Shoot {
		m.ResetKDMiddleware() // Resets KD to its beginning position ...

		m.ShouldMoveFireOrFireMoveMiddleware()

		m.RegBulletMiddleware() // Regs outcoming bullet ...

	}
	// }
	return m
}
