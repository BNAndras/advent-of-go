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
	result := 0
	vals := conv.ToIntSlice(input)
	//vals = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	for i := 1; i < len(vals); i++ {
		if vals[i] > vals[i-1] {
			result += 1
		}
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	vals := conv.ToIntSlice(input)
	//vals = []int{607, 618, 618, 617, 647, 716, 769, 792}

	prev := sum(vals[:3])

	for i := 3; i < len(vals); i++ {
		curr := prev - vals[i-3] + vals[i]
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
