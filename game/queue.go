package game

import (
	"math/rand"
	"time"

	"github.com/Three6ty1/tetrigo/types"
	"github.com/hajimehoshi/ebiten/v2"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// var r = rand.New(rand.NewSource(2))

// Queue is 6 long, but only displays 5
type TetriminoQueue struct {
	queue []Tetrimino
}

func NewTetriminoQueue() *TetriminoQueue {
	q := make([]Tetrimino, 0)

	// [0 - 7)
	for len(q) < 6 {
		q = append(q, NewTetrimino(types.Piece(r.Int31n(6)+1)))
	}

	tq := &TetriminoQueue{
		queue: q,
	}

	return tq
}

func (tq *TetriminoQueue) Next() Tetrimino {
	// https://stackoverflow.com/a/26863706
	next := tq.queue[0]
	tq.queue = append(tq.queue, NewTetrimino(types.Piece(r.Int31n(6)+1)))
	tq.queue = tq.queue[1:]

	return next
}

func (tq TetriminoQueue) Draw(screen *ebiten.Image, pfStart types.Vector, minoOffset float64, gameScale float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(gameScale*0.75, gameScale*0.75)

	// Move to top right corner of the playfield and then some more
	qStartX := pfStart.X + (minoOffset * 13)
	qStartY := pfStart.Y

	op.GeoM.Translate(qStartX, qStartY)

	var y float64

	for i := 0; i < len(tq.queue); i++ {
		y = qStartY + (minoOffset*4)*float64(i)

		op.GeoM.Translate(0, y)
		screen.DrawImage(tq.queue[i].GetSprite(), op)
		op.GeoM.Translate(0, -y)
	}

}
