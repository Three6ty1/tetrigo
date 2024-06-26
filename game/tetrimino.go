package game

import (
	_ "image/png"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2"
)

const T = true
const F = false

const StartingX = 5
const StartingY = 0
const TetriminoPath = "./assets/tetriminos/"

type Tetrimino interface {
	GetColor() types.Mino
	GetPosition() *types.Vector
	SetPosition(x float64, y float64)
	GetMatrix() [][]bool
	Rotate(isLeft bool)
	Rotater(o types.Orientation) [][]bool
	TryRotateLeft(o types.Orientation) types.Orientation
	TryRotateRight(o types.Orientation) types.Orientation
	GetSprite() *ebiten.Image
	GetOrientation() types.Orientation
	SetOrientation(o types.Orientation)
}

type Piece struct {
	piece types.Piece
	color types.Mino
	// Position relative to the playfield array
	position    *types.Vector
	orientation types.Orientation
	sprite      *ebiten.Image
	matrix      [][]bool
}

func NewTetrimino(p types.Piece) Tetrimino {
	var t Tetrimino
	switch p {
	case types.SPiece:
		t = NewSPiece()
	case types.ZPiece:
		t = NewZPiece()
	case types.LPiece:
		t = NewLPiece()
	case types.JPiece:
		t = NewJPiece()
	case types.TPiece:
		t = NewTPiece()
	case types.OPiece:
		t = NewOPiece()
	case types.IPiece:
		t = NewIPiece()
	}

	return t
}

func (t Piece) GetOrientation() types.Orientation {
	return t.orientation
}

func (t Piece) SetOrientation(o types.Orientation) {
	t.orientation = o
}

func (t Piece) GetColor() types.Mino {
	return t.color
}

func (t Piece) GetPosition() *types.Vector {
	return t.position
}

func (t *Piece) SetPosition(x float64, y float64) {
	t.position.X = x
	t.position.Y = y
}

func (t Piece) GetMatrix() [][]bool {
	return t.matrix
}

func (t Piece) GetSprite() *ebiten.Image {
	return t.sprite
}

func (t Piece) TryRotateLeft(o types.Orientation) types.Orientation {
	new := t.orientation - 1
	if new < 0 {
		new = types.O270
	}

	return new
}

func (t Piece) TryRotateRight(o types.Orientation) types.Orientation {
	new := t.orientation + 1
	if new > types.O270 {
		new = types.O0
	}

	return new
}
