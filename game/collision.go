package game

import "github.com/Three6ty1/tetrigo/types"

func getOffsetData(piece types.Piece, from types.Orientation, to types.Orientation) [][]int32 {
	return nil
}

func IsColliding(pf PlayField, startX float64, startY float64, collisionBox [][]bool) bool {
	// Check every position occupied in a collision box and see if it is touching the bottom, sides or another mino
	for row := 0; row < len(collisionBox); row++ {
		realRow := int(startY) + row - 1

		for col := 0; col < len(collisionBox[0]); col++ {
			realCol := int(startX) + col - 1

			// If the piece isnt colliding in this position
			if !collisionBox[row][col] || realRow == -1 {
				continue
			}

			// The piece is colliding
			// 1. Extending past stack borders
			// 2. Within borders and colliding with a placed mino
			if realRow >= len(pf.stack) || realCol >= len(pf.stack[0]) || realCol < 0 || (pf.stack[realRow][realCol] != types.None && collisionBox[row][col]) {
				return true
			}
		}
	}
	return false
}

// func RotateKicker(pf PlayField, float64, startY float64, collisionBox [][]bool) (float64, float64, bool) {
// 	var x, y float64
// 	const dir = BadExpr
// }
