package middlewares

import (
	"battlecity_test/direction"
	"battlecity_test/game"
	"battlecity_test/solver/algorithm"
	"battlecity_test/solver/utils"
	// "fmt"
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

func (m *Middleware) AvoidOuterBulletsMiddleware() {

	my := m.Default.b.GetMe()

	t := game.Point{X: my.X, Y: my.Y + 1}
	r := game.Point{X: my.X + 1, Y: my.Y}
	l := game.Point{X: my.X - 1, Y: my.Y}
	b := game.Point{X: my.X, Y: my.Y - 1}

	top := game.Point{X: my.X, Y: my.Y + 2}
	right := game.Point{X: my.X + 2, Y: my.Y}
	left := game.Point{X: my.X - 2, Y: my.Y}
	bottom := game.Point{X: my.X, Y: my.Y - 2}

	topItem := m.Default.b.GetAt(top)
	//Zone to evaluate the top way to go

	topTopItem := m.Default.b.GetAt(game.Point{X: t.X, Y: t.Y + 1})
	topRightItem := m.Default.b.GetAt(game.Point{X: t.X + 1, Y: t.Y})
	topLeftItem := m.Default.b.GetAt(game.Point{X: t.X - 1, Y: t.Y})

	topTopTopItem := m.Default.b.GetAt(game.Point{X: t.X, Y: t.Y + 2})
	topTopRightItem := m.Default.b.GetAt(game.Point{X: t.X + 2, Y: t.Y})
	topTopLeftItem := m.Default.b.GetAt(game.Point{X: t.X - 2, Y: t.Y})

	rightItem := m.Default.b.GetAt(right)
	//Zone to evaluate the right way to go

	rightTopItem := m.Default.b.GetAt(game.Point{X: r.X, Y: r.Y + 1})
	rightRightItem := m.Default.b.GetAt(game.Point{X: r.X + 1, Y: r.Y})
	rightBottomItem := m.Default.b.GetAt(game.Point{X: r.X, Y: r.Y - 1})

	rightRightTopItem := m.Default.b.GetAt(game.Point{X: r.X, Y: r.Y + 2})
	rightRightRightItem := m.Default.b.GetAt(game.Point{X: r.X + 2, Y: r.Y})
	rightRightBottomItem := m.Default.b.GetAt(game.Point{X: r.X, Y: r.Y - 2})

	leftItem := m.Default.b.GetAt(left)
	//Zone to evaluate the left way to go

	leftTopItem := m.Default.b.GetAt(game.Point{X: l.X, Y: l.Y + 1})
	leftLeftItem := m.Default.b.GetAt(game.Point{X: l.X - 1, Y: l.Y})
	leftBottomItem := m.Default.b.GetAt(game.Point{X: l.X, Y: l.Y - 1})

	leftLeftTopItem := m.Default.b.GetAt(game.Point{X: l.X, Y: l.Y + 2})
	leftLeftLeftItem := m.Default.b.GetAt(game.Point{X: l.X - 2, Y: l.Y})
	leftLeftBottomItem := m.Default.b.GetAt(game.Point{X: l.X, Y: l.Y - 2})

	bottomItem := m.Default.b.GetAt(bottom)
	//Zone to evaluate the left way to go

	bottomBottomItem := m.Default.b.GetAt(game.Point{X: b.X, Y: b.Y - 1})
	bottomRightItem := m.Default.b.GetAt(game.Point{X: b.X + 1, Y: b.Y})
	bottomLeftItem := m.Default.b.GetAt(game.Point{X: b.X - 1, Y: b.Y})

	bottomBottomBottomItem := m.Default.b.GetAt(game.Point{X: b.X, Y: b.Y - 2})
	bottomBottomRightItem := m.Default.b.GetAt(game.Point{X: b.X + 2, Y: b.Y})
	bottomBottomLeftItem := m.Default.b.GetAt(game.Point{X: b.X - 2, Y: b.Y})

	n := utils.GetNoneElements(m.Default.b)

	switch m.Way {
	case direction.UP:
		if topItem != game.BULLET {
			return
		}

		switch {
		case utils.IsWithinElements(rightItem, n) &&
			utils.IsWithinElements(rightTopItem, n) &&
			utils.IsWithinElements(rightRightItem, n) &&
			utils.IsWithinElements(rightBottomItem, n) &&
			utils.IsWithinElements(rightRightTopItem, n) &&
			utils.IsWithinElements(rightRightRightItem, n) &&
			utils.IsWithinElements(rightRightBottomItem, n) && right != *m.MyBullet:
			m.Way = direction.RIGHT
		case utils.IsWithinElements(leftItem, n) &&
			utils.IsWithinElements(leftTopItem, n) &&
			utils.IsWithinElements(leftLeftItem, n) &&
			utils.IsWithinElements(leftBottomItem, n) &&
			utils.IsWithinElements(leftLeftTopItem, n) &&
			utils.IsWithinElements(leftLeftLeftItem, n) &&
			utils.IsWithinElements(leftLeftBottomItem, n) && left != *m.MyBullet:
			m.Way = direction.LEFT
		case utils.IsWithinElements(bottomItem, n) &&
			utils.IsWithinElements(bottomBottomItem, n) &&
			utils.IsWithinElements(bottomLeftItem, n) &&
			utils.IsWithinElements(bottomRightItem, n) &&
			utils.IsWithinElements(bottomBottomBottomItem, n) &&
			utils.IsWithinElements(bottomBottomLeftItem, n) &&
			utils.IsWithinElements(bottomBottomRightItem, n) && bottom != *m.MyBullet:
			m.Way = direction.DOWN
		}
	case direction.RIGHT:
		if rightItem != game.BULLET {
			return
		}

		switch {
		case utils.IsWithinElements(topItem, n) &&
			utils.IsWithinElements(topLeftItem, n) &&
			utils.IsWithinElements(topRightItem, n) &&
			utils.IsWithinElements(topTopItem, n) &&
			utils.IsWithinElements(topTopLeftItem, n) &&
			utils.IsWithinElements(topTopRightItem, n) &&
			utils.IsWithinElements(topTopTopItem, n) && top != *m.MyBullet:
			m.Way = direction.UP
		case utils.IsWithinElements(leftItem, n) &&
			utils.IsWithinElements(leftBottomItem, n) &&
			utils.IsWithinElements(leftLeftItem, n) &&
			utils.IsWithinElements(leftTopItem, n) &&
			utils.IsWithinElements(leftLeftTopItem, n) &&
			utils.IsWithinElements(leftLeftLeftItem, n) &&
			utils.IsWithinElements(leftLeftBottomItem, n) && left != *m.MyBullet:
			m.Way = direction.LEFT
		case utils.IsWithinElements(bottomItem, n) &&
			utils.IsWithinElements(bottomBottomItem, n) &&
			utils.IsWithinElements(bottomLeftItem, n) &&
			utils.IsWithinElements(bottomRightItem, n) &&
			utils.IsWithinElements(bottomBottomBottomItem, n) &&
			utils.IsWithinElements(bottomBottomLeftItem, n) &&
			utils.IsWithinElements(bottomBottomRightItem, n) && bottom != *m.MyBullet:
			m.Way = direction.DOWN
		}
	case direction.LEFT:
		if leftItem != game.BULLET {
			return
		}

		switch {
		case utils.IsWithinElements(topItem, n) &&
			utils.IsWithinElements(topLeftItem, n) &&
			utils.IsWithinElements(topRightItem, n) &&
			utils.IsWithinElements(topTopItem, n) &&
			utils.IsWithinElements(topTopLeftItem, n) &&
			utils.IsWithinElements(topTopRightItem, n) &&
			utils.IsWithinElements(topTopTopItem, n) && top != *m.MyBullet:
			m.Way = direction.UP
		case utils.IsWithinElements(rightItem, n) &&
			utils.IsWithinElements(rightTopItem, n) &&
			utils.IsWithinElements(rightRightItem, n) &&
			utils.IsWithinElements(rightBottomItem, n) &&
			utils.IsWithinElements(rightRightTopItem, n) &&
			utils.IsWithinElements(rightRightRightItem, n) &&
			utils.IsWithinElements(rightRightBottomItem, n) && right != *m.MyBullet:
			m.Way = direction.RIGHT
		case utils.IsWithinElements(bottomItem, n) &&
			utils.IsWithinElements(bottomBottomItem, n) &&
			utils.IsWithinElements(bottomLeftItem, n) &&
			utils.IsWithinElements(bottomRightItem, n) &&
			utils.IsWithinElements(bottomBottomBottomItem, n) &&
			utils.IsWithinElements(bottomBottomLeftItem, n) &&
			utils.IsWithinElements(bottomBottomRightItem, n) && bottom != *m.MyBullet:
			m.Way = direction.DOWN
		}
	case direction.DOWN:
		if bottomItem != game.BULLET {
			return
		}

		switch {
		case utils.IsWithinElements(topItem, n) &&
			utils.IsWithinElements(topLeftItem, n) &&
			utils.IsWithinElements(topTopItem, n) &&
			utils.IsWithinElements(topRightItem, n) &&
			utils.IsWithinElements(topTopLeftItem, n) &&
			utils.IsWithinElements(topTopRightItem, n) &&
			utils.IsWithinElements(topTopTopItem, n) && top != *m.MyBullet:
			m.Way = direction.UP
		case utils.IsWithinElements(rightItem, n) &&
			utils.IsWithinElements(rightTopItem, n) &&
			utils.IsWithinElements(rightRightItem, n) &&
			utils.IsWithinElements(rightBottomItem, n) &&
			utils.IsWithinElements(rightRightTopItem, n) &&
			utils.IsWithinElements(rightRightRightItem, n) &&
			utils.IsWithinElements(rightRightBottomItem, n) && right != *m.MyBullet:
			m.Way = direction.RIGHT
		case utils.IsWithinElements(leftItem, n) &&
			utils.IsWithinElements(leftBottomItem, n) &&
			utils.IsWithinElements(leftLeftItem, n) &&
			utils.IsWithinElements(leftTopItem, n) &&
			utils.IsWithinElements(leftLeftTopItem, n) &&
			utils.IsWithinElements(leftLeftLeftItem, n) &&
			utils.IsWithinElements(leftLeftBottomItem, n) && left != *m.MyBullet:
			m.Way = direction.LEFT
		}
	}
}

func (m *Middleware) ResetTrapMiddleware() {
	*m.StaticWay = direction.NONE
}

func (m *Middleware) UpdatingProcessMiddleware() {
	if utils.IsUpdatingProcess(utils.GetAIEnemies(m.Default.b)) {
		m.Updation = true
	}
}

func (m *Middleware) CanShootMiddleware() {

	if m.Way == direction.NONE {
		return
	}

	s := m.Default.b.GetMe()
	r := []game.Point{}
	var rowNum int

	for {

		if rowNum == 6 {
			break
		}

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

				counter++
			}
		}

		r = append(r, s)

		rowNum++
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

	for i := 2; i >= 0; i-- {

		top := game.Point{X: a.X, Y: a.Y + 1}
		right := game.Point{X: a.X + 1, Y: a.Y}
		left := game.Point{X: a.X - 1, Y: a.Y}
		bottom := game.Point{X: a.X, Y: a.Y - 1}

		if m.MoveFire {
			break
		}

		for _, tankOr := range []game.Element{game.TANK_UP, game.TANK_RIGHT, game.TANK_LEFT, game.TANK_DOWN} {
			for _, enemyPos := range []game.Point{top, right, left, bottom} {
				for _, enemyOr := range []game.Element{game.AI_TANK_UP, game.AI_TANK_RIGHT, game.AI_TANK_LEFT, game.AI_TANK_DOWN, game.OTHER_TANK_UP, game.OTHER_TANK_RIGHT, game.OTHER_TANK_LEFT, game.OTHER_TANK_DOWN} {
					if utils.ElementIs(a, tankOr, m.Default.b) && utils.ElementIs(enemyPos, enemyOr, m.Default.b) {
						m.MoveFire = true
					}
				}
			}
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

			top := game.Point{X: myCoords.X, Y: myCoords.Y + 1}
			right := game.Point{X: myCoords.X + 1, Y: myCoords.Y}
			left := game.Point{X: myCoords.X - 1, Y: myCoords.Y}
			bottom := game.Point{X: myCoords.X, Y: myCoords.Y - 1}

			topTop := game.Point{X: top.X, Y: top.Y + 1}
			topRight := game.Point{X: top.X + 1, Y: top.Y}
			topLeft := game.Point{X: top.X - 1, Y: top.Y}

			rightTop := game.Point{X: right.X, Y: right.Y + 1}
			rightRight := game.Point{X: right.X + 1, Y: right.Y}
			rightBottom := game.Point{X: right.X, Y: right.Y - 1}

			leftTop := game.Point{X: left.X, Y: left.Y + 1}
			leftLeft := game.Point{X: left.X - 1, Y: left.Y}
			leftBottom := game.Point{X: left.X, Y: left.Y - 1}

			bottomBottom := game.Point{X: top.X, Y: top.Y - 1}
			bottomRight := game.Point{X: top.X + 1, Y: top.Y}
			bottomLeft := game.Point{X: top.X - 1, Y: top.Y}

			switch {
			case utils.IsWithin(top, utils.GetAvailableZoneToGo(m.Default.b)) &&
				utils.IsWithin(topTop, utils.GetAvailableZoneToGo(m.Default.b)) &&
				utils.IsWithin(topRight, utils.GetAvailableZoneToGo(m.Default.b)) &&
				utils.IsWithin(topLeft, utils.GetAvailableZoneToGo(m.Default.b)):
				m.Way = direction.UP
			case utils.IsWithin(right, utils.GetAvailableZoneToGo(m.Default.b)) &&
				utils.IsWithin(rightTop, utils.GetAvailableZoneToGo(m.Default.b)) &&
				utils.IsWithin(rightRight, utils.GetAvailableZoneToGo(m.Default.b)) &&
				utils.IsWithin(rightBottom, utils.GetAvailableZoneToGo(m.Default.b)):
				m.Way = direction.RIGHT
			case utils.IsWithin(left, utils.GetAvailableZoneToGo(m.Default.b)) &&
				utils.IsWithin(leftTop, utils.GetAvailableZoneToGo(m.Default.b)) &&
				utils.IsWithin(leftLeft, utils.GetAvailableZoneToGo(m.Default.b)) &&
				utils.IsWithin(leftBottom, utils.GetAvailableZoneToGo(m.Default.b)):
				m.Way = direction.LEFT
			case utils.IsWithin(bottom, utils.GetAvailableZoneToGo(m.Default.b)) &&
				utils.IsWithin(bottomBottom, utils.GetAvailableZoneToGo(m.Default.b)) &&
				utils.IsWithin(bottomRight, utils.GetAvailableZoneToGo(m.Default.b)) &&
				utils.IsWithin(bottomLeft, utils.GetAvailableZoneToGo(m.Default.b)):
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

	m.AvoidOuterBulletsMiddleware()

	m.UpdatingProcessMiddleware()

	m.UpdateBulletMiddleware()

	m.RecessionMiddleware()

	if m.Recession {
		return m
	}

	switch {
	case m.Updation:
		m.ResetKDMiddleware()

	default:
		m.ResetTrapMiddleware()

		m.CanShootMiddleware()

		if m.Shoot {
			m.ResetKDMiddleware() // Resets KD to its beginning position ...

			m.ShouldMoveFireOrFireMoveMiddleware()

			m.RegBulletMiddleware() // Regs outcoming bullet ...

		}
	}
	return m
}
