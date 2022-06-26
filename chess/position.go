package chess

// Position represents a game position.
// Three-fold repetition is not tracked here.
type Position struct {
	Board          []Bitboard // Board describes which pieces are on the board and where.
	CastleRights   CastleRights
	EnPassantRight EnPassantRight
	SideToMove     Color
	HalfMoves      uint8  // HalfMoves is the number of half moves since the last capture or pawn move.
	FullMoves      uint16 // FullMoves starts at 1 and is incremented after Black moves.
}

// NewPosition returns a new position, pre-populated with all starting pieces.
func NewPosition() Position {
	p := Position{
		Board:          make([]Bitboard, 12),
		CastleRights:   AllCastleRights,
		EnPassantRight: NoEnPassantRight,
		SideToMove:     White,
		HalfMoves:      0,
		FullMoves:      1,
	}
	p.Put(WhiteRook, A1)
	p.Put(WhiteKnight, B1)
	p.Put(WhiteBishop, C1)
	p.Put(WhiteQueen, D1)
	p.Put(WhiteKing, E1)
	p.Put(WhiteBishop, F1)
	p.Put(WhiteKnight, G1)
	p.Put(WhiteRook, H1)
	for sq := A2; sq <= H2; sq++ {
		p.Put(WhitePawn, sq)
	}
	for sq := A7; sq <= H7; sq++ {
		p.Put(BlackPawn, sq)
	}
	p.Put(BlackRook, A8)
	p.Put(BlackKnight, B8)
	p.Put(BlackBishop, C8)
	p.Put(BlackQueen, D8)
	p.Put(BlackKing, E8)
	p.Put(BlackBishop, F8)
	p.Put(BlackKnight, G8)
	p.Put(BlackRook, H8)
	return p
}

// Put puts a piece on the board. No other fields are updated.
func (p *Position) Put(pc Piece, s Square) {
	p.Board[pc].Set(s)
}

// Get returns the piece on the given square.
// If there's no piece there, ok is false.
func (p *Position) Get(s Square) (pc Piece, ok bool) {
	for pc := WhitePawn; pc <= BlackKing; pc++ {
		if p.Board[pc].Get(s) {
			return pc, true
		}
	}
	return 0, false
}

// Reset resets the position to the starting position.
func (p *Position) Reset() {
	// TODO: Optimize this to eliminate all allocations.
	*p = NewPosition()
}
