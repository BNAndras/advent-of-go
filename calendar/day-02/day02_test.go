package main

import "testing"

func TestSolvePart1(t *testing.T) {
	inp := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2"}
	got := solvePart1(inp)
	want := 150
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSolvePart2(t *testing.T) {
	inp := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2"}
	got := solvePart2(inp)
	want := 900
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
