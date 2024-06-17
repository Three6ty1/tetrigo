package main

import (
	"fmt"
	"log"

	"github.com/Three6ty1/tetrigo/game"
	"github.com/Three6ty1/tetrigo/types"

	"github.com/hajimehoshi/ebiten/v2"
)

var GameScale = 0.5

type Game struct {
	tick uint
	// queue
	// held tetrimino
	lines     uint32
	state     GameState
	playfield *game.PlayField
	active    game.Tetrimino
}

type GameState int32

const (
	playing GameState = iota
	win
	lose
)

func (g *Game) Update() error {
	g.tick++

	if g.tick == ^uint(0) {
		g.tick = 0
	}

	// Natural block falling
	if g.tick%5 == 0 {
		currentTetrimino := g.active
		currentPosition := currentTetrimino.GetPosition()
		currentTetrimino.SetPosition(currentPosition.X, currentPosition.Y+1)
		if g.playfield.IsColliding(currentTetrimino) {
			fmt.Println("Colliding")
			// Revert position
			currentTetrimino.SetPosition(currentPosition.X, currentPosition.Y-1)

			// TODO: Switch to delay based hard drop
			// Drop the tetrimino on the stack
			fmt.Println("Updating Stack")
			err := g.playfield.UpdateStack(currentTetrimino)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Reassigning active?")
			g.active = game.NewTPiece()
			// Automatically garbage collection yay
			// TODO: Queue up the next tetrimino
		}

	}

	// update all objects in the game...?
	// Increment the current falling block
	// Generate the next block in the queue

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.playfield.Draw(screen, GameScale)
	g.active.Draw(screen, g.playfield, GameScale)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	s := ebiten.Monitor().DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

func main() {
	ebiten.SetWindowSize(1600, 900)
	ebiten.SetWindowTitle("Tetrigo")

	g := &Game{
		lines:     0,
		state:     playing,
		playfield: game.NewPlayField(),
		active:    game.NewTetrimino(types.IPiece),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
