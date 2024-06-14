package game

import (
	_ "image/png"
	"log"

	"github.com/Three6ty1/tetrigo/helper"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Orientation int32

const (
	O0 Orientation = iota
	O90
	O180
	O270
)

type Piece int32

const (
	S Piece = iota
	Z
	L
	J
	T
	O
	I
)

type Tetrimino struct {
	piece Piece
	// Position relative to the playfield array
	position    *helper.Vector
	orientation Orientation
	sprite      *ebiten.Image
}

func NewTetrimino(p Piece) *Tetrimino {
	t := &Tetrimino{
		piece:       p,
		orientation: O0,
		position:    helper.NewVector(5, 0),
	}

	var newSprite *ebiten.Image
	var err error
	switch p {
	case S:
		newSprite, _, err = ebitenutil.NewImageFromFile("./assets/S.png")
	case Z:
		newSprite, _, err = ebitenutil.NewImageFromFile("./assets/Z.png")
	case L:
		newSprite, _, err = ebitenutil.NewImageFromFile("./assets/L.png")
	case J:
		newSprite, _, err = ebitenutil.NewImageFromFile("./assets/J.png")
	case T:
		newSprite, _, err = ebitenutil.NewImageFromFile("./assets/T.png")
	case O:
		newSprite, _, err = ebitenutil.NewImageFromFile("./assets/O.png")
	case I:
		newSprite, _, err = ebitenutil.NewImageFromFile("./assets/I.png")
	}

	if err != nil {
		log.Fatal(err)
	}

	t.sprite = newSprite

	return t
}

func (t Tetrimino) Draw(screen *ebiten.Image, pf *PlayField, gameScale float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(gameScale, gameScale)

	// +1 due to border
	x := pf.playfieldStart.X + (float64(pf.minoOffset) * (t.position.X))
	y := pf.playfieldStart.Y + (float64(pf.minoOffset) * (t.position.Y + 1.0))

	op.GeoM.Translate(x, y)

	screen.DrawImage(t.sprite, op)
}

func (t Tetrimino) RotateLeft(pf PlayField) {

}

func (t Tetrimino) RotateRight(pf PlayField) {

}

func (t Tetrimino) GetPosition() {

}

func (t Tetrimino) SetPosition() {

}
