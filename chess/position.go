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

// AllPieces returns a bitboard of all piece locations.
func (p *Position) AllPieces() Bitboard {
	var b Bitboard
	for _, bb := range p.Board {
		b |= bb
	}
	return b
}

// WhitePieces returns a bitboard of all white piece locations.
func (p *Position) WhitePieces() Bitboard {
	var b Bitboard
	for pc := WhitePawn; pc <= WhiteKing; pc++ {
		b |= p.Board[pc]
	}
	return b
}

// BlackPieces returns a bitboard of all black piece locations.
func (p *Position) BlackPieces() Bitboard {
	var b Bitboard
	for pc := BlackPawn; pc <= BlackKing; pc++ {
		b |= p.Board[pc]
	}
	return b
}

/*
// LegalMoves returns all legal moves in the position.
func (p *Position) LegalMoves() []Move {
	var (
		result      []Move
		friends     []Bitboard // Bitboards for all friendly pieces.
		enemies     []Bitboard // Bitboards for all enemy pieces.
		shortCastle bool       // Whether or not the side to move can castle short.
		longCastle  bool       // Whether or not the side to move can castle long.
	)

	if p.SideToMove == White {
		friends = p.Board[WhitePawn : WhiteKing+1]
		enemies = p.Board[BlackPawn : BlackKing+1]
		shortCastle = p.CastleRights.Get(WhiteShortCastleRight)
		longCastle = p.CastleRights.Get(WhiteLongCastleRight)
	} else {
		friends = p.Board[BlackPawn : BlackKing+1]
		enemies = p.Board[WhitePawn : WhiteKing+1]
		shortCastle = p.CastleRights.Get(BlackShortCastleRight)
		longCastle = p.CastleRights.Get(BlackLongCastleRight)
	}

	var friendOccupancy Bitboard // Bitboard of all friendly pieces' locations.
	for _, bb := range friends {
		friendOccupancy |= bb
	}

	for s := A1; s <= H8; s++ {
		// Skip if there's no piece on the square.
		if !friendOccupancy.Get(s) {
			continue
		}

		// If it's a king...
		if friends[King].Get(s) {
			// ... add legal non-castling moves.
			continue
		}
	}

	return result
}
*/
