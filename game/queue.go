package game

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// var r = rand.New(rand.NewSource(2))

var TEST = false // Enable testing variables

// Queue is 6 long, but only displays 5
type TetriminoQueue struct {
	queue []Tetrimino
	next  []Tetrimino
}

func newQueue() *[]Tetrimino {
	// https://tetris.fandom.com/wiki/Random_Generator
	// 7-piece bag
	q := make([]Tetrimino, 0)

	for i := 1; i <= 7; i++ {
		q = append(q, NewTetrimino(types.Piece(i)))
	}

	r.Shuffle(len(q), func(i, j int) {
		q[i], q[j] = q[j], q[i]
	})

	return &q
}

func NewTetriminoQueue() *TetriminoQueue {
	var tq *TetriminoQueue
	if !TEST {
		tq = &TetriminoQueue{
			queue: *newQueue(),
			next:  *newQueue(),
		}
	} else {
		tq = &TetriminoQueue{
			queue: *newTestQueue(),
			next:  *newTestQueue(),
		}
	}

	return tq
}

func (tq *TetriminoQueue) Next() Tetrimino {
	// https://stackoverflow.com/a/26863706
	next := tq.queue[0]
	tq.queue = tq.queue[1:]

	if len(tq.queue) == 0 {
		tq.queue = tq.next
		if !TEST {
			tq.next = *newQueue()
		} else {
			tq.next = *newTestQueue()
		}
	}

	return next
}

func (tq TetriminoQueue) Draw(screen *ebiten.Image, pfStart types.Vector, minoOffset float64, gameScale float64) {
	var q []Tetrimino

	q = append(q, tq.queue...)
	q = append(q, tq.next[:7-len(tq.queue)]...)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(gameScale*0.75, gameScale*0.75)

	// Move to top right corner of the playfield and then some more
	qStartX := pfStart.X + (minoOffset * 12)
	qStartY := pfStart.Y

	op.GeoM.Translate(qStartX, qStartY)

	var y float64

	for i := 0; i < 5; i++ {

		y = qStartY + (minoOffset*3)*float64(i)

		// TODO: Remove. Here for stack debugging
		if i > len(tq.queue) {
			vector.DrawFilledRect(screen, float32(qStartX), float32(y), 50.0, 2.0, color.RGBA{255, 0, 0, 255}, false)
		}

		op.GeoM.Translate(0, y)
		screen.DrawImage(q[i].GetSprite(), op)
		op.GeoM.Translate(0, -y)
	}

}

// Testing function
func newTestQueue() *[]Tetrimino {
	q := make([]Tetrimino, 0)

	for i := 1; i <= 7; i++ {
		q = append(q, NewSPiece())
	}

	return &q
}
