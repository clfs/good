// Package fen implements parsing and generation for FEN notation.
package fen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/clfs/good/chess"
)

// Starting is the FEN for the starting position.
const Starting = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func init() {
	colorTo = make(map[chess.Color]string)
	pieceTo = make(map[chess.Piece]rune)
	castleRightsTo = make(map[chess.CastleRights]string)
	enPassantRightTo = make(map[chess.EnPassantRight]string)
	for k, v := range colorFrom {
		colorTo[v] = k
	}
	for k, v := range pieceFrom {
		pieceTo[v] = k
	}
	for k, v := range castleRightsFrom {
		castleRightsTo[v] = k
	}
	for k, v := range enPassantRightFrom {
		enPassantRightTo[v] = k
	}
}

var (
	colorTo          map[chess.Color]string
	pieceTo          map[chess.Piece]rune
	castleRightsTo   map[chess.CastleRights]string
	enPassantRightTo map[chess.EnPassantRight]string
)

var colorFrom = map[string]chess.Color{
	"w": chess.White,
	"b": chess.Black,
}

var pieceFrom = map[rune]chess.Piece{
	'P': chess.WhitePawn,
	'N': chess.WhiteKnight,
	'B': chess.WhiteBishop,
	'R': chess.WhiteRook,
	'Q': chess.WhiteQueen,
	'K': chess.WhiteKing,
	'p': chess.BlackPawn,
	'n': chess.BlackKnight,
	'b': chess.BlackBishop,
	'r': chess.BlackRook,
	'q': chess.BlackQueen,
	'k': chess.BlackKing,
}

var castleRightsFrom = map[string]chess.CastleRights{
	"-":    chess.NoCastleRights,
	"K":    chess.CastleRights(chess.WhiteShortCastleRight),
	"Q":    chess.CastleRights(chess.WhiteLongCastleRight),
	"k":    chess.CastleRights(chess.BlackShortCastleRight),
	"q":    chess.CastleRights(chess.BlackLongCastleRight),
	"KQ":   chess.CastleRights(chess.WhiteShortCastleRight | chess.WhiteLongCastleRight),
	"Kk":   chess.CastleRights(chess.WhiteShortCastleRight | chess.BlackShortCastleRight),
	"Kq":   chess.CastleRights(chess.WhiteShortCastleRight | chess.BlackLongCastleRight),
	"Qk":   chess.CastleRights(chess.WhiteLongCastleRight | chess.BlackShortCastleRight),
	"Qq":   chess.CastleRights(chess.WhiteLongCastleRight | chess.BlackLongCastleRight),
	"kq":   chess.CastleRights(chess.BlackShortCastleRight | chess.BlackLongCastleRight),
	"KQk":  chess.CastleRights(chess.WhiteShortCastleRight | chess.WhiteLongCastleRight | chess.BlackShortCastleRight),
	"KQq":  chess.CastleRights(chess.WhiteShortCastleRight | chess.WhiteLongCastleRight | chess.BlackLongCastleRight),
	"Kkq":  chess.CastleRights(chess.WhiteShortCastleRight | chess.BlackShortCastleRight | chess.BlackLongCastleRight),
	"Qkq":  chess.CastleRights(chess.WhiteLongCastleRight | chess.BlackShortCastleRight | chess.BlackLongCastleRight),
	"KQkq": chess.AllCastleRights,
}

var enPassantRightFrom = map[string]chess.EnPassantRight{
	"-":  chess.NoEnPassantRight,
	"a3": chess.EnPassantRight(chess.A3),
	"b3": chess.EnPassantRight(chess.B3),
	"c3": chess.EnPassantRight(chess.C3),
	"d3": chess.EnPassantRight(chess.D3),
	"e3": chess.EnPassantRight(chess.E3),
	"f3": chess.EnPassantRight(chess.F3),
	"g3": chess.EnPassantRight(chess.G3),
	"h3": chess.EnPassantRight(chess.H3),
	"a6": chess.EnPassantRight(chess.A6),
	"b6": chess.EnPassantRight(chess.B6),
	"c6": chess.EnPassantRight(chess.C6),
	"d6": chess.EnPassantRight(chess.D6),
	"e6": chess.EnPassantRight(chess.E6),
	"f6": chess.EnPassantRight(chess.F6),
	"g6": chess.EnPassantRight(chess.G6),
	"h6": chess.EnPassantRight(chess.H6),
}

