package game

import (
	"log"

	"github.com/Three6ty1/tetrigo/helper"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var boardImg *ebiten.Image

type Mino int32

const (
	blue Mino = iota
	green
	lightBlue
	orange
	purple
	red
	yellow
	empty
)

type PlayField struct {
	stack          [][]Mino
	minoOffset     float64
	playfieldStart helper.Vector
}

func NewPlayField() *PlayField {
	s := make([][]Mino, 20)

	for i := range s {
		s[i] = make([]Mino, 10)
	}

	pf := &PlayField{
		stack: s,
	}

	return pf
}

func (pf *PlayField) Draw(screen *ebiten.Image, gameScale float64) {
	boardImg, _, err := ebitenutil.NewImageFromFile("./assets/board.png")

	if err != nil {
		log.Fatal(err)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(gameScale, gameScale)

	bw := float64(boardImg.Bounds().Dx())
	bh := float64(boardImg.Bounds().Dy())
	sw := float64(screen.Bounds().Dx())

	// 4 because scale means the original bounds is 2x larger, therefore we need /2 again
	startX := float64(sw/2 - (bw/2)*gameScale)
	startY := float64((bh / 8) * gameScale)
	op.GeoM.Translate(startX, startY)

	screen.DrawImage(boardImg, op)

	if pf.minoOffset == 0.0 {
		// Set the playfield variables during runtime
		// Board with borders is 12 * 22
		pf.minoOffset = bw * gameScale / 12
		pf.playfieldStart = *helper.NewVector(startX, startY)
	}
}

func (pf PlayField) IsColliding(t Tetrimino) bool {
	return false
}

func (pf PlayField) UpdateStack(t Tetrimino) {

}

func (pf PlayField) ClearLines() {

}

func (pf PlayField) CheckOutOfBounds() bool {
	return false
}
