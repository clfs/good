package chess

// Bitboard is an integer where each bit represents one square.
// From LSB to MSB, the bits represent a1, b1, ..., h1, a2, ..., h8.
type Bitboard uint64

// Set sets the bit at a square to 1.
func (b *Bitboard) Set(s Square) {
	*b |= s.Bitboard()
}

// Clear clears the bit at a square to 0.
func (b *Bitboard) Clear(s Square) {
	*b &^= s.Bitboard()
}

// Get returns true if the bit at a square is 1.
func (b *Bitboard) Get(s Square) bool {
	return *b&s.Bitboard()>>s != 0
}

// Toggle toggles the bit at a square.
func (b *Bitboard) Toggle(s Square) {
	*b ^= s.Bitboard()
}

// SetIf sets the bit at a square to 1 if cond is true.
func (b *Bitboard) SetIf(s Square, cond bool) {
	if cond {
		b.Set(s)
	} else {
		b.Clear(s)
	}
}
