package solver

// import (


// 	"battlecity_test/direction"
// 	"battlecity_test/game"
// )

// var (
// 	Bulets [][]game.Point
// 	Steps  []game.Point

// 	SPAWNED_IN_DANGEROUS_ZONE = false

// 	DangerousPoints   = []game.Point{{X: 16, Y: 30}, {X: 17, Y: 30}, {X: 32, Y: 16}, {X: 1, Y: 16}, {X: 15, Y: 23}, {X: 16, Y: 23}, {X: 17, Y: 23}}
// 	DangerousElements = []rune{game.BULLET, game.TANK_DOWN, game.TANK_LEFT, game.TANK_RIGHT, game.TANK_UP, game.RIVER}

// 	ZERO_DIRECTION = direction.Direction(1000)
// )

// //Main hub of hero movement decisions ;) ...
// func GetWayToGo(b *game.Board)direction.Direction {
// 	myCoords := b.GetMe()

// 	if isUpdatingProcess(b.GetBullets()) {

// 		if InBlindZone(myCoords) {
// 			setSpawnedInDangerousZone(true)
// 		}

// 		setIsNotAvailableZoneActive(false)
// 		setZoneAvailablility(getCurrentZone(myCoords))

// 		return direction.NONE
// 	}

// 	if InBlindZone(myCoords) {
// 		setSpawnedInDangerousZone(false)
// 	}

// 	if inSafety(myCoords) {
// 		setIsNotAvailableZoneActive(true)
// 		setZoneAvailablility(getCurrentZone(myCoords))
// 	}

// 	switch getTheBestTactic(myCoords, getTheNearestEnemy(b.GetEnemies(), myCoords), b) {
// 	case UP:
// 		return direction.UP
// 	case RIGHT:
// 		return direction.RIGHT
// 	case LEFT:
// 		return direction.LEFT
// 	case DOWN:
// 		return direction.DOWN
// 	}

// 	return direction.NONE
// }
