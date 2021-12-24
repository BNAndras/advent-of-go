package main

import (
	"advent-of-go/utils/files"
	"math"
	"sort"
	"strconv"
)

var (
	invalidCellValue = -1
	basinEdgeValue   = 9
)

type point struct {
	x int
	y int
}

type grid [][]int

func (g grid) isInBounds(p point) bool {
	if p.y < 0 || p.y >= len(g) {
		return false
	}

	if p.x < 0 || p.x >= len(g[0]) {
		return false
	}

	return true
}

func (g grid) getNeighborsOf(p point) []point {

	var possible = []point{
		{x: p.x, y: p.y + 1}, // up
		{x: p.x + 1, y: p.y}, // right
		{x: p.x, y: p.y - 1}, // down
		{x: p.x - 1, y: p.y}, // left
	}

	var valid []point
	for _, p := range possible {
		if g.isInBounds(p) == false {
			continue
		}

		valid = append(valid, p)
	}

	return valid
}

func (g grid) getLowestPoints() []point {
	var lowPoints []point
	for y, row := range g {
		for x, val := range row {
			min := math.MaxInt
			pos := point{x, y}

			for _, neighbor := range g.getNeighborsOf(pos) {

				if h := g[neighbor.y][neighbor.x]; h < min {
					min = h
				}
			}

			if val < min {
				lowPoints = append(lowPoints, pos)
			}
		}
	}

	return lowPoints
}

func (g grid) getValueAt(p point) int {
	if g.isInBounds(p) == false {
		return invalidCellValue
	}

	return g[p.y][p.x]
}

func (g grid) getBasinFrom(end point) []point {
	queue := stack{end}
	seen := map[point]bool{
		end: true,
	}
	for queue.isEmpty() == false {
		current, _ := queue.pop()
		for _, neighbor := range g.getNeighborsOf(current) {
			if g.getValueAt(neighbor) == basinEdgeValue {
				continue
			}
			if _, ok := seen[neighbor]; ok {
				continue
			}

			queue = append(queue, neighbor)
			seen[neighbor] = true
		}
	}

	var keys []point
	for k := range seen {
		keys = append(keys, k)
	}
	return keys
}

type stack []point

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) pop() (point, bool) {
	if s.isEmpty() {
		return point{}, false
	}

	last := len(*s) - 1
	element := (*s)[last]
	*s = (*s)[:last]
	return element, true
}

func (s *stack) push(p point) {
	*s = append(*s, p)
}

func main() {
	input := files.ReadFile(9, "\n")
	println(solvePartOne(input))
	println(solvePartTwo(input))
}

func solvePartOne(input []string) int {
	g := buildGridFrom(input)

	var totalRisk int
	lowPoints := g.getLowestPoints()
	for _, p := range lowPoints {
		totalRisk += g.getValueAt(p) + 1
	}

	return totalRisk
}

func solvePartTwo(input []string) int {
	g := buildGridFrom(input)
	var basinSizes []int
	lowPoints := g.getLowestPoints() // low points for each basin
	for _, lowPoint := range lowPoints {
		basin := g.getBasinFrom(lowPoint)
		basinSizes = append(basinSizes, len(basin))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func buildGridFrom(input []string) grid {
	g := make(grid, len(input))
	for i, row := range input {
		heights := make([]int, len(row))
		for j, c := range row {
			heights[j], _ = strconv.Atoi(string(c))
		}
		g[i] = heights
	}

	return g
}
