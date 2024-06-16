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
	s, _, err := ebitenutil.NewImageFromFile("./assets/I.png")
	if err != nil {
		log.Fatal(err)
	}
	return &OPiece{
		Piece: Piece{
			piece:       types.OPiece,
			color:       types.Mino(types.OPiece),
			orientation: types.O0,
			position:    types.NewVector(5, 0),
			sprite:      s,
			matrix: [][]bool{
				{F, T, T, F},
				{F, T, T, F},
				{F, F, F, F},
			},
		},
	}
}

// No Rotation
func (t OPiece) Rotater(o types.Orientation) [][]bool {
	return [][]bool{
		{F, T, T, F},
		{F, T, T, F},
		{F, F, F, F},
	}
}
