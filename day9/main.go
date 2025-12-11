package main

import (
	"fmt"
	"strings"

	"github.com/DChivers98/Advent-of-Code-2025/utils"
)

type Coordinate struct{ X, Y int }

func main() {
	lines := utils.ReadFileLines("./data/day9.txt")

	redTiles := make([]Coordinate, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, ",")
		redTiles = append(redTiles, Coordinate{utils.ToInt(parts[0]), utils.ToInt(parts[1])})
	}

	part1(redTiles)
}

func part1(redTiles []Coordinate) {
	var maxArea int
	for i, firstTile := range redTiles {
		for _, secondTile := range redTiles[i+1:] {
			maxArea = max(maxArea, (utils.AbsVal(secondTile.X-firstTile.X)+1)*(utils.AbsVal(secondTile.Y-firstTile.Y)+1))
		}
	}

	fmt.Printf("Part 1 result: %d\n", maxArea)
}
