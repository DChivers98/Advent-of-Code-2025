package main

import (
	"fmt"

	"github.com/DChivers98/Advent-of-Code-2025/utils"
)

type Position struct {
	row, col int
}

func main() {
	lines := utils.ReadFileLines("./data/day7.txt")

	var startPosition Position
	var manifold [][]string
	for row, line := range lines {
		var chars []string
		for col, char := range line {
			chars = append(chars, string(char))
			if char == 'S' {
				startPosition = Position{row, col}
			}
		}
		manifold = append(manifold, chars)
	}

	partA(manifold, startPosition)
	partB(manifold, startPosition)
}

func partA(manifold [][]string, start Position) {
	visited := make(map[Position]bool)

	var walkDown func(int, int) int
	walkDown = func(row, col int) int {
		position := Position{row, col}

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

	fmt.Printf("Part A result: %d\n", walkDown(start.row, start.col))
}

func partB(manifold [][]string, start Position) {
	visited := make(map[Position]bool)
	pathCountCache := map[Position]int{}

	var walkDown func(int, int) int
	walkDown = func(row, col int) int {
		position := Position{row, col}

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

	fmt.Printf("Part B result: %d\n", walkDown(start.row, start.col))
}
