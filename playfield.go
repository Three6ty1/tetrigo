package main

type PlayField struct {
	stack [][]bool
}

func (pf PlayField) IsColliding(t Tetrimino) bool {
	return false
}

func (pf PlayField) UpdateStack(t Tetrimino) {

}

func (pf PlayField) ClearLines() {

}

func (pf PlayField) CheckOutOfBounds() bool {
	return false
}
