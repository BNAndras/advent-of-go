package main

import (
	"advent-of-go/utils/conv"
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"sort"
	"strings"
)

func main() {
	input := files.ReadFile(7, "\n")
	println(solvePartOne(input))
	//println(solvePartTwo(input))
}

func solvePartOne(input []string) int {
	minFuel := maths.MaxInt()

	nums := strings.Split(input[0], ",")
	positions := conv.ToIntSlice(nums)
	sort.Ints(positions)
	minPos := positions[0]
	maxPos := positions[len(positions)-1]
	for i := minPos; i < maxPos; i++ {
		currFuel := 0
		for _, position := range positions {
			currFuel += maths.Abs(i - position)
			//fmt.Println(currFuel)
		}
		if currFuel < minFuel {
			minFuel = currFuel
		}
	}

	return minFuel
}

func solvePartTwo(input []string) int {
	results := 0

	return results
}
