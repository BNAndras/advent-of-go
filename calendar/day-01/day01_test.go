package main

import "testing"

func TestSolvePart1(t *testing.T) {
	inp := []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}
	got := solvePart1(inp)
	want := 7
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	inp := []string{"607", "618", "618", "617", "647", "716", "769", "792"}
	got := solvePart2(inp)
	want := 5
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
