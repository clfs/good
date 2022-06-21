package chess

import "testing"

func TestColor_Opposite(t *testing.T) {
	if White.Opposite() != Black {
		t.Error("White.Opposite() != Black")
	}
	if Black.Opposite() != White {
		t.Error("Black.Opposite() != White")
	}
}
