package chess

// TargetTable is a lookup table for squares a piece can target. For example,
// TargetTable[WhitePawn][A2] returns a bitboard with A3, B3, and A4 set.
var TargetTable [][]Bitboard

type magicParameters struct {
	scale uint64
	shift uint8
}

// rookMagicParameters holds parameters for indexing into rookMagicTable.
var rookMagicParameters []magicParameters

// bishopMagicParameters holds parameters for indexing into bishopMagicTable.
var bishopMagicParameters []magicParameters

// rookMagicTable is a lookup table for rook moves that also respects occupied
// squares.
var rookMagicTable []Bitboard

// bishopMagicTable is a lookup table for bishop moves that also respects
// occupied squares.
var bishopMagicTable []Bitboard

func init() {
	// Movement constants for target square generation.
	var (
		knightDeltas = [][]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
		bishopDeltas = [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
		rookDeltas   = [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
		kingDeltas   = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	)

	TargetTable = make([][]Bitboard, 12)
	bitboards := make([]Bitboard, 12*64) // backing slice for TargetTable
	for i := range TargetTable {
		TargetTable[i], bitboards = bitboards[:64], bitboards[64:]
	}

	// Targets for white pieces.
	for s := A1; s <= H8; s++ {
		f, r := s.File(), s.Rank()

		// White pawn targets.
		if r != Rank1 && r != Rank8 {
			TargetTable[WhitePawn][s].Set(NewSquare(f, r+1)) // single push
			if r == Rank2 {
				TargetTable[WhitePawn][s].Set(NewSquare(f, r+2)) // double push
			}
			if f != FileA {
				TargetTable[WhitePawn][s].Set(NewSquare(f-1, r+1)) // left-up
			}
			if f != FileH {
				TargetTable[WhitePawn][s].Set(NewSquare(f+1, r+1)) // right-up
			}
		}

		// White knight targets.
		for _, d := range knightDeltas {
			f := f + File(d[0])
			r := r + Rank(d[1])
			if f.Valid() && r.Valid() {
				TargetTable[WhiteKnight][s].Set(NewSquare(f, r))
			}
		}

		// White bishop targets.
		for _, d := range bishopDeltas {
			f := f + File(d[0])
			r := r + Rank(d[1])
			for f.Valid() && r.Valid() {
				TargetTable[WhiteBishop][s].Set(NewSquare(f, r))
				f = f + File(d[0])
				r = r + Rank(d[1])
			}
		}

		// White rook targets.
		for _, d := range rookDeltas {
			f := f + File(d[0])
			r := r + Rank(d[1])
			for f.Valid() && r.Valid() {
				TargetTable[WhiteRook][s].Set(NewSquare(f, r))
				f += File(d[0])
				r += Rank(d[1])
			}
		}

		// White queen targets.
		TargetTable[WhiteQueen][s] = TargetTable[WhiteBishop][s] | TargetTable[WhiteRook][s]

		// White king targets.
		for _, d := range kingDeltas {
			f := f + File(d[0])
			r := r + Rank(d[1])
			if f.Valid() && r.Valid() {
				TargetTable[WhiteKing][s].Set(NewSquare(f, r))
			}
		}
	}

	// Targets for black pieces.
	copy(TargetTable[BlackPawn], TargetTable[WhitePawn])
	for s := A1; s < H8; s++ {
		TargetTable[BlackPawn][s].Mirror()
	}
	copy(TargetTable[BlackKnight], TargetTable[WhiteKnight])
	copy(TargetTable[BlackBishop], TargetTable[WhiteBishop])
	copy(TargetTable[BlackRook], TargetTable[WhiteRook])
	copy(TargetTable[BlackQueen], TargetTable[WhiteQueen])
	copy(TargetTable[BlackKing], TargetTable[WhiteKing])
}
