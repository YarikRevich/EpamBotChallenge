package utils

import (
	"math"
	"battlecity_test/game"
)


var (
	Bulets [][]game.Point
)

func CheckEqual(a, b []game.Point) bool {
	for _, av := range a {
		for _, bv := range b {
			if av != bv {
				return false
			}
		}
	}
	return true
}

func IsBattleWall(a game.Point, b []game.Point) bool {
	for _, v := range b {
		if a == v {
			return true
		}
	}
	return false
}

func IsWithin(a game.Point, b []game.Point) bool {
	//Checks whether a within b

	for _, v := range b {
		if a == v {
			return true
		}
	}
	return false

}


func GetTheNearestElement(o game.Point, c []game.Point) game.Point {
	var theBest game.Point
	var theBestLength float64

	for _, v := range c {
		if theBest.X == 0 && theBest.Y == 0 {
			theBest = v
			theBestLength = math.Sqrt((math.Pow(float64(v.X-o.X), 2) + math.Pow(float64(v.Y-o.Y), 2)))
			continue
		}
		if n := math.Sqrt((math.Pow(float64(v.X-o.X), 2) + math.Pow(float64(v.Y-o.Y), 2))); n < theBestLength {
			theBest = v
			theBestLength = n
		}
	}
	return theBest
}



func IsUpdatingProcess(e []game.Point) bool {
	Bulets = append(Bulets, e)

	if len(Bulets) == 1 {
		return false
	}

	if len(Bulets) == 2 {
		r := CheckEqual(Bulets[len(Bulets)-1], Bulets[len(Bulets)-2])
		Bulets = [][]game.Point{}
		return r
	}
	return true
}
