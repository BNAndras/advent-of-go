package main

import "testing"

func TestPartOne(t *testing.T) {
	inp := []string{
		"3,4,3,1,2",
	}
	got := solvePartOne(inp)
	want := 5934
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	inp := []string{
		"3,4,3,1,2",
	}
	got := solvePartTwo(inp)
	want := 26984457539
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
