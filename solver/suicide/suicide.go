package suicide

import (
	"battlecity_test/action"
	"battlecity_test/game"
	"math"
)

func IsBetterToSuicide(me, target game.Point) bool {
	distance := math.Sqrt(
		math.Pow(
			float64(target.X)-float64(me.X), 2) +
			math.Pow(
				float64(target.Y)-float64(me.Y), 2))
	return distance > 20
}

func GetMoveToSoicide() action.Action {
	return action.DoNothing()
}
