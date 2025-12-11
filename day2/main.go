package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DChivers98/Advent-of-Code-2025/utils"
)

func main() {
	data := utils.ReadFile("./data/day2.txt")

	productIDRanges := strings.Split(data, ",")

	var lowerBounds []int
	var upperBounds []int
	for _, productIDRange := range productIDRanges {
		parts := strings.Split(productIDRange, "-")
		lowerBounds = append(lowerBounds, utils.ToInt(parts[0]))
		upperBounds = append(upperBounds, utils.ToInt(parts[1]))
	}

	part1(lowerBounds, upperBounds)
	part2(lowerBounds, upperBounds)
}

func part1(lowerBounds []int, upperBounds []int) {
	count := 0
	for idx := range lowerBounds {
		lowerBound, upperBound := lowerBounds[idx], upperBounds[idx]
		for value := lowerBound; value <= upperBound; value++ {
			stringValue := strconv.Itoa(value)
			length := len(stringValue)

			if stringValue[:length/2] == stringValue[length/2:] {
				count += value
			}

		}
	}

	fmt.Printf("Part 1 result: %d\n", count)
}

func part2(lowerBounds []int, upperBounds []int) {
	count := 0
	for idx := range lowerBounds {
		lowerBound, upperBound := lowerBounds[idx], upperBounds[idx]
		for value := lowerBound; value <= upperBound; value++ {
			stringValue := strconv.Itoa(value)
			length := len(stringValue)

			for patternLength := 1; patternLength <= length/2; patternLength++ {
				pattern := stringValue[:patternLength]
				if strings.Repeat(pattern, length/patternLength) == stringValue {
					count += value
					break
				}
			}
		}
	}

	fmt.Printf("Part 2 result: %d\n", count)
}