// To returns the FEN for a position.
func To(p chess.Position) string {
	var b strings.Builder

	// Piece placement.
	for r := chess.Rank8; r <= chess.Rank8; r-- {
		skip := 0
		for f := chess.FileA; f <= chess.FileH; f++ {
			sq := chess.NewSquare(f, r)
			piece, ok := p.Get(sq)
			if !ok {
				skip++
				continue
			}
			if skip > 0 {
				fmt.Fprintf(&b, "%d", skip)
				skip = 0
			}
			fmt.Fprintf(&b, "%c", pieceTo[piece]) // lookup is guaranteed ok
		}
		if skip > 0 {
			fmt.Fprintf(&b, "%d", skip)
		}
		if r != chess.Rank1 {
			fmt.Fprintf(&b, "/")
		}
	}

	// Active color.
	fmt.Fprintf(&b, " %s", colorTo[p.SideToMove])

	// Castling rights.
	fmt.Fprintf(&b, " %s", castleRightsTo[p.CastleRights])

	// En passant target square.
	fmt.Fprintf(&b, " %s", enPassantRightTo[p.EnPassantRight])

	// Half-move clock.
	fmt.Fprintf(&b, " %d", p.HalfMoves)

	// Full-move count.
	fmt.Fprintf(&b, " %d", p.FullMoves)

	return b.String()
}

// BUG(clfs): From accepts invalid piece placements without returning an error.

// From returns the position described by the FEN string.
//
// These are the only deviations from the PGN standard:
//
//   - Adjacent fields must be separated by one or more consecutive white space
//     characters, as defined by unicode.IsSpace.
//   - The en passant target square, if any, must be on the third or sixth rank.
//   - If the full move number is 0, it is interpreted as if it were 1.
func From(s string) (chess.Position, error) {
	var p chess.Position

	fields := strings.Fields(s)
	if l := len(fields); l != 6 {
		return p, fmt.Errorf("fen: invalid number of fields: %d", l)
	}

	// Piece placement.
	// TODO: check that all 64 squares are specified.
	square := chess.A8
	for _, r := range fields[0] {
		switch r {
		case '1', '2', '3', '4', '5', '6', '7', '8':
			square += chess.Square(r - '0') // advance rightwards
		case '/':
			square -= 16 // move to the leftmost square in the rank below
		default:
			piece, ok := pieceFrom[r]
			if !ok {
				return p, fmt.Errorf("fen: invalid board rune: %c", r)
			}
			p.Put(piece, square)
			square++
		}
	}

	// Active color.
	color, ok := colorFrom[fields[1]]
	if !ok {
		return p, fmt.Errorf("fen: invalid side to move: %s", fields[1])
	}
	p.SideToMove = color

	// Castling rights.
	castleRights, ok := castleRightsFrom[fields[2]]
	if !ok {
		return p, fmt.Errorf("fen: invalid castle rights: %s", fields[2])
	}
	p.CastleRights = castleRights

	// En passant square.
	enPassantRight, ok := enPassantRightFrom[fields[3]]
	if !ok {
		return p, fmt.Errorf("fen: invalid en passant square: %s", fields[3])
	}
	p.EnPassantRight = enPassantRight

	// Half-move clock.
	halfMoves, err := strconv.ParseUint(fields[4], 10, 8)
	if err != nil {
		return p, fmt.Errorf("fen: invalid half-move clock: %s", fields[4])
	}
	p.HalfMoves = uint8(halfMoves)

	// Full-move count.
	fullMoves, err := strconv.ParseUint(fields[5], 10, 16)
	if err != nil {
		return p, fmt.Errorf("fen: invalid full-move count: %s", fields[5])
	}
	if fullMoves == 0 { // Fix a common mistake in various FEN strings.
		fullMoves = 1
	}
	p.FullMoves = uint16(fullMoves)

	return p, nil
}
