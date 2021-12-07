package main

import "testing"

func TestPartOne(t *testing.T) {
	inp := []string{
		"16,1,2,0,4,2,7,1,2,14",
	}
	got := solvePartOne(inp)
	want := 37
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
