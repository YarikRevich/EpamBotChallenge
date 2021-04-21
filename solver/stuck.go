package solver

import "battlecity_test/game"

func rightTopStuckAngle(top, right game.Element) bool {
	//Checks whether hero is stuck in the right top angle

	return (top != game.NONE && right != game.NONE) || (top == game.ICE && right == game.ICE)
}

func leftTopStuckAngle(top, left game.Element) bool {
	//Checks whether hero is stuck in the left top angle
	
	return (top != game.NONE && left != game.NONE) || (top == game.ICE && left == game.ICE)
}

func rightBottomStuckAngle(bottom, right game.Element) bool {
	//Checks whether hero is stuck in the right bottom angle

	return (bottom != game.NONE && right != game.NONE) || (bottom == game.ICE && right == game.ICE)
}

func leftBottomStuckAngle(bottom, left game.Element) bool {
	//Checks whether hero is stuck in the left bottom angle

	return (bottom != game.NONE && left != game.NONE) || (bottom == game.ICE && left == game.ICE)
}