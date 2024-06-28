package game

import (
	"log"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2"
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

func (t *TPiece) GetAltSprite() *ebiten.Image {
	return t.GetSprite()
}

func (t *TPiece) Rotate(isLeft bool) {
	if isLeft {
		t.orientation = t.TryRotateLeft(t.orientation)
		t.matrix = t.Rotater(t.orientation)
	} else {
		t.orientation = t.TryRotateRight(t.orientation)
		t.matrix = t.Rotater(t.orientation)
	}

}

func (t TPiece) Rotater(o types.Orientation) [][]bool {
	switch o {
	case types.O0:
		return [][]bool{
			{F, T, F},
			{T, T, T},
			{F, F, F},
		}
	case types.O90:
		return [][]bool{
			{F, T, F},
			{F, T, T},
			{F, T, F},
		}
	case types.O180:
		return [][]bool{
			{F, F, F},
			{T, T, T},
			{F, T, F},
		}
	default:
		return [][]bool{
			{F, T, F},
			{T, T, F},
			{F, T, F},
		}
	}
}
