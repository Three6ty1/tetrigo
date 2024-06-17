package game

import (
	"log"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type IPiece struct {
	Piece
}

func NewIPiece() *IPiece {
	s, _, err := ebitenutil.NewImageFromFile(TetriminoPath + "I.png")
	if err != nil {
		log.Fatal(err)
	}
	return &IPiece{
		Piece: Piece{
			piece:       types.IPiece,
			color:       types.Mino(types.IPiece),
			orientation: types.O0,
			position:    types.NewVector(StartingX, StartingY),
			sprite:      s,
			matrix: [][]bool{
				{F, F, F, F},
				{T, T, T, T},
				{F, F, F, F},
				{F, F, F, F},
			},
		},
	}
}

func (t IPiece) Rotater(o types.Orientation) [][]bool {
	if o == types.O0 {
		return [][]bool{
			{F, F, F, F},
			{T, T, T, T},
			{F, F, F, F},
			{F, F, F, F},
		}
	} else if o == types.O90 {
		return [][]bool{
			{F, F, T, F},
			{F, F, T, F},
			{F, F, T, F},
			{F, F, T, F},
		}
	} else if o == types.O180 {
		return [][]bool{
			{F, F, F, F},
			{F, F, F, F},
			{T, T, T, T},
			{F, F, F, F},
		}
	} else {
		return [][]bool{
			{F, T, F, F},
			{F, T, F, F},
			{F, T, F, F},
			{F, T, F, F},
		}
	}
}
