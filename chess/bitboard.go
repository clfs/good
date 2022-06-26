package chess

import "strings"

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

// IsEmpty returns true if the bitboard has no bits set.
func (b *Bitboard) IsEmpty() bool {
	return *b == 0
}

// Intersects returns true if two bitboards have any bits in common.
func (b *Bitboard) Intersects(other *Bitboard) bool {
	return *b&*other != 0
}

// SetIf sets the bit at a square to 1 if cond is true.
func (b *Bitboard) SetIf(s Square, cond bool) {
	if cond {
		b.Set(s)
	} else {
		b.Clear(s)
	}
}

// DebugString returns a multi-line string representation of the bitboard.
func (b *Bitboard) DebugString() string {
	var s strings.Builder
	for r := Rank8; r <= Rank8; r-- {
		for f := FileA; f <= FileH; f++ {
			if b.Get(NewSquare(f, r)) {
				s.WriteByte('1')
			} else {
				s.WriteByte('.')
			}
		}
		s.WriteByte('\n')
	}
	return s.String()
}
