package game

import (
	"log"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ZPiece struct {
	Piece
}

func NewZPiece() *ZPiece {
	s, _, err := ebitenutil.NewImageFromFile(TetriminoPath + "Z.png")
	if err != nil {
		log.Fatal(err)
	}
	return &ZPiece{
		Piece: Piece{
			piece:       types.ZPiece,
			color:       types.Mino(types.ZPiece),
			orientation: types.O0,
			position:    types.NewVector(StartingX, StartingY),
			sprite:      s,
			matrix: [][]bool{
				{T, T, F},
				{F, T, T},
				{F, F, F},
			},
		},
	}
}

func (t ZPiece) Rotater(o types.Orientation) [][]bool {
	if o == types.O0 {
		return [][]bool{
			{T, T, F},
			{F, T, T},
			{F, F, F},
		}
	} else if o == types.O90 {
		return [][]bool{
			{F, F, T},
			{F, T, T},
			{F, T, F},
		}
	} else if o == types.O180 {
		return [][]bool{
			{F, F, F},
			{T, T, F},
			{F, T, T},
		}
	} else {
		return [][]bool{
			{F, T, F},
			{T, T, F},
			{T, F, F},
		}
	}
}
