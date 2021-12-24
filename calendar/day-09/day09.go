package main

import (
	"advent-of-go/utils/files"
	"strconv"
)

type point struct {
	x int
	y int
}

func (p point) getNeighbors() [4]point {
	return [4]point{
		{x: p.x, y: p.y + 1}, // up
		{x: p.x + 1, y: p.y}, // right
		{x: p.x, y: p.y - 1}, // down
		{x: p.x - 1, y: p.y}, // left
	}
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

func main() {
	input := files.ReadFile(9, "\n")
	println(solvePartOne(input))
	//println(solvePartTwo(input))
}

func solvePartOne(input []string) int {
	totalRisk := 0
	heightmap := make(grid, len(input))

	for i, row := range input {
		heights := make([]int, len(row))
		for j, c := range row {
			heights[j], _ = strconv.Atoi(string(c))
		}
		heightmap[i] = heights
	}

	for y, row := range heightmap {
		for x, val := range row {
			minHeight := 10
			pos := point{x, y}
			for _, neighbor := range pos.getNeighbors() {

				if heightmap.isInBounds(neighbor) == false {
					continue
				}

				if h := heightmap[neighbor.y][neighbor.x]; h < minHeight {
					minHeight = h
				}
			}

			if val < minHeight {
				totalRisk += 1 + val
				//fmt.Println(pos, val)
			}
		}
	}

	return totalRisk
}
