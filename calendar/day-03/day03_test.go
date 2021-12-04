package main

import "testing"

func TestPartOne(t *testing.T) {
	inp := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010"}
	got := solvePart1(inp)
	var want int64 = 198
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
