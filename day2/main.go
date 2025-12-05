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

	partA(lowerBounds, upperBounds)
	partB(lowerBounds, upperBounds)
}

func partA(lowerBounds []int, upperBounds []int) {
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

	fmt.Printf("Part A result: %d\n", count)
}

func partB(lowerBounds []int, upperBounds []int) {
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

	fmt.Printf("Part B result: %d\n", count)
}
