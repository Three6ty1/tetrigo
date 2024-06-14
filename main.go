package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var boardImg *ebiten.Image
var LPiece *ebiten.Image

type Game struct{}

type GameState int32

const (
	playing GameState = iota
	win
	lose
)

func init() {
	LPiece, _, _ = ebitenutil.NewImageFromFile("./assets/l.png")
	boardImg, _, _ = ebitenutil.NewImageFromFile("./assets/board.png")
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	s := ebiten.Monitor().DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

func main() {
	ebiten.SetWindowSize(1600, 900)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
