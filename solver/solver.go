/*-
 * #%L
 * Codenjoy - it's a dojo-like platform from developers to developers.
 * %%
 * Copyright (C) 2018 - 2021 Codenjoy
 * %%
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public
 * License along with this program.  If not, see
 * <http://www.gnu.org/licenses/gpl-3.0.html>.
 * #L%
 */
package solver

import (

	"battlecity_test/action"
	"battlecity_test/game"
	"battlecity_test/solver/middlewares"
)

var (
	PreviousStep game.Point
)

type Solver struct {
	KD *int
	MyBullet *game.Point
}

func New() Solver {
	var k int = 4
	return Solver{KD: &k, MyBullet: new(game.Point)}
}

func (s *Solver) GetNextAction(b *game.Board) action.Action {
	//todo: your code here

	m := middlewares.Run(b, s.KD, s.MyBullet)

	if m.Way.IsValid() {
		if m.Shoot{
			if m.MoveFire{
				return action.MoveFire(m.Way)
			}
			return action.FireMove(m.Way)
		}
		return action.Move(m.Way)
	}

	return action.DoNothing()
}
