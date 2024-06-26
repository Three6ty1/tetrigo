package game

import (
	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2"
)

type Hold struct {
	piece   Tetrimino
	canHold bool
}

func NewHold() *Hold {
	return &Hold{
		piece:   nil,
		canHold: true,
	}
}

func (h Hold) Draw(screen *ebiten.Image, pfStart types.Vector, minoOffset float64, gameScale float64) {
	if h.piece == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(gameScale, gameScale)

	x := pfStart.X - float64(minoOffset*7)
	y := pfStart.Y + minoOffset*2

	op.GeoM.Translate(x, y)

	screen.DrawImage(h.piece.GetSprite(), op)
}

func (h Hold) CanHold() bool {
	return h.canHold
}

func (h *Hold) ResetCanHold() {
	h.canHold = true
}

func (h *Hold) Swap(t Tetrimino) Tetrimino {
	new := h.piece
	h.piece = t
	t.SetPosition(StartingX, StartingY)

	// Rotate to default orientation
	for t.GetOrientation() != types.O0 {
		t.Rotate(true)
	}

	h.canHold = false
	return new
}
