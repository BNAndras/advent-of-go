package main

import (
	"advent-of-go/utils/files"
	"strings"
)

var uniqueSegmentCounts = map[int]bool{
	1: false, // not  a digit
	2: true,  // digit 1
	3: true,  // digit 7
	4: true,  // digit 4
	5: false, // digit 2, 3, or 5
	6: false, // digit 0, 6, or 9
	7: true,  // digit 8
}

func main() {
	input := files.ReadFile(8, "\n")
	println(solvePartOne(input))
	println(solvePartTwo(input))
}

func solvePartOne(input []string) int {
	results := 0

	for _, line := range input {
		sections := strings.Split(line, "|")
		outputValue := sections[1]
		digits := strings.Fields(outputValue)
		for _, digit := range digits {
			numSegments := len(digit)
			if uniqueSegmentCounts[numSegments] {
				results += 1
			}
		}
	}
	return results
}

func solvePartTwo(input []string) int {
	_ = input
	results := 0

	return results
}
