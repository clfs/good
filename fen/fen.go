// Package fen implements parsing and generation for FEN notation.
package fen

import "github.com/clfs/good/chess"

// Starting is the FEN for the starting position.
const Starting = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

// To returns the FEN for a position.
func To(p chess.Position) string {
	return ""
}

// From returns the position for a FEN string.
func From(s string) (*chess.Position, error) {
	p := chess.NewPosition()
	return p, nil
}
