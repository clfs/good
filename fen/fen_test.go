package fen

import (
	"testing"

	"github.com/clfs/good/chess"
	"github.com/google/go-cmp/cmp"
)

func TestFrom_Starting(t *testing.T) {
	p1 := chess.NewPosition()
	p2, err := From(Starting)
	if err != nil {
		t.Errorf("failed to parse starting position: %v", err)
	}
	if diff := cmp.Diff(p1, p2); diff != "" {
		t.Errorf("(-want +got)\n%s", diff)
	}
}

func FuzzFrom_NoPanic(f *testing.F) {
	f.Add(Starting)
	f.Fuzz(func(t *testing.T, s string) {
		_, _ = From(s)
	})
}
