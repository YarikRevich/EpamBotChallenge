package solver

import (
	"battlecity_test/direction"
	"battlecity_test/game"
	"fmt"
)

var (
	Bulets [][]game.Point
	Steps  []game.Point

	SPAWNED_IN_DANGEROUS_ZONE = false

	DangerousPoints   = []game.Point{{X: 16, Y: 29}, {X: 17, Y: 29}, {X: 32, Y: 16}, {X: 0, Y: 16}, {X: 1, Y: 16}, {X: 15, Y: 23}, {X: 16, Y: 23}, {X: 17, Y: 23}}
	DangerousElements = []rune{game.BULLET, game.TANK_DOWN, game.TANK_LEFT, game.TANK_RIGHT, game.TANK_UP, game.RIVER}

	ZERO_DIRECTION = direction.Direction(1000)
)

//Main hub of hero movement decisions ;) ...
func GetWayToGo(b *game.Board) direction.Direction {

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

	if isUpdatingProcess(b.GetBullets()) {

		if InBlindZone(myCoords) {
			setSpawnedInDangerousZone(true)
		}

		setIsNotAvailableZoneActive(false)
		setZoneAvailablility(getCurrentZone(myCoords))

		clearSteps()

		return ZERO_DIRECTION
	}

	if CURRENT_TACTIC == ZERO_VALUE {
		setNextTactic(getTacticByCodeZone(getCurrentZone(myCoords)))
	}

	if InBlindZone(myCoords) {
		setSpawnedInDangerousZone(false)
	}

	if inSafety(myCoords) {
		setIsNotAvailableZoneActive(true)
		setZoneAvailablility(getCurrentZone(myCoords))
	}

	// if gotTheLastPointOfZone(myCoords) {
	// 	setNextTactic(getNextTactic())
	// }
	// Checks if stuck ...

	// if gotStuck(myCoords) {
	// 	fmt.Println("STUCK")
	// 	setNextTactic(getRandomTactic())
	// }

	if ifEnemyInAvailableZone(getTheNearestEnemy(b.GetEnemies(), myCoords)) {
		setNextTactic(getTacticToGetTheEnemy(myCoords, getTheNearestEnemy(b.GetEnemies(), myCoords)))
	}

	fmt.Println(myCoords)

	//Due to the pipeline chose the next direction ...
	// fmt.Println(getActionSetByTactic())

	for _, p := range getActionSetByTactic() {
		fmt.Println(p)
		switch p {
		case UP:
			if notDangerous(topItem) && InAvailableZone(top) && NotBlindZone(top) {
				Steps = append(Steps, top)
				return direction.UP
			}
			continue
		case RIGHT:
			if notDangerous(rightItem) && InAvailableZone(right) && NotBlindZone(right) {
				Steps = append(Steps, right)
				return direction.RIGHT
			}
			continue
		case LEFT:
			if notDangerous(leftItem) && InAvailableZone(left) && NotBlindZone(left) {
				Steps = append(Steps, left)
				return direction.LEFT
			}
			continue
		case DOWN:
			if notDangerous(bottomItem) && InAvailableZone(bottom) && NotBlindZone(bottom) {
				Steps = append(Steps, bottom)
				return direction.DOWN
			}
			continue
		}
	}

	fmt.Println("there are not variants")
	return ZERO_DIRECTION
}
