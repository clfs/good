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
	WhitePawn Piece = iota
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
	return Piece(uint8(c)*6 + uint8(r))
}

// Color returns the color of the piece.
func (p Piece) Color() Color {
	return Color(p / 6)
}

// Role returns the role of the piece.
func (p Piece) Role() Role {
	return Role(p % 6)
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

// Left returns the file to the left. For FileA, the result is invalid.
func (f File) Left() File {
	return f - 1
}

// Right returns the file to the right. For FileH, the result is invalid.
func (f File) Right() File {
	return f + 1
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

// Above returns the rank above. For Rank8, the result is invalid.
func (r Rank) Above() Rank {
	return r + 1
}

// Below returns the rank below. For Rank1, the result is invalid.
func (r Rank) Below() Rank {
	return r - 1
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
)

// NewSquare returns a new Square at the given file and rank.
func NewSquare(f File, r Rank) Square {
	return Square(f) | Square(r)<<3
}

// IsValidCoords returns true if the given coordinates reference a valid square.
func IsValidCoords(f File, r Rank) bool {
	return f.Valid() && r.Valid()
}

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

// CastleRight represents a single castle right.
type CastleRight uint8

const (
	WhiteShortCastleRight CastleRight = 1 << iota
	WhiteLongCastleRight
	BlackShortCastleRight
	BlackLongCastleRight
)

func (c CastleRight) String() string {
	return []string{"WhiteShortCastleRight", "WhiteLongCastleRight", "BlackShortCastleRight", "BlackLongCastleRight"}[c]
}

// CastleRights represents the available castle rights of both players.
type CastleRights uint8

const (
	// NoCastleRights represents the state where no castle rights are available.
	NoCastleRights CastleRights = 0
	// AllCastleRights represents the state where all castle rights are available to both players.
	AllCastleRights CastleRights = 0xF
)

// Get returns true if a castle right is available.
func (c *CastleRights) Get(r CastleRight) bool {
	return *c&CastleRights(r) != 0
}

// Enable enables a castle right.
func (c *CastleRights) Enable(r CastleRight) {
	*c |= CastleRights(r)
}

// Disable disables a castle right.
func (c *CastleRights) Disable(r CastleRight) {
	*c &^= CastleRights(r)
}

// Move represents an engine move, or equivalently, a transition between
// two positions. In chess terminology, this would be a ply.
type Move uint16

// NewMove returns a new Move.
//
// To represent promotion moves, use NewPromotionMove instead.
//
// To represent castling moves, use the king's original and destination squares
// as the from and to squares respectively.
func NewMove(from, to Square) Move {
	return Move(from) | Move(to)<<6
}

// NewPromotionMove returns a new Move that records the given promotion.
func NewPromotionMove(from, to Square, p Piece) Move {
	return Move(from) | Move(to)<<6 | Move(p)<<12
}

// From returns the square at which the move starts.
//
// If the move is castling, the from square is the king's original square.
func (m Move) From() Square {
	return Square(m & 0x3F)
}

// To returns the square at which the move ends.
//
// If the move is castling, the to square is the king's destination square.
func (m Move) To() Square {
	return Square((m >> 6) & 0x3F)
}

// Promotion returns promotion information for the move. If the move is not
// a promotion move, ok is false.
func (m Move) Promotion() (p Piece, ok bool) {
	p = Piece(m >> 12)
	return p, p == 0
}

// EnPassantRight represents an en passant right.
//
// A Square can be cast directly to this type, like EnPassantRight(D6).
//
// To represent the lack of an en passant right, use NoEnPasssantRight.
type EnPassantRight uint8

// NoEnPassantRight represents the lack of an en passant right.
const NoEnPassantRight EnPassantRight = 0xFF
