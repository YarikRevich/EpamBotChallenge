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
package main

import (
	"battlecity_test/game"
	"battlecity_test/solver"
	"log"
)

func main() {

	//URL FOR PROD ;) https://epam-botchallenge.com/codenjoy-contest/board/player/qf4t9nh7yhww7507kyq6?code=260516936731504358
	
	browserURL := "https://practice.epam-botchallenge.com/codenjoy-contest/board/player/n8x403qatme2wqwyeajg?code=6999680204656435641&game=battlecity"

	game, c := game.StartGame(browserURL)
	b := game.GetBoard()
	s := solver.New()

	for {
		select {
		case <-c.Done:
			log.Fatal("It's done")
		case <-c.Read:
			// Set next move
			game.SetNextAction(s.GetNextAction(b))
			c.Write <- struct{}{}
		}
	}
}
