package target

import (
	"battlecity_test/game"
	"math"
)

const (
	maxDistBetweenDiffTypeEnemies = 5
)

type Target struct {
	b *game.Board
}

func (t *Target) getDistanceToPoint(p game.Point) int {
	me := t.b.GetMe()

	return int(math.Sqrt(
		math.Pow(
			float64(p.X) - float64(me.X), 2) +
		math.Pow(
				float64(p.Y) - float64(me.Y), 2)))
}

func (t *Target) getNearestEnemy(enemyTypes ...game.Element) game.Point{
	var (
		nearest game.Point
		distanceToNearest int)

	for _, p := range t.b.GetAllPoints(enemyTypes...){
		if distanceToNearest == 0 || t.getDistanceToPoint(p) < distanceToNearest{
			nearest = p
		}
	}

	return nearest
}

func (t *Target) getNearestAIEnemy() game.Point {
	return t.getNearestEnemy(
		game.AI_TANK_DOWN,
		game.AI_TANK_LEFT,
		game.AI_TANK_PRIZE,
		game.AI_TANK_RIGHT,
		game.AI_TANK_UP,
	)
}

func (t *Target) getNearestLiveEnemy() game.Point {
	return t.getNearestEnemy(
		game.TANK_DOWN,
		game.TANK_LEFT,
		game.TANK_RIGHT,
		game.TANK_UP,
	)
}

func (t *Target) GetTarget() game.Point {
	aiEnemy := t.getNearestAIEnemy()
	aiEnemyDistance := t.getDistanceToPoint(aiEnemy)

	liveEnemy := t.getNearestLiveEnemy()
	liveEnemyDistance := t.getDistanceToPoint(liveEnemy)
	if math.Abs(float64(aiEnemyDistance-liveEnemyDistance)) <= maxDistBetweenDiffTypeEnemies {
		return liveEnemy
	}
	return aiEnemy
}

func New(b *game.Board) *Target {
	return &Target{
		b: b}
}
