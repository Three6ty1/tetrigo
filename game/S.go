package game

import (
	"log"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SPiece struct {
	Piece
}

func NewSPiece() *SPiece {
	s, _, err := ebitenutil.NewImageFromFile(TetriminoPath + "S.png")
	if err != nil {
		log.Fatal(err)
	}
	return &SPiece{
		Piece: Piece{
			piece:       types.SPiece,
			color:       types.Mino(types.SPiece),
			orientation: types.O0,
			position:    types.NewVector(StartingX, StartingY),
			sprite:      s,
			matrix: [][]bool{
				{F, T, T},
				{T, T, F},
				{F, F, F},
			},
		},
	}
}

func (t SPiece) Rotater(o types.Orientation) [][]bool {
	if o == types.O0 {
		return [][]bool{
			{F, T, T},
			{T, T, F},
			{F, F, F},
		}
	} else if o == types.O90 {
		return [][]bool{
			{F, T, F},
			{F, T, T},
			{F, F, T},
		}
	} else if o == types.O180 {
		return [][]bool{
			{F, F, F},
			{F, T, T},
			{T, T, F},
		}
	} else {
		return [][]bool{
			{T, F, F},
			{T, T, F},
			{F, T, F},
		}
	}
}
