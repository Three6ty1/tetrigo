package main

import (
	_ "image/png"
	"log"

	"github.com/Three6ty1/tetrigo/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var LPiece *ebiten.Image

type Game struct {
	tick uint
	// queue
	lines     *uint32
	state     GameState
	playfield *game.PlayField
}

type GameState int32

const (
	playing GameState = iota
	win
	lose
)

func init() {
	LPiece, _, _ = ebitenutil.NewImageFromFile("./assets/l.png")

}

func (g *Game) Update() error {
	g.tick++

	if g.tick == ^uint(0) {
		g.tick = 0
	}

	// update all objects in the game...?
	// Increment the current falling block
	// Generate the next block in the queue

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.playfield.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	s := ebiten.Monitor().DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

func main() {
	ebiten.SetWindowSize(1600, 900)
	ebiten.SetWindowTitle("Tetrigo")
	var lines uint32 = 0

	g := &Game{
		lines:     &lines,
		state:     playing,
		playfield: game.NewPlayField(),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
