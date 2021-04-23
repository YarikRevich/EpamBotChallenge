package solver

import "math/rand"

func getTacticByCodeZone(zone int) int {
	switch zone {
	case FIRST_ZONE_CODE:
		return FIRST_TACTIC
	case SECOND_ZONE_CODE:
		return SECOND_TACTIC
	case THIRD_ZONE_CODE:
		return THIRD_TACTIC
	case FORTH_ZONE_CODE:
		return FORTH_TACTIC
	}
	return ZERO_VALUE
}

//Due to the previous tactic returns the next one ...
func getNextTactic() int {
	switch CURRENT_TACTIC {
	case FIRST_TACTIC:
		return SECOND_TACTIC
	case SECOND_TACTIC:
		return THIRD_TACTIC
	case THIRD_TACTIC:
		return FORTH_TACTIC
	}
	return ZERO_VALUE
}

func setNextTactic(tactic int) {
	CURRENT_TACTIC = tactic
}

//Due to the tactic name returns the action set
func getActionSetByTactic() []string {
	switch CURRENT_TACTIC {
	case FIRST_TACTIC:
		return []string{UP, LEFT, RIGHT}
	case SECOND_TACTIC:
		return []string{UP, RIGHT, DOWN}
	case THIRD_TACTIC:
		return []string{DOWN, LEFT, RIGHT}
	case FORTH_TACTIC:
		return []string{UP, RIGHT, LEFT}
	case FIRST_TACTIC_ENEMY:
		return []string{RIGHT, LEFT, UP, DOWN}
	case SECOND_TACTIC_ENEMY:
		return []string{LEFT, RIGHT, UP, DOWN}
	case THIRD_TACTIC_ENEMY:
		return []string{UP, RIGHT, LEFT, DOWN}
	case FORTH_TACTIC_ENEMY:
		return []string{UP, LEFT, RIGHT, DOWN}
	case FIFTH_TACTIC_ENEMY:
		return []string{DOWN, RIGHT, LEFT, UP}
	case SIXTH_TACTIC_ENEMY:
		return []string{DOWN, LEFT, RIGHT, UP}
	}
	return nil
}



func getRandomTactic() int {
	var a []int
	f := []int{FIRST_TACTIC, SECOND_TACTIC, THIRD_TACTIC, FORTH_TACTIC}
	e := []int{FIRST_TACTIC_ENEMY, SECOND_TACTIC_ENEMY, THIRD_TACTIC_ENEMY, FORTH_TACTIC_ENEMY, FIFTH_TACTIC_ENEMY, SIXTH_TACTIC_ENEMY}

	if checkIfIn(CURRENT_TACTIC, f) {
		a = f
	} else {
		a = e
	}

	for {
		c := a[rand.Intn(3)]
		if c != CURRENT_TACTIC {
			return c
		}
	}
}