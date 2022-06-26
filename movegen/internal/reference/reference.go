// Package refgen is a slow reference implementation of a move generator.
package refgen

import "github.com/clfs/good/chess"

// IsLegalMove returns true if a move is legal in a position.
func IsLegalMove(p chess.Position, m chess.Move) bool {
	return false
}

// LegalMoves returns a slice of all legal moves in a position.
func LegalMoves(p chess.Position) []chess.Move {
	return nil
}