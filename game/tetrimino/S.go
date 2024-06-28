package tetrimino

import (
	"log"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SPiece struct {
	Piece
}

func (t *SPiece) GetAltSprite() *ebiten.Image {
	return t.GetSprite()
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

func (t *SPiece) Rotate(isLeft bool) {
	if isLeft {
		t.orientation = t.TryRotateLeft(t.orientation)
		t.matrix = t.Rotater(t.orientation)
	} else {
		t.orientation = t.TryRotateRight(t.orientation)
		t.matrix = t.Rotater(t.orientation)
	}

}

func (t SPiece) Rotater(o types.Orientation) [][]bool {
	switch o {
	case types.O0:
		return [][]bool{
			{F, T, T},
			{T, T, F},
			{F, F, F},
		}
	case types.O90:
		return [][]bool{
			{F, T, F},
			{F, T, T},
			{F, F, T},
		}
	case types.O180:
		return [][]bool{
			{F, F, F},
			{F, T, T},
			{T, T, F},
		}
	default:
		return [][]bool{
			{T, F, F},
			{T, T, F},
			{F, T, F},
		}
	}
}
