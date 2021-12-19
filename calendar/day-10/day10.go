package main

import (
	"advent-of-go/utils/files"
	"sort"
)

var partOneScores = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var partTwoScores = map[string]int{
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

var delimiterPairs = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	last := len(*s) - 1
	element := (*s)[last]
	*s = (*s)[:last]
	return element, true
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
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
			str := string(char)
			if opposite, ok := delimiterPairs[str]; ok {
				if last, ok := stack.Pop(); ok {
					if last != opposite {
						results += partOneScores[str]
					}
				}
			} else {
				stack.Push(str)
			}
		}
	}

	return results
}

func solvePartTwo(input []string) int {
	var allScores []int
	for _, line := range input {
		var stack Stack
		var isCorrupted bool
		for _, char := range line {
			str := string(char)
			if opposite, ok := delimiterPairs[str]; ok {
				if last, ok := stack.Pop(); ok {
					if last != opposite {
						isCorrupted = true
						break
					}
				}
			} else {
				stack.Push(str)
			}
		}

		if isCorrupted {
			continue
		}

		var stackScore int
		for !stack.IsEmpty() {
			currItem, _ := stack.Pop()

			score := partTwoScores[currItem]
			stackScore = 5*stackScore + score
		}
		allScores = append(allScores, stackScore)
	}

	sort.Ints(allScores)
	return allScores[len(allScores)/2]
}
