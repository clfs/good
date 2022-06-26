package fastgen

import "github.com/clfs/good/chess"

// Table is a collection of move lookup tables.
var Table [12]map[chess.Square][]chess.Move

// Lookup returns a slice of all possible moves for the given piece and square.
// Moves are not guaranteed to be legal or in any particular order.
func Lookup(p chess.Piece, s chess.Square) []chess.Move {
	return Table[p][s]
}

func init() {
	// Instantiation.
	Table = [12]map[chess.Square][]chess.Move{}
	for i := 0; i < 12; i++ {
		Table[i] = make(map[chess.Square][]chess.Move)
	}

	// [DONE] White pawns moving one square up without promoting.
	for s := chess.A2; s <= chess.H6; s++ {
		Table[chess.WhitePawn][s] = append(Table[chess.WhitePawn][s], chess.NewMove(s, s.Up()))
	}

	// [DONE] White pawns moving two squares up.
	for s := chess.A2; s <= chess.H2; s++ {
		Table[chess.WhitePawn][s] = append(Table[chess.WhitePawn][s], chess.NewMove(s, s.UpN(2)))
	}

	// [DONE] White pawns capturing without promoting.
	for s := chess.A2; s <= chess.H6; s++ {
		if !s.IsLeftEdge() {
			Table[chess.WhitePawn][s] = append(Table[chess.WhitePawn][s], chess.NewMove(s, s.Up().Left()))
		}
		if !s.IsRightEdge() {
			Table[chess.WhitePawn][s] = append(Table[chess.WhitePawn][s], chess.NewMove(s, s.Up().Right()))
		}
	}

	// [DONE] White pawns promoting without capturing.
	for s := chess.A7; s <= chess.H7; s++ {
		for pc := chess.WhiteKnight; pc <= chess.WhiteQueen; pc++ {
			Table[chess.WhitePawn][s] = append(Table[chess.WhitePawn][s], chess.NewPromotionMove(s, s.Up(), pc))
		}
	}

	// [DONE] White pawns both capturing and promoting.
	for s := chess.A7; s <= chess.H7; s++ {
		for pc := chess.WhiteKnight; pc <= chess.WhiteQueen; pc++ {
			if !s.IsLeftEdge() {
				Table[chess.WhitePawn][s] = append(Table[chess.WhitePawn][s], chess.NewPromotionMove(s, s.Up().Left(), pc))
			}
			if !s.IsRightEdge() {
				Table[chess.WhitePawn][s] = append(Table[chess.WhitePawn][s], chess.NewPromotionMove(s, s.Up().Right(), pc))
			}
		}
	}

	// [DONE] Black pawns moving one square down without promoting.
	for s := chess.A3; s <= chess.H7; s++ {
		Table[chess.BlackPawn][s] = append(Table[chess.BlackPawn][s], chess.NewMove(s, s.Down()))
	}

	// [DONE] Black pawns moving two squares down.
	for s := chess.A7; s <= chess.H7; s++ {
		Table[chess.BlackPawn][s] = append(Table[chess.BlackPawn][s], chess.NewMove(s, s.DownN(2)))
	}

	// Black pawns capturing without promoting.

	// Black pawns promoting without capturing.

	// Black pawns both capturing and promoting.

	// White knight moves.

	// White bishop moves.

	// White rook moves.

	// White queen moves.

	// White king moves, not including castling.

	// White castling moves.

	// [DONE] Black knight, bishop, rook, and queen moves.
	Table[chess.BlackKnight] = Table[chess.WhiteKnight]
	Table[chess.BlackBishop] = Table[chess.WhiteBishop]
	Table[chess.BlackRook] = Table[chess.WhiteRook]
	Table[chess.BlackQueen] = Table[chess.WhiteQueen]

	// Black king moves, not including castling.

	// Black castling moves.
}
