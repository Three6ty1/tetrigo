package game

import (
	"log"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type JPiece struct {
	Piece
}

func NewJPiece() *JPiece {
	s, _, err := ebitenutil.NewImageFromFile(TetriminoPath + "J.png")
	if err != nil {
		log.Fatal(err)
	}
	return &JPiece{
		Piece: Piece{
			piece:       types.JPiece,
			color:       types.Mino(types.JPiece),
			orientation: types.O0,
			position:    types.NewVector(StartingX, StartingY),
			sprite:      s,
			matrix: [][]bool{
				{T, F, F},
				{T, T, T},
				{F, F, F},
			},
		},
	}
}

func (t JPiece) Rotater(o types.Orientation) [][]bool {
	if o == types.O0 {
		return [][]bool{
			{T, F, F},
			{T, T, T},
			{F, F, F},
		}
	} else if o == types.O90 {
		return [][]bool{
			{F, T, T},
			{F, T, F},
			{F, T, F},
		}
	} else if o == types.O180 {
		return [][]bool{
			{F, F, F},
			{T, T, T},
			{F, F, T},
		}
	} else {
		return [][]bool{
			{F, T, F},
			{F, T, F},
			{T, T, F},
		}
	}
}
