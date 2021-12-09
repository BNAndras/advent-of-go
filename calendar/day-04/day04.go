package main

import (
	"advent-of-go/utils/conv"
	"advent-of-go/utils/files"
	"fmt"
	"github.com/golang-collections/collections/grid"
	"strings"
)

type Board grid.Grid

func main() {
	input := files.ReadFile(4, "\n\n")
	println(solvePartOne(input))
	//println(solvePartTwo(input))
}

func solvePartOne(input []string) int {
	results := 0
	numbersDrawn := conv.ToIntSlice(strings.Split(input[0], ","))
	_ = numbersDrawn
	boardsToSplit := input[1:]
	for i, b := range boardsToSplit {
		fmt.Println(i, b)
		cols := 0
		rows := 0
		board := grid.New(cols, rows)
		_ = board
	}
	/*
		fmt.Println(boards[0])

		for _, num := range numbersDrawn {
			markNumber(num, boards)
			if hasWinner {
				unmarked := 0
				return num * unmarked
			}
		}
	*/

	return results
}

func solvePartTwo(input []string) int {
	results := 0

	return results
}
