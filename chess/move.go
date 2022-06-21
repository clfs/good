package chess

// Promotion is a flag that indicates promotion information for a Move.
//
// Notably, promotions can be cast to the corresponding piece.
// For example, Piece(PromoteToKnight) == Knight.
type Promotion uint16

const (
	NoPromotion Promotion = iota
	PromoteToKnight
	PromoteToBishop
	PromoteToRook
	PromoteToQueen
)

// Valid returns true if the promotion's underlying integer value is valid.
func (p Promotion) Valid() bool {
	return p <= PromoteToQueen
}

// Move represents an engine move, or equivalently, a transition between
// two positions. In chess terminology, this would be a ply.
type Move uint16

// NewMove returns a new Move.
//
// To represent castling moves, use the king and rook's current squares as the
// from and to squares respectively.
func NewMove(from, to Square, promotion Promotion) Move {
	return Move(from) | Move(to)<<6 | Move(promotion)<<12
}

// From returns the square at which the move starts, except when the move is
// a castling move.
//
// If the move is a castling move, the from square will be the king's current
// position instead.
func (m Move) From() Square {
	return Square(m & 0x3F)
}

// To returns the square at which the move ends, except when the move is a
// castling move.
//
// If the move is a castling move, the to square will be the rook's current
// position instead.
func (m Move) To() Square {
	return Square((m >> 6) & 0x3F)
}

// Promotion returns promotion information for the move.
func (m Move) Promotion() Promotion {
	return Promotion((m >> 12) & 0xF)
}

// IsPromotion returns true if the move is a promotion.
func (m Move) IsPromotion() bool {
	return m.Promotion() != NoPromotion
}
