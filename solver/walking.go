package solver

import (
	"battlecity_test/direction"
	"battlecity_test/game"
	"fmt"
)

var (
	Bulets [][]game.Point
	Steps  []game.Point

	DangerousPoints   = []game.Point{{X: 16, Y: 29}, {X: 17, Y: 29}, {X: 32, Y: 16}}
	DangerousElements = []rune{game.BULLET, game.TANK_DOWN, game.TANK_LEFT, game.TANK_RIGHT, game.TANK_UP, game.RIVER}
)

func checkEqual(a, b []game.Point) bool {
	for _, av := range a {
		for _, bv := range b {
			if av != bv {
				return false
			}
		}
	}
	return true
}

func isUpdatingProcess(e []game.Point) bool {
	Bulets = append(Bulets, e)

	if len(Bulets) == 1 {
		return false
	}

	if len(Bulets) == 2 {
		r := checkEqual(Bulets[len(Bulets)-1], Bulets[len(Bulets)-2])
		Bulets = [][]game.Point{}
		return r
	}
	return true
}

func NotBlindZone(e game.Point) bool {
	for _, v := range DangerousPoints {
		if e == v {
			return false
		}
	}
	return true
}

func NotDangerous(e game.Element) bool {
	return (e == game.NONE || e == game.ICE || e == game.OTHER_TANK_DOWN || e == game.OTHER_TANK_LEFT || e == game.OTHER_TANK_RIGHT || e == game.OTHER_TANK_UP) &&
		e != game.TREE &&
		e != game.RIVER &&
		e != game.BULLET
}

// //Checks whether hero is stuck ...
func gotStuck(c game.Point) bool {
	if len(Steps) > 4 {
		return Steps[len(Steps)-1] == Steps[len(Steps)-3] &&
			Steps[len(Steps)-2] == Steps[len(Steps)-4]

	}
	return false
}

//Main hub of hero movement decisions ;) ...
func GetReliableWayToGo(b *game.Board) direction.Direction {

	if isUpdatingProcess(b.GetBullets()) {
		Steps = []game.Point{}
		return direction.Direction(1000)
	}

	//Gets the coords of hero and elements ...

	myCoords := b.GetMe()

	top := game.Point{X: myCoords.X, Y: myCoords.Y + 1}
	bottom := game.Point{X: myCoords.X, Y: myCoords.Y - 1}
	right := game.Point{X: myCoords.X + 1, Y: myCoords.Y}
	left := game.Point{X: myCoords.X - 1, Y: myCoords.Y}

	topItem := b.GetAt(top)
	bottomItem := b.GetAt(bottom)
	rightItem := b.GetAt(right)
	leftItem := b.GetAt(left)

	fmt.Println(gotTheLastPointOfZone(myCoords), getTacticByCodeZone(getCurrentZone(myCoords)), getCurrentZone(myCoords))
	fmt.Println(myCoords)

	if CURRENT_TACTIC == ZERO_VALUE {
		setNextTactic(getTacticByCodeZone(getCurrentZone(myCoords)))
	}

	if gotTheLastPointOfZone(myCoords) {
		setNextTactic(getNextTactic())
	}
	//Checks if stuck ...

	if gotStuck(myCoords) {
		setNextTactic(getRandomTactic())
	}

	//Due to the pipeline chose the next direction ...

	for _, p := range getActionSetByTactic() {
		switch p {
		case UP:
			if NotDangerous(topItem) && NotBlindZone(top) {
				Steps = append(Steps, top)
				return direction.UP
			}
			continue
		case RIGHT:
			if NotDangerous(rightItem) && NotBlindZone(right) {
				Steps = append(Steps, right)
				return direction.RIGHT
			}
			continue
		case LEFT:
			if NotDangerous(leftItem) && NotBlindZone(left) {
				Steps = append(Steps, left)
				return direction.LEFT
			}
			continue
		case DOWN:
			if NotDangerous(bottomItem) && NotBlindZone(bottom) {
				Steps = append(Steps, bottom)
				return direction.DOWN
			}
			continue
		}
	}

	fmt.Println("there are not variants")
	return direction.Direction(1000)
}
