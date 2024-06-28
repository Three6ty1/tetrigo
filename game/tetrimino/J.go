package tetrimino

import (
	"log"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2"
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

func (t *JPiece) GetAltSprite() *ebiten.Image {
	return t.GetSprite()
}

func (t *JPiece) Rotate(isLeft bool) {
	if isLeft {
		t.orientation = t.TryRotateLeft(t.orientation)
		t.matrix = t.Rotater(t.orientation)
	} else {
		t.orientation = t.TryRotateRight(t.orientation)
		t.matrix = t.Rotater(t.orientation)
	}

}

func (t JPiece) Rotater(o types.Orientation) [][]bool {
	switch o {
	case types.O0:
		return [][]bool{
			{T, F, F},
			{T, T, T},
			{F, F, F},
		}
	case types.O90:
		return [][]bool{
			{F, T, T},
			{F, T, F},
			{F, T, F},
		}
	case types.O180:
		return [][]bool{
			{F, F, F},
			{T, T, T},
			{F, F, T},
		}
	default:
		return [][]bool{
			{F, T, F},
			{F, T, F},
			{T, T, F},
		}
	}
}
