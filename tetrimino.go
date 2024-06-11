package main

import "github.com/hajimehoshi/ebiten/v2"

type Vector struct {
	X float64
	Y float64
}

type Orientation int32

const (
	Original Orientation = iota
	Right
	OneHundredEighty
	Left
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
	locked      bool
	position    Vector
	orientation Orientation
	sprite      *ebiten.Image
}

func (t Tetrimino) RotateLeft(pf PlayField) {

}
