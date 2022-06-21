// Package chess provides basic chess constants and functions.
package chess

// Color represents either white or black.
type Color uint8

const (
	White Color = iota
	Black
)

// Opposite returns the opposite color.
func (c Color) Opposite() Color {
	return c ^ 1
}

// Valid returns true if the color is valid.
func (c Color) Valid() bool {
	return c <= Black
}

func (c Color) String() string {
	return []string{"White", "Black"}[c]
}

// Role represents a piece's role.
type Role uint8

const (
	Pawn Role = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

// Valid returns true if the role is valid.
func (r Role) Valid() bool {
	return r <= King
}

func (r Role) String() string {
	return []string{"Pawn", "Knight", "Bishop", "Rook", "Queen", "King"}[r]
}

// Piece represents a piece.
type Piece uint8

const (
	WhitePawn = iota
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing
	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
)

func NewPiece(c Color, r Role) Piece {
	return Piece(c*BlackPawn + Color(r))
}

// Color returns the color of the piece.
func (p Piece) Color() Color {
	return Color(p / BlackPawn)
}

// Role returns the role of the piece.
func (p Piece) Role() Role {
	return Role(p % BlackPawn)
}

// Valid returns true if the piece is valid.
func (p Piece) Valid() bool {
	return p <= BlackKing
}

func (p Piece) String() string {
	return []string{
		"WhitePawn", "WhiteKnight", "WhiteBishop", "WhiteRook", "WhiteQueen", "WhiteKing",
		"BlackPawn", "BlackKnight", "BlackBishop", "BlackRook", "BlackQueen", "BlackKing",
	}[p]
}

// File represents a board file.
type File uint8

const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

func (f File) Valid() bool {
	return f <= FileH
}

func (f File) String() string {
	return []string{"FileA", "FileB", "FileC", "FileD", "FileE", "FileF", "FileG", "FileH"}[f]
}

// Rank represents a board rank. Note that Rank1 is represented by 0.
type Rank uint8

const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)

func (r Rank) Valid() bool {
	return r <= Rank8
}

func (r Rank) String() string {
	return []string{"Rank1", "Rank2", "Rank3", "Rank4", "Rank5", "Rank6", "Rank7", "Rank8"}[r]
}

// Square represents a board square.
// A1 is represented by 0, B1 by 1, and H8 by 63.
type Square uint8

const (
	A1 Square = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8

	// Special cases. Square.Valid returns false. Not all functions can
	// meaningfully handle these.
	NoEnPassant Square = 255
)

// Valid returns true if the square is valid.
func (s Square) Valid() bool {
	return s <= H8
}

// File returns the file of the square.
func (s Square) File() File {
	return File(s % 8)
}

// Rank returns the rank of the square.
func (s Square) Rank() Rank {
	return Rank(s / 8)
}

// Bitboard returns the bitboard representation of the square.
func (s Square) Bitboard() Bitboard {
	return Bitboard(1 << s)
}

func (s Square) String() string {
	return []string{
		"A1", "B1", "C1", "D1", "E1", "F1", "G1", "H1",
		"A2", "B2", "C2", "D2", "E2", "F2", "G2", "H2",
		"A3", "B3", "C3", "D3", "E3", "F3", "G3", "H3",
		"A4", "B4", "C4", "D4", "E4", "F4", "G4", "H4",
		"A5", "B5", "C5", "D5", "E5", "F5", "G5", "H5",
		"A6", "B6", "C6", "D6", "E6", "F6", "G6", "H6",
		"A7", "B7", "C7", "D7", "E7", "F7", "G7", "H7",
		"A8", "B8", "C8", "D8", "E8", "F8", "G8", "H8",
	}[s]
}

// CastlingFlag represents a single castling right.
type CastlingFlag uint8

const (
	WhiteShortCastle CastlingFlag = 1 << iota
	WhiteLongCastle
	BlackShortCastle
	BlackLongCastle
)

func (c CastlingFlag) String() string {
	return []string{"WhiteShortCastle", "WhiteLongCastle", "BlackShortCastle", "BlackLongCastle"}[c]
}

// CastlingRights represents the castling rights of both players.
type CastlingRights uint8

// Get returns true if a castling right is available.
func (c *CastlingRights) Get(f CastlingFlag) bool {
	return *c&CastlingRights(f) != 0
}

// Disable disables a castling right.
func (c *CastlingRights) Disable(f CastlingFlag) {
	*c &^= CastlingRights(f)
}

// NewCastlingRights returns a new CastlingRights with all rights enabled.
func NewCastlingRights() CastlingRights {
	return CastlingRights(WhiteShortCastle | WhiteLongCastle | BlackShortCastle | BlackLongCastle)
}
