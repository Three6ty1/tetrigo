package main

import "github.com/hajimehoshi/ebiten/v2"

type Vector struct {
	X float64
	Y float64
}

type Orientation int32

const (
	O0 Orientation = iota
	O90
	O180
	O270
)

type Piece int32

const (
	S Piece = iota
	Z
	L
	J
	T
	O
	I
)

type Tetrimino struct {
	piece       Piece
	position    Vector
	orientation Orientation
	sprite      *ebiten.Image
}

func (t Tetrimino) Draw(screen *ebiten.Image) {

}

func (t Tetrimino) RotateLeft(pf PlayField) {

}

func (t Tetrimino) RotateRight(pf PlayField) {

}

func (t Tetrimino) GetPosition() {

}

func (t Tetrimino) SetPosition() {

}
