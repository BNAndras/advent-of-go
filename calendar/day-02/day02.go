package main

import (
	"advent-of-go/utils/files"
	"strconv"
	"strings"
)

const (
	forward = "forward"
	up      = "up"
	down    = "down"
)

func main() {
	input := files.ReadFile(2, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	var x, y int

	for _, line := range input {
		words := strings.Split(line, " ")
		direction := words[0]
		magnitude, _ := strconv.Atoi(words[1])
		calculateChangeP1(direction, magnitude, &x, &y)
	}
	return x * y
}

func solvePart2(input []string) int {
	var x, y, aim int

	for _, line := range input {
		words := strings.Split(line, " ")
		direction := words[0]
		magnitude, _ := strconv.Atoi(words[1])
		calculateChangeP2(direction, magnitude, &x, &y, &aim)
	}
	return x * y
}

func calculateChangeP1(direction string, magnitude int, x, y *int) {
	switch direction {
	case forward:
		*x += magnitude
		break
	case down:
		*y += magnitude
		break
	case up:
		*y -= magnitude
		break
	}
}

func calculateChangeP2(direction string, magnitude int, x, y, aim *int) {
	switch direction {
	case forward:
		*x += magnitude
		*y += *aim * magnitude
		break
	case down:
		*aim += magnitude
		break
	case up:
		*aim -= magnitude
		break
	}
}
