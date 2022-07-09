package chess

// targetTable is a lookup table for squares a piece can target. For example,
// targetTable[WhitePawn][A2] returns a bitboard with A3, B3, and A4 set.
var targetTable [][]Bitboard

// Targets returns the squares a piece can target from a given square. For
// example, Targets(WhitePawn, A2) returns a bitboard with A3, B3, and A4 set.
func Targets(p Piece, s Square) Bitboard {
	return targetTable[p][s]
}

// magicTable is a lookup table for squares that sliding pieces (bishops, rooks,
// and queens) can target, while also respecting occupied squares.
var magicTable []Bitboard

type magicParameters struct {
	index uint64 // index into magicTable
	mask  uint64
	scale uint64
	shift uint8
}

// rookMagicParameters holds parameters for indexing into rookMagicTable.
var rookMagicParameters []magicParameters

// bishopMagicParameters holds parameters for indexing into bishopMagicTable.
var bishopMagicParameters []magicParameters

// BishopAttacks returns the squares a bishop can attack from a given square,
// accounting for occupied squares.
func BishopAttacks(s Square, occupied Bitboard) Bitboard {
	p := bishopMagicParameters[s]
	offset := uint64(occupied) & p.mask * p.scale >> p.shift
	return magicTable[p.index+offset]
}

// RookAttacks returns the squares a rook can attack from a given square,
// accounting for occupied squares.
func RookAttacks(s Square, occupied Bitboard) Bitboard {
	p := rookMagicParameters[s]
	offset := uint64(occupied) & p.mask * p.scale >> p.shift
	return magicTable[p.index+offset]
}

// QueenAttacks returns the squares a queen can attack from a given square,
// accounting for occupied squares.
func QueenAttacks(s Square, occupied Bitboard) Bitboard {
	return BishopAttacks(s, occupied) | RookAttacks(s, occupied)
}

func init() {
	// Movement constants for target square generation.
	var (
		knightDeltas = [][]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
		bishopDeltas = [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
		rookDeltas   = [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
		kingDeltas   = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	)

	targetTable = make([][]Bitboard, 12)
	bitboards := make([]Bitboard, 12*64) // Backing slice for targetTable
	for i := range targetTable {
		targetTable[i], bitboards = bitboards[:64], bitboards[64:]
	}

	// Targets for white pieces.
	for s := A1; s <= H8; s++ {
		f, r := s.File(), s.Rank()

		// White pawn targets.
		if r != Rank1 && r != Rank8 {
			targetTable[WhitePawn][s].Set(NewSquare(f, r+1)) // single push
			if r == Rank2 {
				targetTable[WhitePawn][s].Set(NewSquare(f, r+2)) // double push
			}
			if f != FileA {
				targetTable[WhitePawn][s].Set(NewSquare(f-1, r+1)) // left-up
			}
			if f != FileH {
				targetTable[WhitePawn][s].Set(NewSquare(f+1, r+1)) // right-up
			}
		}

		// White knight targets.
		for _, d := range knightDeltas {
			f := f + File(d[0])
			r := r + Rank(d[1])
			if f.Valid() && r.Valid() {
				targetTable[WhiteKnight][s].Set(NewSquare(f, r))
			}
		}

		// White bishop targets.
		for _, d := range bishopDeltas {
			f := f + File(d[0])
			r := r + Rank(d[1])
			for f.Valid() && r.Valid() {
				targetTable[WhiteBishop][s].Set(NewSquare(f, r))
				f = f + File(d[0])
				r = r + Rank(d[1])
			}
		}

		// White rook targets.
		for _, d := range rookDeltas {
			f := f + File(d[0])
			r := r + Rank(d[1])
			for f.Valid() && r.Valid() {
				targetTable[WhiteRook][s].Set(NewSquare(f, r))
				f += File(d[0])
				r += Rank(d[1])
			}
		}

		// White queen targets.
		targetTable[WhiteQueen][s] = targetTable[WhiteBishop][s] | targetTable[WhiteRook][s]

		// White king targets.
		for _, d := range kingDeltas {
			f := f + File(d[0])
			r := r + Rank(d[1])
			if f.Valid() && r.Valid() {
				targetTable[WhiteKing][s].Set(NewSquare(f, r))
			}
		}
	}

	// Targets for black pieces.
	copy(targetTable[BlackPawn], targetTable[WhitePawn])
	for s := A1; s < H8; s++ {
		targetTable[BlackPawn][s].Mirror()
	}
	copy(targetTable[BlackKnight], targetTable[WhiteKnight])
	copy(targetTable[BlackBishop], targetTable[WhiteBishop])
	copy(targetTable[BlackRook], targetTable[WhiteRook])
	copy(targetTable[BlackQueen], targetTable[WhiteQueen])
	copy(targetTable[BlackKing], targetTable[WhiteKing])
}
