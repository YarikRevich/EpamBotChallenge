package solver

import "battlecity_test/game"

func checkIfIn(e int, b []int) bool {
	for _, v := range b {
		if v == e {
			return true
		}
	}
	return false
}

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

func clearSteps() {
	Steps = []game.Point{}
}
