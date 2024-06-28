package tetrimino

import (
	"log"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ZPiece struct {
	Piece
}

func (t *ZPiece) GetAltSprite() *ebiten.Image {
	return t.GetSprite()
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

func (t *ZPiece) Rotate(isLeft bool) {
	if isLeft {
		t.orientation = t.TryRotateLeft(t.orientation)
		t.matrix = t.Rotater(t.orientation)
	} else {
		t.orientation = t.TryRotateRight(t.orientation)
		t.matrix = t.Rotater(t.orientation)
	}

}

func (t ZPiece) Rotater(o types.Orientation) [][]bool {
	switch o {
	case types.O0:
		return [][]bool{
			{T, T, F},
			{F, T, T},
			{F, F, F},
		}
	case types.O90:
		return [][]bool{
			{F, F, T},
			{F, T, T},
			{F, T, F},
		}
	case types.O180:
		return [][]bool{
			{F, F, F},
			{T, T, F},
			{F, T, T},
		}
	default:
		return [][]bool{
			{F, T, F},
			{T, T, F},
			{T, F, F},
		}
	}
}
