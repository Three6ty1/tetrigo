package game

import (
	"github.com/Three6ty1/tetrigo/types"
)

func getOrientationData(piece types.Piece, o types.Orientation) [][]int32 {
	var data [][]int32
	if piece == types.IPiece {
		switch o {
		case types.O0:
			data = [][]int32{{0, 0}, {-1, 0}, {+2, 0}, {-1, 0}, {+2, 0}}
		case types.O90:
			data = [][]int32{{-1, 0}, {0, 0}, {0, 0}, {0, +1}, {0, -2}}
		case types.O180:
			data = [][]int32{{-1, +1}, {+1, +1}, {-2, +1}, {+1, 0}, {-2, 0}}
		case types.O270:
			data = [][]int32{{0, +1}, {0, +1}, {0, +1}, {0, -1}, {0, +2}}
		}
	} else {
		switch o {
		case types.O0:
			data = [][]int32{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
		case types.O90:
			data = [][]int32{{0, 0}, {+1, 0}, {+1, -1}, {0, +2}, {+1, +2}}
		case types.O180:
			data = [][]int32{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
		case types.O270:
			data = [][]int32{{0, 0}, {-1, 0}, {-1, -1}, {0, +2}, {-1, +2}}
		}
	}
	return data
}

func getOffsetData(piece types.Piece, from types.Orientation, to types.Orientation) [][]int32 {
	offsets := getOrientationData(piece, from)
	toData := getOrientationData(piece, to)

	// 5 offsets
	for i := 0; i < 5; i++ {
		offsets[i][0] -= toData[i][0]
		offsets[i][1] -= toData[i][1]
	}

	return offsets
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

func RotateKicker(pf PlayField, t Tetrimino, isLeft bool) (types.Vector, bool) {
	from := t.GetOrientation()
	var to types.Orientation
	if isLeft {
		to = t.TryRotateLeft(from)
	} else {
		to = t.TryRotateRight(from)
	}

	collisionBox := t.Rotater(to)

	startPos := t.GetPosition()
	var x, y float64

	offsetData := getOffsetData(types.Piece(t.GetColor()), from, to)

	for i := 0; i < 5; i++ {
		x = startPos.X - float64(offsetData[i][0])
		y = startPos.Y - float64(offsetData[i][1])

		if !IsColliding(pf, x, y, collisionBox) {
			return *types.NewVector(x, y), true
		}
	}

	return *types.NewVector(-1, -1), false
}
