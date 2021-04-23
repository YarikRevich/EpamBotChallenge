package solver

import "battlecity_test/game"

//IT IS ONLY ABOUT DENIED ZONE TRAKING ...

func inSafety(p game.Point) bool{
	return p.Y >= NOT_AVAILABLE_ZONE.Y
}

func InAvailableZone(p game.Point) bool {
	if IS_NOT_AVAILABLE_ZONE_ACTIVE {
		return p.Y >= NOT_AVAILABLE_ZONE.Y
	}
	return true
}

func setIsNotAvailableZoneActive(r bool){
	IS_NOT_AVAILABLE_ZONE_ACTIVE = r
}

func setZoneAvailablility(z int) {
	if !IS_NOT_AVAILABLE_ZONE_ACTIVE {
		if z == FIRST_ZONE_CODE || z == FORTH_ZONE_CODE {
			IS_NOT_AVAILABLE_ZONE_ACTIVE = false
		} else {
			IS_NOT_AVAILABLE_ZONE_ACTIVE = true
		}
	}
}
