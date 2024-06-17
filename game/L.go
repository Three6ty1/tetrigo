package game

import (
	"log"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type LPiece struct {
	Piece
}

func NewLPiece() *LPiece {
	s, _, err := ebitenutil.NewImageFromFile(TetriminoPath + "L.png")
	if err != nil {
		log.Fatal(err)
	}
	return &LPiece{
		Piece: Piece{
			piece:       types.LPiece,
			color:       types.Mino(types.LPiece),
			orientation: types.O0,
			position:    types.NewVector(StartingX, StartingY),
			sprite:      s,
			matrix: [][]bool{
				{F, F, T},
				{T, T, T},
				{F, F, F},
			},
		},
	}
}

func (t LPiece) Rotater(o types.Orientation) [][]bool {
	if o == types.O0 {
		return [][]bool{
			{F, F, T},
			{T, T, T},
			{F, F, F},
		}
	} else if o == types.O90 {
		return [][]bool{
			{F, T, F},
			{F, T, F},
			{F, T, T},
		}
	} else if o == types.O180 {
		return [][]bool{
			{F, F, F},
			{T, T, T},
			{T, F, F},
		}
	} else {
		return [][]bool{
			{T, T, F},
			{F, T, F},
			{F, T, F},
		}
	}
}
