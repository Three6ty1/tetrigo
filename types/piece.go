package types

type Piece int32

const (
	Empty Piece = iota
	SPiece
	ZPiece
	LPiece
	JPiece
	TPiece
	OPiece
	IPiece
)

type Mino int32

// These two are equivalents
const (
	None Mino = iota
	Green
	Red
	Orange
	Blue
	Purple
	Yellow
	LightBlue
)

type Orientation int32

const (
	O0 Orientation = iota
	O90
	O180
	O270
)
