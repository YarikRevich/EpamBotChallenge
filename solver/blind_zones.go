package solver

import "battlecity_test/game"

//IT IS ONLY ABOUT BLIND ZONES TRAKING ....

func InBlindZone(e game.Point) bool {
	for _, v := range DangerousPoints {
		if e == v {
			return false
		}
	}
	return true
}

func setSpawnedInDangerousZone(r bool){
	SPAWNED_IN_DANGEROUS_ZONE = r
}

func NotBlindZone(e game.Point) bool {
	if !SPAWNED_IN_DANGEROUS_ZONE {
		for _, v := range DangerousPoints {
			if e == v {
				return false
			}
		}
	}
	return true
}
