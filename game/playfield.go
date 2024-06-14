package game

import (
	"log"

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
	stack [][]Mino
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

func (pf PlayField) Draw(screen *ebiten.Image) {
	boardImg, _, err := ebitenutil.NewImageFromFile("./assets/board.png")
	if err != nil {
		log.Fatal(err)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	scale := ebiten.Monitor().DeviceScaleFactor()
	op.GeoM.Scale(scale, scale)

	bw := boardImg.Bounds().Dx()
	bh := boardImg.Bounds().Dy()
	sw := screen.Bounds().Dx()

	// 4 because scale means the original bounds is 2x larger, therefore we need /2 again
	op.GeoM.Translate(float64(sw/2-bw/4), float64(bh/16))

	screen.DrawImage(boardImg, op)
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
