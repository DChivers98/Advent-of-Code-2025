package main

import (
	"fmt"

	"github.com/DChivers98/Advent-of-Code-2025/utils"
)

func main() {
	grid, _ := utils.ReadFileLinesIntoGrid("./data/day4.txt", 0)

	part1(grid)
	part2(grid)
}

func part1(grid [][]string) {
	total := 0
	directions := [][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}

	for i := range len(grid) {
		for j := range len(grid[i]) {
			if grid[i][j] != "@" {
				continue
			}
			count := 0
			for _, direction := range directions {
				newX, newY := i+direction[0], j+direction[1]
				if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[newX]) {
					if grid[newX][newY] == "@" {
						count++
					}
				}
			}
			if count < 4 {
				total++
			}
		}
	}

	fmt.Printf("Part 1 result: %d\n", total)
}

func part2(grid [][]string) {
	total := 0
	directions := [][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}

	var removeRolls func(x, y int)
	removeRolls = func(x, y int) {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) {
			return
		}
		if grid[x][y] != "@" {
			return
		}

		neighbour := 0
		for _, direction := range directions {
			newX, newY := x+direction[0], y+direction[1]
			if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[newX]) && grid[newX][newY] == "@" {
				neighbour++
			}
		}

		if neighbour < 4 {
			total++
			grid[x][y] = "."
			for _, direction := range directions {
				removeRolls(x+direction[0], y+direction[1])
			}
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "@" {
				removeRolls(i, j)
			}
		}
	}

	fmt.Printf("Part 2 result: %d\n", total)
}
