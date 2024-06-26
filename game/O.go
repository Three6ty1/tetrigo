package game

import (
	"log"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type OPiece struct {
	Piece
}

func NewOPiece() *OPiece {
	s, _, err := ebitenutil.NewImageFromFile(TetriminoPath + "O.png")
	if err != nil {
		log.Fatal(err)
	}
	return &OPiece{
		Piece: Piece{
			piece:       types.OPiece,
			color:       types.Mino(types.OPiece),
			orientation: types.O0,
			position:    types.NewVector(StartingX, StartingY),
			sprite:      s,
			matrix: [][]bool{
				{T, T},
				{T, T},
			},
		},
	}
}

func (t *OPiece) Rotate(isLeft bool) {
	if isLeft {
		t.orientation = t.TryRotateLeft(t.orientation)
		t.matrix = t.Rotater(t.orientation)
	} else {
		t.orientation = t.TryRotateRight(t.orientation)
		t.matrix = t.Rotater(t.orientation)
	}

}

// No Rotation
func (t OPiece) Rotater(o types.Orientation) [][]bool {
	return [][]bool{
		{T, T},
		{T, T},
	}
}
