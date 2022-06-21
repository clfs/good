package chess

// Bitboard is an integer where each bit represents one square.
// From LSB to MSB, the bits represent a1, b1, ..., h8.
type Bitboard uint64

// Set sets the given square to 1.
func (b *Bitboard) Set(s Square) {
	*b |= s.Bitboard()
}

// Clear sets the given square to 0.
func (b *Bitboard) Clear(s Square) {
	*b &^= s.Bitboard()
}

// Get returns true if the given square is set.
func (b *Bitboard) Get(s Square) bool {
	return *b&s.Bitboard()>>s != 0
}

// Toggle toggles the bit at the given square.
func (b *Bitboard) Toggle(s Square) {
	*b ^= s.Bitboard()
}

// Assign assigns the provided value at a given square.
func (b *Bitboard) Assign(s Square, v bool) {
	if v {
		b.Set(s)
	} else {
		b.Clear(s)
	}
}
