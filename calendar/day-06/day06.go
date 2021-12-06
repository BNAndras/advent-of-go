package main

import (
	"advent-of-go/utils/conv"
	"advent-of-go/utils/files"
	"strings"
)

func main() {
	input := files.ReadFile(6, "\n")
	println(solvePartOne(input))
	println(solvePartTwo(input))
}

func solvePartOne(input []string) int {
	var results int
	nums := strings.Split(input[0], ",")
	fishes := conv.ToIntSlice(nums)

	fishAges := make(map[int]int)
	for _, day := range fishes {
		fishAges[day] += 1
	}

	for i := 0; i < 80; i++ {

		currentAges := make(map[int]int)

		for d, c := range fishAges {
			if d > 0 {
				currentAges[d-1] += c
			} else {
				currentAges[6] += c
				currentAges[8] += c
			}
		}

		for d := range fishAges {
			delete(fishAges, d)
		}

		for d, c := range currentAges {
			fishAges[d] = c
		}
	}

	for _, c := range fishAges {
		results += c
	}

	return results
}

func solvePartTwo(input []string) int {
	var results int
	nums := strings.Split(input[0], ",")
	fishes := conv.ToIntSlice(nums)

	fishAges := make(map[int]int)
	for _, day := range fishes {
		fishAges[day] += 1
	}

	for i := 0; i < 256; i++ {

		currentAges := make(map[int]int)

		for d, c := range fishAges {
			if d > 0 {
				currentAges[d-1] += c
			} else {
				currentAges[6] += c
				currentAges[8] += c
			}
		}

		for d := range fishAges {
			delete(fishAges, d)
		}

		for d, c := range currentAges {
			fishAges[d] = c
		}
	}

	for _, c := range fishAges {
		results += c
	}

	return results
}
