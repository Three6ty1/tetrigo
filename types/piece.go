package types

type Mino int32

// These are equivalents

const (
	green Mino = iota
	red
	blue
	orange
	purple
	yellow
	lightBlue
)

type Piece int32

const (
	SPiece Piece = iota
	ZPiece
	LPiece
	JPiece
	TPiece
	OPiece
	IPiece
)

type Orientation int32

const (
	O0 Orientation = iota
	O90
	O180
	O270
)
