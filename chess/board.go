package chess

// Color represents either white or black.
type Color uint8

const (
	White Color = iota
	Black
)

// Opposite returns the opposite color.
func (c Color) Opposite() Color {
	return c ^ 1
}

// Valid returns true if the color is valid.
func (c Color) Valid() bool {
	return c <= Black
}

// FEN returns the FEN representation of the color.
func (c Color) FEN() string {
	return []string{"w", "b"}[c]
}

// String returns the string representation of the color.
func (c Color) String() string {
	return []string{"White", "Black"}[c]
}
