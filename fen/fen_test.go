package fen

import "testing"

func FuzzFrom_NoPanic(f *testing.F) {
	f.Add(Starting)
	f.Fuzz(func(t *testing.T, s string) {
		_, _ = From(s)
	})
}
