package main

import (
	"fmt"

	"github.com/DChivers98/Advent-of-Code-2025/utils"
)

type Position = utils.Position

func main() {
	manifold, start := utils.ReadFileLinesIntoGrid("./data/day7.txt", 'S')

	part1(manifold, start)
	part2(manifold, start)
}

func part1(manifold [][]string, start Position) {
	visited := make(map[Position]bool)

	var walkDown func(int, int) int
	walkDown = func(row, col int) int {
		position := Position{Row: row, Col: col}

		if row < 0 || row >= len(manifold) || col < 0 || col >= len(manifold[row]) || visited[position] {
			return 0
		}

		visited[position] = true

		if manifold[row][col] != "^" {
			return walkDown(row+1, col)
		}

		splitCount := 1

		if col-1 >= 0 {
			splitCount += walkDown(row, col-1)
		}
		if col+1 < len(manifold[row]) {
			splitCount += walkDown(row, col+1)
		}

		return splitCount
	}

	fmt.Printf("Part 1 result: %d\n", walkDown(start.Row, start.Col))
}

func part2(manifold [][]string, start Position) {
	visited := make(map[Position]bool)
	pathCountCache := map[Position]int{}

	var walkDown func(int, int) int
	walkDown = func(row, col int) int {
		position := Position{Row: row, Col: col}

		if row < 0 || row >= len(manifold) || col < 0 || col >= len(manifold[row]) || visited[position] {
			return 0
		}

		if row == len(manifold)-1 {
			return 1
		}

		if value, ok := pathCountCache[position]; ok {
			return value
		}
		visited[position] = true

		var pathCount int
		if manifold[row][col] == "^" {
			if col-1 >= 0 {
				pathCount += walkDown(row, col-1)
			}
			if col+1 < len(manifold[row]) {
				pathCount += walkDown(row, col+1)
			}
		} else {
			pathCount = walkDown(row+1, col)
		}

		visited[position] = false
		pathCountCache[position] = pathCount

		return pathCount
	}

	fmt.Printf("Part 2 result: %d\n", walkDown(start.Row, start.Col))
}
