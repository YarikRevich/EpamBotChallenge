/*-
 * #%L
 * Codenjoy - it's a dojo-like platform from developers to developers.
 * %%
 * Copyright (C) 2018 - 2020 Codenjoy
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

package direction

type Direction int

const (
	UP Direction = iota
	DOWN
	RIGHT
	LEFT
	NONE
)

func (d Direction) String() string {
	return [...]string{"UP", "DOWN", "RIGHT", "LEFT", "NONE"}[d]
}

func (d Direction) IsValid() bool {
	switch d {
	case UP, DOWN, RIGHT, LEFT:
		return true
	default:
		return false
	}
}
