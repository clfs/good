package fen

import (
	"testing"

	"github.com/clfs/good/chess"
	"github.com/google/go-cmp/cmp"
)

func TestTo_Starting(t *testing.T) {
	if s := To(chess.NewPosition()); s != Starting {
		t.Errorf("want %s, got %s", Starting, s)
	}
}

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
