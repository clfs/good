package eval

import (
	"github.com/clfs/good/chess"
	"github.com/clfs/good/eval/internal/refeval"
)

// Position returns the value of a position. Positive values are good for white,
// and negative values are good for black.
func Position(p chess.Position) int {
	return refeval.Position(p)
}
