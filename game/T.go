package game

import (
	"log"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TPiece struct {
	Piece
}

func NewTPiece() *TPiece {
	s, _, err := ebitenutil.NewImageFromFile(TetriminoPath + "T.png")
	if err != nil {
		log.Fatal(err)
	}
	return &TPiece{
		Piece: Piece{
			piece:       types.TPiece,
			color:       types.Mino(types.TPiece),
			orientation: types.O0,
			position:    types.NewVector(StartingX, StartingY),
			sprite:      s,
			matrix: [][]bool{
				{F, T, F},
				{T, T, T},
				{F, F, F},
			},
		},
	}
}

func (t TPiece) Rotater(o types.Orientation) [][]bool {
	if o == types.O0 {
		return [][]bool{
			{F, T, F},
			{T, T, T},
			{F, F, F},
		}
	} else if o == types.O90 {
		return [][]bool{
			{F, T, F},
			{F, T, T},
			{F, T, F},
		}
	} else if o == types.O180 {
		return [][]bool{
			{F, F, F},
			{T, T, T},
			{F, T, F},
		}
	} else {
		return [][]bool{
			{F, T, F},
			{T, T, F},
			{F, T, F},
		}
	}
}
