package main

import (
	"advent-of-go/utils/files"
)

var partOneScores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var delimiters = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

type Stack []rune

func (s *Stack) Pop() (rune, bool) {
	if len(*s) == 0 {
		return 0, false
	} else {
		last := len(*s) - 1
		element := (*s)[last]
		*s = (*s)[:last]
		return element, true
	}
}

func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

func main() {
	input := files.ReadFile(10, "\n")
	println(solvePartOne(input))
	println(solvePartTwo(input))
}

func solvePartOne(input []string) int {
	results := 0

	for _, line := range input {
		var stack Stack
		for _, char := range line {
			if opposite, ok := delimiters[char]; ok {
				if last, ok := stack.Pop(); ok {
					if last != opposite {
						results += partOneScores[char]
					}
				}
			} else {
				stack.Push(char)
			}
		}
	}

	return results
}

func solvePartTwo(input []string) int {
	results := 0

	return results
}
