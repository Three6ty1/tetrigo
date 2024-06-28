package tetrimino

import (
	"log"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type LPiece struct {
	Piece
}

func (t *LPiece) GetAltSprite() *ebiten.Image {
	return t.GetSprite()
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

func (t *LPiece) Rotate(isLeft bool) {
	if isLeft {
		t.orientation = t.TryRotateLeft(t.orientation)
		t.matrix = t.Rotater(t.orientation)
	} else {
		t.orientation = t.TryRotateRight(t.orientation)
		t.matrix = t.Rotater(t.orientation)
	}

}

func (t LPiece) Rotater(o types.Orientation) [][]bool {
	switch o {
	case types.O0:
		return [][]bool{
			{F, F, T},
			{T, T, T},
			{F, F, F},
		}
	case types.O90:
		return [][]bool{
			{F, T, F},
			{F, T, F},
			{F, T, T},
		}
	case types.O180:
		return [][]bool{
			{F, F, F},
			{T, T, T},
			{T, F, F},
		}
	default:
		return [][]bool{
			{T, T, F},
			{F, T, F},
			{F, T, F},
		}
	}
}
