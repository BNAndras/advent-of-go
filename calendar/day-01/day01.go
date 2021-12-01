package main

import (
	"advent-of-go/utils/conv"
	"advent-of-go/utils/files"
)

func main() {
	input := files.ReadFile(1, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	vals := conv.ToIntSlice(input)

	return checkIncreasingDepths(vals, 1)
}

func solvePart2(input []string) int {
	vals := conv.ToIntSlice(input)

	return checkIncreasingDepths(vals, 3)
}

func checkIncreasingDepths(depths []int, n int) int {
	result := 0
	prev := sum(depths[:n])
	for i := n; i < len(depths); i++ {
		curr := prev - depths[i-n] + depths[i]
		if curr > prev {
			result += 1
		}
	}

	return result
}

func sum(arr []int) int {
	result := 0
	for _, v := range arr {
		result += v
	}

	return result
}
