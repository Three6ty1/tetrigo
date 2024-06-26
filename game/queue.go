package game

import (
	"math/rand"
	"time"

	"github.com/Three6ty1/tetrigo/game/tetrimino"
	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// var r = rand.New(rand.NewSource(2))

var TEST = false // Enable testing variables

// Queue is 6 long, but only displays 5
type TetriminoQueue struct {
	queue []tetrimino.Tetrimino
	next  []tetrimino.Tetrimino
}

func newQueue() *[]tetrimino.Tetrimino {
	// https://tetris.fandom.com/wiki/Random_Generator
	// 7-piece bag
	q := make([]tetrimino.Tetrimino, 0)

	for i := 1; i <= 7; i++ {
		q = append(q, tetrimino.NewTetrimino(types.Piece(i)))
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

func (tq *TetriminoQueue) Next() tetrimino.Tetrimino {
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
	var q []tetrimino.Tetrimino

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

		op.GeoM.Translate(0, y)
		screen.DrawImage(q[i].GetAltSprite(), op)
		op.GeoM.Translate(0, -y)
	}

}

// Testing function
func newTestQueue() *[]tetrimino.Tetrimino {
	q := make([]tetrimino.Tetrimino, 0)

	for i := 1; i <= 7; i++ {
		q = append(q, tetrimino.NewSPiece())
	}

	return &q
}
