// Package refeval implements an evaluation function that only accounts for
// material and checkmate.
package refeval

import (
	"github.com/clfs/good/chess"
)

// PieceValues contains the values of pieces, measured in centipawns.
var PieceValues = map[chess.Piece]int{
	chess.WhitePawn:   100,
	chess.WhiteKnight: 300,
	chess.WhiteBishop: 300,
	chess.WhiteRook:   500,
	chess.WhiteQueen:  900,
	chess.BlackPawn:   -100,
	chess.BlackKnight: -300,
	chess.BlackBishop: -300,
	chess.BlackRook:   -500,
	chess.BlackQueen:  -900,
}

// Position returns the value of a position.
func Position(p chess.Position) int {
	// TODO: Account for checkmate.
	var score int
	for s := chess.A1; s <= chess.H8; s++ {
		pc, ok := p.Get(s)
		if !ok {
			continue
		}
		score += PieceValues[pc]
	}
	return score
}
