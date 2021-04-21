package solver

// import (
// 	// "battlecity/action"
// 	"battlecity_test/direction"
// 	"battlecity_test/game"
// 	"errors"
// )

// type Middleware func(*game.Board) (direction.Direction, error)

// func FreePlaceMiddlware() Middleware {
// 	return Middleware(func(b *game.Board) (direction.Direction, error) {
// 		myCoords := b.GetMe()

		

// 		top := game.Point{X: myCoords.X, Y: myCoords.Y + 1}
// 		bottom := game.Point{X: myCoords.X, Y: myCoords.Y - 1}
// 		right := game.Point{X: myCoords.X + 1, Y: myCoords.Y}
// 		left := game.Point{X: myCoords.X - 1, Y: myCoords.Y}

// 		switch {
// 		case NotDangerous(b.GetAt(top)) && !UP_LOCKED:
// 			UP_LOCKED = true
// 			DOWN_LOCKED = true
// 			LEFT_LOCKED = true
// 			RIGHT_LOCKED = false
// 			return direction.UP, nil
// 		case NotDangerous(b.GetAt(bottom)) && !DOWN_LOCKED:
// 			UP_LOCKED = false
// 			DOWN_LOCKED = true
// 			LEFT_LOCKED = true
// 			RIGHT_LOCKED = true
// 			return direction.DOWN, nil
// 		case NotDangerous(b.GetAt(right)) && !RIGHT_LOCKED:
// 			UP_LOCKED = false
// 			DOWN_LOCKED = false
// 			LEFT_LOCKED = true
// 			RIGHT_LOCKED = true
// 			return direction.RIGHT, nil
// 		case NotDangerous(b.GetAt(left)) && !LEFT_LOCKED:
// 			UP_LOCKED = false
// 			DOWN_LOCKED = false
// 			LEFT_LOCKED = true
// 			RIGHT_LOCKED = true
// 			return direction.LEFT, nil
// 		}

// 		return direction.Direction(0), errors.New("there are no free place arround")
// 	})
// }

// func EnemiesMiddleware() Middleware {
// 	return Middleware(func(b *game.Board) (direction.Direction, error) {

// 		// if b.IsAnyAt(myCoords, []game.Element{game.OTHER_TANK_DOWN}) {
// 		// 	return direction.UP, nil
// 		// }

// 		// if b.IsAnyAt(myCoords, []game.Element{game.OTHER_TANK_UP}) {
// 		// 	return direction.DOWN, nil
// 		// }

// 		// if b.IsAnyAt(myCoords, []game.Element{game.OTHER_TANK_LEFT}) {
// 		// 	return direction.RIGHT, nil
// 		// }

// 		// if b.IsAnyAt(myCoords, []game.Element{game.OTHER_TANK_RIGHT}) {
// 		// 	return direction.LEFT, nil
// 		// }

// 		return direction.Direction(0), errors.New("there are no tanks arround")
// 	})
// }

// func RiverMiddleware() string {
// 	return ""
// }

// func ApplyMiddlewares(b *game.Board, m ...Middleware) (*game.Board, []Middleware) {
// 	return b, m
// }

// func RunMiddlewares(b *game.Board, m []Middleware) (direction.Direction, error) {
// 	for _, v := range m {
// 		a, err := v(b)
// 		if err != nil {
// 			continue
// 		}
// 		return a, nil
// 	}
// 	return direction.Direction(0), errors.New("there are not available actions for middlewares")
// }
