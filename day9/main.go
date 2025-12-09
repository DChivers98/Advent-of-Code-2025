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

	partA(redTiles)
}

func partA(redTiles []Coordinate) {
	var maxArea int
	for firstTileIndex := range redTiles {
		for secondTileIndex := firstTileIndex + 1; secondTileIndex < len(redTiles); secondTileIndex++ {
			x1, y1 := redTiles[firstTileIndex].X, redTiles[firstTileIndex].Y
			x2, y2 := redTiles[secondTileIndex].X, redTiles[secondTileIndex].Y

			maxArea = max(maxArea, (utils.AbsVal(x2-x1)+1)*(utils.AbsVal(y2-y1)+1))
		}
	}

	fmt.Printf("Part A result: %d\n", maxArea)
}
