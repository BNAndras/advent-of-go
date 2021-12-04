package main

import (
	"advent-of-go/utils/files"
	"strconv"
)

const (
	one  = '1'
	zero = '0'
)

func main() {
	input := files.ReadFile(3, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int64 {
	var sGamma, sEpsilon string

	for i := range input[0] {
		ones, zeroes := countOnesAndZeroes(input, i)

		if ones > zeroes {
			sGamma += string(one)
			sEpsilon += string(zero)
		} else {
			sGamma += string(zero)
			sEpsilon += string(one)
		}
	}

	//https://stackoverflow.com/questions/9271469/go-convert-string-which-represent-binary-number-into-int
	gamma, _ := strconv.ParseInt(sGamma, 2, 64)
	epsilon, _ := strconv.ParseInt(sEpsilon, 2, 64)
	return gamma * epsilon
}

func solvePart2(input []string) int {
	results := 0

	return results
}

func countOnesAndZeroes(input []string, position int) (int, int) {
	var ones, zeroes int
	for _, line := range input {

		switch line[position] {
		case one:
			ones += 1
		case zero:
			zeroes += 1
		}
	}

	return ones, zeroes
}
