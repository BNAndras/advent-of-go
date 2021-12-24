package main

import "testing"

func TestPartOne(t *testing.T) {
	inp := []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678"}
	got := solvePartOne(inp)
	want := 15
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
