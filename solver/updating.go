package solver

import "battlecity_test/game"

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