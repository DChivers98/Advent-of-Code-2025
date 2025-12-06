package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/DChivers98/Advent-of-Code-2025/utils"
)

func main() {
	freshIngredientIDs, availableIngredientIDs := utils.ReadFileLinesSplitOnBlank("./data/day5.txt")

	freshIDRanges := make([]freshIDRange, 0, len(freshIngredientIDs))
	for _, freshIngredientID := range freshIngredientIDs {
		if idx := strings.IndexByte(freshIngredientID, '-'); idx >= 0 {
			lowerBound := utils.ToInt(freshIngredientID[:idx])
			higherBound := utils.ToInt(freshIngredientID[idx+1:])
			freshIDRanges = append(freshIDRanges, freshIDRange{lowerBound, higherBound})
		}
	}

	partA(freshIDRanges, availableIngredientIDs)
	partB(freshIDRanges)
}

type freshIDRange struct {
	lowerBound, higherBound int
}

func partA(freshIDRanges []freshIDRange, availableIngredientIDs []string) {
	matched := 0

	for _, id := range availableIngredientIDs {
		availableIngredientID := utils.ToInt(id)
		for _, freshIDRange := range freshIDRanges {
			if availableIngredientID > freshIDRange.lowerBound && availableIngredientID <= freshIDRange.higherBound {
				matched++
				break
			}
		}
	}

	fmt.Printf("Part A result: %d\n", matched)
}

func partB(freshIDRanges []freshIDRange) {
	sort.Slice(freshIDRanges, func(a, b int) bool {
		return freshIDRanges[a].lowerBound < freshIDRanges[b].lowerBound
	})

	total, currentID := 0, freshIDRanges[0]
	for _, freshIDRange := range freshIDRanges[1:] {
		if freshIDRange.lowerBound <= currentID.higherBound+1 {
			if freshIDRange.higherBound > currentID.higherBound {
				currentID.higherBound = freshIDRange.higherBound
			}
			continue
		}

		total += currentID.higherBound - currentID.lowerBound + 1
		currentID = freshIDRange
	}

	total += currentID.higherBound - currentID.lowerBound + 1

	fmt.Printf("Part B result: %d\n", total)
}
