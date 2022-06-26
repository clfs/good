package movegen

import (
	"github.com/clfs/good/chess"
	"github.com/clfs/good/movegen/internal/refgen"
)

// IsLegalMove returns true if a move is legal in a position.
func IsLegalMove(p chess.Position, m chess.Move) bool {
	return refgen.IsLegalMove(p, m)
}

// LegalMoves returns a slice of all legal moves in a position.
func LegalMoves(p chess.Position) []chess.Move {
	return refgen.LegalMoves(p)
}
