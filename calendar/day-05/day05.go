package main

import (
	"advent-of-go/utils/files"
	"advent-of-go/utils/maths"
	"math"
	"regexp"
	"strconv"
)

const (
	X = Axis('X')
	Y = Axis('Y')
)

type Axis rune

type Coordinate struct {
	X int
	Y int
}

type Line struct {
	start   Coordinate
	current Coordinate
	end     Coordinate
}

func (l Line) direction(axis Axis) int {
	var a, b = l.chooseCoordinates(axis)

	if a < b {
		return 1
	} else if a > b {
		return -1
	}
	return 0
}

func (l Line) delta(axis Axis) int {
	var a, b = l.chooseCoordinates(axis)

	return maths.Abs(b - a)
}

func (l Line) chooseCoordinates(axis Axis) (int, int) {
	var a, b int
	if axis == X {
		a = l.start.X
		b = l.end.X
	} else {
		a = l.start.Y
		b = l.end.Y
	}
	return a, b
}

func main() {
	input := files.ReadFile(5, "\n")
	println(solvePart1(input))
	println(solvePart2(input))
}

func solvePart1(input []string) int {
	markedCoordinates := map[Coordinate]int{}
	overlaps := 0
	for _, line := range input {
		re, _ := regexp.Compile("(\\d+),(\\d+)\\s->\\s(\\d+),(\\d+)")
		nums := re.FindStringSubmatch(line)
		start := Coordinate{
			X: MustAtoi(nums[1]),
			Y: MustAtoi(nums[2]),
		}
		end := Coordinate{
			X: MustAtoi(nums[3]),
			Y: MustAtoi(nums[4]),
		}
		line := Line{start, start, end}
		//fmt.Println(line)
		directionX := line.direction(X)
		directionY := line.direction(Y)
		deltaX := line.delta(X)
		deltaY := line.delta(Y)

		// no diagonals yet!
		if deltaX != 0 && deltaY != 0 {
			continue
		}
		deltaMax := int(math.Max(float64(deltaX), float64(deltaY)))
		for i := 0; i <= deltaMax; i++ {
			markedCoordinates[line.current] += 1
			if markedCoordinates[line.current] == 2 {
				overlaps += 1
			}

			line.current.X += directionX
			line.current.Y += directionY
		}
	}
	return overlaps
}

func solvePart2(input []string) int {
	markedCoordinates := map[Coordinate]int{}
	overlaps := 0
	for _, line := range input {
		re, _ := regexp.Compile("(\\d+),(\\d+)\\s->\\s(\\d+),(\\d+)")
		nums := re.FindStringSubmatch(line)
		start := Coordinate{
			X: MustAtoi(nums[1]),
			Y: MustAtoi(nums[2]),
		}
		end := Coordinate{
			X: MustAtoi(nums[3]),
			Y: MustAtoi(nums[4]),
		}
		line := Line{start, start, end}
		//fmt.Println(line)
		directionX := line.direction(X)
		directionY := line.direction(Y)
		deltaX := line.delta(X)
		deltaY := line.delta(Y)

		// only straight or 45-degree lines allowed
		if deltaX != 0 && deltaY != 0 && deltaX != deltaY {
			continue
		}
		deltaMax := int(math.Max(float64(deltaX), float64(deltaY)))
		for i := 0; i <= deltaMax; i++ {
			markedCoordinates[line.current] += 1
			if markedCoordinates[line.current] == 2 {
				overlaps += 1
			}

			line.current.X += directionX
			line.current.Y += directionY
		}
	}
	return overlaps
}

func MustAtoi(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}
