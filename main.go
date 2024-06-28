package main

import (
	"fmt"
	"log"
	"math"

	"github.com/Three6ty1/tetrigo/game"
	"github.com/Three6ty1/tetrigo/game/tetrimino"
	"github.com/Three6ty1/tetrigo/types"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var GameScale = 0.5
var DAS_DELAY = 10 // in Frames. Ebiten is in 60fps

type Game struct {
	tick      uint
	queue     *game.TetriminoQueue
	hold      *game.Hold
	lines     uint32
	state     GameState
	playfield *game.PlayField
	active    tetrimino.Tetrimino
	das       *game.DelayAutoShift
	lock      *game.LockDelay
}

type Direction bool

const (
	Left  = true
	Right = false
)

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

	if g.lock.IsLockActive(g.tick) {
		handleDrop(g, true)
	}

	controls(g, g.tick)

	// Natural block falling
	if g.tick%15 == 0 {
		handleDrop(g, false)
	}

	return nil
}

func handleDrop(g *Game, isHardDrop bool) bool {
	currentTetrimino := g.active
	currentPosition := currentTetrimino.GetPosition()

	if game.IsColliding(*g.playfield, currentPosition.X, currentPosition.Y+1, currentTetrimino.GetMatrix()) {

		if !isHardDrop && !g.lock.IsLockActive(g.tick) {
			g.lock.InitiateLockDelay(g.tick)
			return false
		}

		fmt.Printf("Updating stack Tick: %v\n", g.tick)
		// Drop the tetrimino on the stack
		err := g.playfield.UpdateStack(currentTetrimino)
		if err != nil {
			log.Fatal(err)
		}

		g.playfield.ClearLines()
		g.active = g.queue.Next()
		g.hold.ResetCanHold()
		g.lock.ResetLockDelay()
		return true
	} else {
		currentTetrimino.SetPosition(currentPosition.X, currentPosition.Y+1)

		return false
	}
}

func controls(g *Game, tick uint) {
	currentTetrimino := g.active
	currentPosition := currentTetrimino.GetPosition()
	currentMatrix := currentTetrimino.GetMatrix()
	// MOVE LEFT
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		if !game.IsColliding(*g.playfield, currentPosition.X+1, currentPosition.Y, currentMatrix) {
			currentTetrimino.SetPosition(currentPosition.X+1, currentPosition.Y)
			g.das.InitiateDAS(Left, g.tick)
			g.lock.IncrementLockMovement(g.tick)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.das.IsDASActive(Left, g.tick) && g.tick%2 == 0 {
		if !game.IsColliding(*g.playfield, currentPosition.X+1, currentPosition.Y, currentMatrix) {
			currentTetrimino.SetPosition(currentPosition.X+1, currentPosition.Y)
			g.lock.IncrementLockMovement(g.tick)
		}
	}
	// MOVE RIGHT
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		if !game.IsColliding(*g.playfield, currentPosition.X-1, currentPosition.Y, currentMatrix) {
			currentTetrimino.SetPosition(currentPosition.X-1, currentPosition.Y)
			g.das.InitiateDAS(Right, g.tick)
			g.lock.IncrementLockMovement(g.tick)
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.das.IsDASActive(Right, g.tick) && g.tick%2 == 0 {
		if !game.IsColliding(*g.playfield, currentPosition.X-1, currentPosition.Y, currentMatrix) {
			currentTetrimino.SetPosition(currentPosition.X-1, currentPosition.Y)
			g.lock.IncrementLockMovement(g.tick)
		}
	}
	// SOFT DROP
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && tick%2 == 0 {
		handleDrop(g, false)
		g.lock.IncrementLockMovement(g.tick)
	}
	// HARD DROP
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		for !handleDrop(g, true) {
		}
	}
	// TURN LEFT
	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		newPos, valid := game.RotateKicker(*g.playfield, g.active, Left)
		if valid {
			g.active.SetPosition(newPos.X, newPos.Y)
			g.active.Rotate(Left)
			g.lock.IncrementLockMovement(g.tick)
		}

	}
	// TURN RIGHT
	if inpututil.IsKeyJustPressed(ebiten.KeyX) || inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		newPos, valid := game.RotateKicker(*g.playfield, g.active, Right)
		if valid {
			g.active.SetPosition(newPos.X, newPos.Y)
			g.active.Rotate(Right)
			g.lock.IncrementLockMovement(g.tick)
		}
	}
	// HOLD/SWAP PIECE
	if inpututil.IsKeyJustPressed(ebiten.KeyShift) || inpututil.IsKeyJustPressed(ebiten.KeyC) {
		if g.hold.CanHold() {
			g.active = g.hold.Swap(g.active)

			// First swap
			if g.active == nil {
				g.active = g.queue.Next()
			}

			g.lock.ResetLockDelay()
		}
	}

	// RESET BOARD (TESTING)
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.playfield = game.NewPlayField()
		g.lock.ResetLockDelay()
		g.active.SetPosition(tetrimino.StartingX, tetrimino.StartingY)
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
	s := g.active.GetSprite().Bounds().Size()

	op.GeoM.Translate(-float64(s.X)/2, -float64(s.Y)/2)
	op.GeoM.Rotate((90 * float64(g.active.GetOrientation())) * (math.Pi / 180))
	op.GeoM.Translate(float64(s.X)/2, float64(s.Y)/2)
	op.GeoM.Scale(GameScale, GameScale)

	tPosition := g.active.GetPosition()
	x := pfStart.X + (float64(minoOffset) * tPosition.X)
	y := pfStart.Y + (float64(minoOffset) * tPosition.Y)

	// Deep copy the op
	g.drawGhost(screen, pfStart, minoOffset, *op)

	op.GeoM.Translate(x, y)

	screen.DrawImage(g.active.GetSprite(), op)
}

func (g *Game) drawGhost(screen *ebiten.Image, pfStart types.Vector, minoOffset float64, op ebiten.DrawImageOptions) {
	collisionBox := g.active.GetMatrix()
	tPosition := g.active.GetPosition()
	y := tPosition.Y
	for !game.IsColliding(*g.playfield, tPosition.X, y+1, collisionBox) {
		y++
	}

	x := pfStart.X + (float64(minoOffset) * tPosition.X)
	y = pfStart.Y + (float64(minoOffset) * y)

	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleAlpha(0.4)

	screen.DrawImage(g.active.GetSprite(), &op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	s := ebiten.Monitor().DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

func main() {
	// ebiten.SetWindowSize(800, 450)
	ebiten.SetWindowSize(1600, 900)
	ebiten.SetWindowTitle("Tetrigo")
	// ebiten.SetFullscreen(true)
	g := &Game{
		lines:     0,
		state:     playing,
		playfield: game.NewPlayField(),
		queue:     game.NewTetriminoQueue(),
		hold:      game.NewHold(),
		active:    nil,
		das:       game.NewDelayAutoShift(),
		lock:      game.NewLockDelay(15, 32),
	}

	g.active = g.queue.Next()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
