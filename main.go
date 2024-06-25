package main

import (
	"fmt"
	"log"

	"github.com/Three6ty1/tetrigo/game"
	"github.com/Three6ty1/tetrigo/types"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var GameScale = 0.25

type Game struct {
	tick      uint
	queue     *game.TetriminoQueue
	hold      *game.Hold
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

	controls(g, g.tick)

	// Natural block falling
	if g.tick%15 == 0 {
		handleDrop(g)
	}

	return nil
}

func handleDrop(g *Game) {
	currentTetrimino := g.active
	currentPosition := currentTetrimino.GetPosition()

	if game.IsColliding(*g.playfield, currentPosition.X, currentPosition.Y+1, currentTetrimino.GetMatrix()) {
		fmt.Printf("Tick: %v\n", g.tick)
		fmt.Println("Colliding")

		// TODO: Switch to delay based hard drop
		// Drop the tetrimino on the stack
		fmt.Println("Updating Stack")
		err := g.playfield.UpdateStack(currentTetrimino)
		if err != nil {
			log.Fatal(err)
		}

		// Queue up the next tetrimino
		g.active = g.queue.Next()
		g.hold.ResetCanHold()

	} else {
		currentTetrimino.SetPosition(currentPosition.X, currentPosition.Y+1)
	}
}

func controls(g *Game, tick uint) {
	currentTetrimino := g.active
	currentPosition := currentTetrimino.GetPosition()
	currentMatrix := currentTetrimino.GetMatrix()
	// TODO: Change to hotkeys
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && tick%3 == 0 {
		if !game.IsColliding(*g.playfield, currentPosition.X+1, currentPosition.Y, currentMatrix) {
			currentTetrimino.SetPosition(currentPosition.X+1, currentPosition.Y)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && tick%3 == 0 {
		if !game.IsColliding(*g.playfield, currentPosition.X-1, currentPosition.Y, currentMatrix) {
			currentTetrimino.SetPosition(currentPosition.X-1, currentPosition.Y)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && tick%2 == 0 {
		handleDrop(g)

	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		current := g.active
		for current == g.active {
			handleDrop(g)
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		newPos, valid := game.RotateKicker(*g.playfield, g.active, true)
		if valid {
			g.active.SetPosition(newPos.X, newPos.Y)
			g.active.Rotate(true)
		}

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyX) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		newPos, valid := game.RotateKicker(*g.playfield, g.active, false)
		if valid {
			g.active.SetPosition(newPos.X, newPos.Y)
			g.active.Rotate(false)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyShift) {
		if g.hold.CanHold() {
			g.active = g.hold.Swap(g.active)

			// First swap
			if g.active == nil {
				g.active = g.queue.Next()
			}
		}
	}
	// } else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {

	// }
}

func (g *Game) Draw(screen *ebiten.Image) {
	pfStart := g.playfield.GetPlayFieldStart()
	minoOffset := g.playfield.GetMinoOffset()

	g.playfield.Draw(screen, GameScale)
	g.DrawActive(screen, pfStart, minoOffset)
	g.queue.Draw(screen, pfStart, minoOffset, GameScale)
	g.hold.Draw(screen, pfStart, minoOffset, GameScale)
}

func (g *Game) DrawActive(screen *ebiten.Image, pfStart types.Vector, minoOffset float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(GameScale, GameScale)

	tPosition := g.active.GetPosition()
	x := pfStart.X + (float64(minoOffset) * tPosition.X)
	y := pfStart.Y + (float64(minoOffset) * tPosition.Y)

	op.GeoM.Translate(x, y)

	screen.DrawImage(g.active.GetSprite(), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	s := ebiten.Monitor().DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

func main() {
	ebiten.SetWindowSize(800, 450)
	ebiten.SetWindowTitle("Tetrigo")

	g := &Game{
		lines:     0,
		state:     playing,
		playfield: game.NewPlayField(),
		queue:     game.NewTetriminoQueue(),
		hold:      game.NewHold(),
		active:    nil,
	}

	g.active = g.queue.Next()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
