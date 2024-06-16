package game

import (
	_ "image/png"
	"log"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tetrimino interface {
	Draw(screen *ebiten.Image, pf *PlayField, gameScale float64)
	GetPosition() *types.Vector
	SetPosition(x float64, y float64)
	GetMatrix() [][]bool
	Rotater(types.Orientation) [][]bool
	RotateLeft()
	RotateRight()
	TryRotateLeft() ([][]bool, types.Orientation)
	TryRotateRight() ([][]bool, types.Orientation)
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

	case types.ZPiece:

	case types.LPiece:

	case types.JPiece:

	case types.TPiece:

	case types.OPiece:

	case types.IPiece:
		t = NewIPiece()
	}

	return t
}

func (t Piece) Draw(screen *ebiten.Image, pf *PlayField, gameScale float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(gameScale, gameScale)

	// +1 due to border
	x := pf.playfieldStart.X + (float64(pf.minoOffset) * (t.position.X))
	y := pf.playfieldStart.Y + (float64(pf.minoOffset) * (t.position.Y + 1.0))

	op.GeoM.Translate(x, y)

	screen.DrawImage(t.sprite, op)
}

func (t Piece) GetPosition() *types.Vector {
	return t.position
}

func (t Piece) SetPosition(x float64, y float64) {
	t.position.X = x
	t.position.Y = y
}

func (t Piece) GetMatrix() [][]bool {
	return t.matrix
}

func (t Piece) Rotater(o types.Orientation) [][]bool {
	log.Fatal("Error: Rotator not implemented")
	return [][]bool{}
}

func (t Piece) RotateLeft() {
	t.matrix, t.orientation = t.TryRotateLeft()
}

func (t Piece) RotateRight() {
	t.matrix, t.orientation = t.TryRotateRight()
}

func (t Piece) TryRotateLeft() ([][]bool, types.Orientation) {
	o := t.orientation - 1
	if o < 0 {
		o = types.O270
	}

	return t.Rotater(o), o
}

func (t Piece) TryRotateRight() ([][]bool, types.Orientation) {
	o := t.orientation + 1
	if o > types.O270 {
		o = types.O0
	}

	return t.Rotater(o), o
}
