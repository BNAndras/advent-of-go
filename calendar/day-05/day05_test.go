package main

import "testing"

func TestPartOne(t *testing.T) {
	inp := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2"}
	got := solvePart1(inp)
	want := 5
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	inp := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2"}
	got := solvePart2(inp)
	want := 12
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
