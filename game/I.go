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

func (t *IPiece) Rotate(isLeft bool) {
	if isLeft {
		t.orientation = t.TryRotateLeft(t.orientation)
		t.matrix = t.Rotater(t.orientation)
	} else {
		t.orientation = t.TryRotateRight(t.orientation)
		t.matrix = t.Rotater(t.orientation)
	}

}

func (t IPiece) Rotater(o types.Orientation) [][]bool {
	switch o {
	case types.O0:
		return [][]bool{
			{F, F, F, F},
			{T, T, T, T},
			{F, F, F, F},
			{F, F, F, F},
		}
	case types.O90:
		return [][]bool{
			{F, F, T, F},
			{F, F, T, F},
			{F, F, T, F},
			{F, F, T, F},
		}
	case types.O180:
		return [][]bool{
			{F, F, F, F},
			{F, F, F, F},
			{T, T, T, T},
			{F, F, F, F},
		}
	default:
		return [][]bool{
			{F, T, F, F},
			{F, T, F, F},
			{F, T, F, F},
			{F, T, F, F},
		}
	}
}
