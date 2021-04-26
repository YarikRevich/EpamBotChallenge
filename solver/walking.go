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

	DangerousPoints   = []game.Point{{X: 16, Y: 30}, {X: 17, Y: 30}, {X: 32, Y: 16}, {X: 1, Y: 16}, {X: 15, Y: 23}, {X: 16, Y: 23}, {X: 17, Y: 23}}
	DangerousElements = []rune{game.BULLET, game.TANK_DOWN, game.TANK_LEFT, game.TANK_RIGHT, game.TANK_UP, game.RIVER}

	ZERO_DIRECTION = direction.Direction(1000)
)

//Main hub of hero movement decisions ;) ...
func GetWayToGo(b *game.Board) direction.Direction {
	myCoords := b.GetMe()

	if isUpdatingProcess(b.GetBullets()) {

		if InBlindZone(myCoords) {
			setSpawnedInDangerousZone(true)
		}

		setIsNotAvailableZoneActive(false)
		setZoneAvailablility(getCurrentZone(myCoords))

		clearSteps()

		return ZERO_DIRECTION
	}

	if InBlindZone(myCoords) {
		setSpawnedInDangerousZone(false)
	}

	if inSafety(myCoords) {
		setIsNotAvailableZoneActive(true)
		setZoneAvailablility(getCurrentZone(myCoords))
	}

	if specleIsStart() {
		setSpecle(getTheNearestEnemy(getEnemies(b), myCoords))
	} else if n, ok := isEnemyAlive(getEnemies(b)); ok {
		fmt.Println("ALIVE")
		setSpecle(n)
	} else if _, ok := isEnemyAlive(getEnemies(b)); !ok {
		fmt.Println("NOT ALIVE")
		setSpecle(getTheNearestEnemy(getEnemies(b), myCoords))
	}

	setAction(getTheBestTactic(myCoords, getSpecle(), b))

	// if ifEnemyInAvailableZone(getTheNearestEnemy(getEnemies(b), myCoords)) {
	// 	setNextTactic(getTacticToGetTheEnemy(myCoords, getTheNearestEnemy(b.GetEnemies(), myCoords)))
	// }

	//getActionSetByTactic

	switch getAction() {
	case UP:
		return direction.UP
	case RIGHT:
		return direction.RIGHT
	case LEFT:
		return direction.LEFT
	case DOWN:
		return direction.DOWN
	}

	fmt.Println("there are not variants")
	return ZERO_DIRECTION
}
