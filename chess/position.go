package chess

// Position represents a game position.
// Three-fold repetition is not tracked here.
type Position struct {
	Board      [12]Bitboard
	Castling   CastlingRights
	EnPassant  Square
	SideToMove Color
	HalfMoves  uint8
	FullMoves  uint16
}

func NewPosition() *Position {
	return &Position{
		Board:      [12]Bitboard{},
		Castling:   NewCastlingRights(),
		EnPassant:  NoEnPassant,
		SideToMove: White,
		HalfMoves:  0,
		FullMoves:  0,
	}
}

func (p *Position) Put(piece Piece, square Square) {
	p.Board[piece] |= Bitboard(1) << square
}
