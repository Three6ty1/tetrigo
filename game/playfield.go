package game

import (
	"fmt"
	"log"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var boardImg *ebiten.Image
var minoGreen *ebiten.Image
var minoRed *ebiten.Image
var minoBlue *ebiten.Image
var minoOrange *ebiten.Image
var minoPurple *ebiten.Image
var minoYellow *ebiten.Image
var minoLightBlue *ebiten.Image

type PlayField struct {
	stack          [][]types.Mino
	minoOffset     float64
	playfieldStart types.Vector
}

func NewPlayField() *PlayField {
	s := make([][]types.Mino, 20)

	for i := range s {
		s[i] = make([]types.Mino, 10)

	}

	pf := &PlayField{
		stack: s,
	}

	initImages()

	return pf
}

func (pf *PlayField) Draw(screen *ebiten.Image, gameScale float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(gameScale, gameScale)

	bw := float64(boardImg.Bounds().Dx())
	bh := float64(boardImg.Bounds().Dy())
	sw := float64(screen.Bounds().Dx())

	// 4 because scale means the original bounds is 2x larger, therefore we need /2 again
	startX := float64(sw/2 - (bw/2)*gameScale)
	startY := float64((bh / 8) * gameScale)
	op.GeoM.Translate(startX, startY)

	// Have to set this because we don't know until runtime what the offset is for the window
	if pf.minoOffset == 0.0 {
		pf.minoOffset = bw * gameScale / 12
		pf.playfieldStart = *types.NewVector(startX, startY)
	}

	screen.DrawImage(boardImg, op)

	pf.drawStack(screen, gameScale)
}

func (pf PlayField) drawStack(screen *ebiten.Image, gameScale float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(gameScale, gameScale)
	var x, y float64
	for row := 0; row < len(pf.stack); row++ {
		for col := 0; col < len(pf.stack[0]); col++ {
			if pf.stack[row][col] == types.None {
				continue
			}
			// +1 due to upper border
			x = pf.playfieldStart.X + (float64(pf.minoOffset) * float64(col+1))
			y = pf.playfieldStart.Y + (float64(pf.minoOffset) * float64(row+1))

			op.GeoM.Translate(x, y)

			switch pf.stack[row][col] {
			case types.Green:
				screen.DrawImage(minoGreen, op)
			case types.Red:
				screen.DrawImage(minoRed, op)
			case types.Blue:
				screen.DrawImage(minoBlue, op)
			case types.Orange:
				screen.DrawImage(minoOrange, op)
			case types.Purple:
				screen.DrawImage(minoPurple, op)
			case types.Yellow:
				screen.DrawImage(minoYellow, op)
			case types.LightBlue:
				screen.DrawImage(minoLightBlue, op)
			}

			op.GeoM.Translate(-x, -y)
		}
	}
}

func (pf PlayField) IsColliding(t Tetrimino) bool {
	startPos := t.GetPosition()
	collisionBox := t.GetMatrix()

	// Check every position occupied in a collision box and see if it is touching the bottom, sides or another mino
	for row := 0; row < len(collisionBox); row++ {
		realRow := int(startPos.Y) + row - 1

		for col := 0; col < len(collisionBox[0]); col++ {
			realCol := int(startPos.X) + col - 1

			// If the piece isnt colliding in this position
			if !collisionBox[row][col] {
				continue
			}

			// The piece is colliding
			// 1. Extending past stack borders
			// 2. Within borders and colliding with a placed mino
			if realRow >= len(pf.stack) || realCol >= len(pf.stack[0]) || (pf.stack[realRow][realCol] != types.None && collisionBox[row][col]) {
				return true
			}
		}
	}
	return false
}

func (pf PlayField) UpdateStack(t Tetrimino) error {
	startPos := t.GetPosition()
	collisionBox := t.GetMatrix()
	color := t.GetColor()

	// Check every position occupied in a collision box and see if it is touching the bottom, sides or another mino
	for row := 0; row < len(collisionBox); row++ {
		realRow := int(startPos.Y) + row - 1

		for col := 0; col < len(collisionBox[0]); col++ {
			realCol := int(startPos.X) + col - 1

			// If the piece isnt colliding in this position
			if !collisionBox[row][col] {
				continue
			}

			// Out of bounds
			if realRow < 0 || realCol < 0 {
				return fmt.Errorf("out of bounds: Game Over")
			}

			if pf.stack[realRow][realCol] != types.None && collisionBox[row][col] {
				return fmt.Errorf("tried to insert Tetrimino into non-empty space in stack at position: X: %v Y:%v", realRow, realCol)
			}

			pf.stack[realRow][realCol] = color
		}
	}

	return nil
}

func (pf PlayField) ClearLines() {

}

func (pf PlayField) CheckOutOfBounds() bool {
	return false
}

func initImages() {
	var currImg *ebiten.Image
	var err error

	currImg, _, err = ebitenutil.NewImageFromFile("./assets/board.png")
	if err != nil {
		log.Fatal(err)
	}
	boardImg = currImg

	currImg, _, err = ebitenutil.NewImageFromFile("./assets/Green.png")
	if err != nil {
		log.Fatal(err)
	}
	minoGreen = currImg

	currImg, _, err = ebitenutil.NewImageFromFile("./assets/Red.png")
	if err != nil {
		log.Fatal(err)
	}
	minoRed = currImg

	currImg, _, err = ebitenutil.NewImageFromFile("./assets/Blue.png")
	if err != nil {
		log.Fatal(err)
	}
	minoBlue = currImg

	currImg, _, err = ebitenutil.NewImageFromFile("./assets/Orange.png")
	if err != nil {
		log.Fatal(err)
	}
	minoOrange = currImg

	currImg, _, err = ebitenutil.NewImageFromFile("./assets/Purple.png")
	if err != nil {
		log.Fatal(err)
	}
	minoPurple = currImg

	currImg, _, err = ebitenutil.NewImageFromFile("./assets/Yellow.png")
	if err != nil {
		log.Fatal(err)
	}
	minoYellow = currImg

	currImg, _, err = ebitenutil.NewImageFromFile("./assets/LightBlue.png")
	if err != nil {
		log.Fatal(err)
	}
	minoLightBlue = currImg
}
