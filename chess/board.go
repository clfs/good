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

// FEN returns the FEN representation of the color.
func (c Color) FEN() string {
	return []string{"w", "b"}[c]
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

// FEN returns the FEN representation of the piece.
func (p Piece) FEN() string {
	return []string{
		"p", "n", "b", "r", "q", "k",
		"P", "N", "B", "R", "Q", "K",
	}[p]
}

func (p Piece) String() string {
	return []string{
		"WhitePawn", "WhiteKnight", "WhiteBishop", "WhiteRook", "WhiteQueen", "WhiteKing",
		"BlackPawn", "BlackKnight", "BlackBishop", "BlackRook", "BlackQueen", "BlackKing",
	}[p]
}
