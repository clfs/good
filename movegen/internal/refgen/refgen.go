// Package refgen is a slow reference implementation of a move generator.
package refgen

import "github.com/clfs/good/chess"

func IsLegalMove(p chess.Position, m chess.Move) bool {
	return false
}

func LegalMoves(p chess.Position) []chess.Move {
	return nil
}
