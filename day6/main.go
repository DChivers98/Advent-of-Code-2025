package main

import (
	"fmt"
	"strings"

	"github.com/DChivers98/Advent-of-Code-2025/utils"
)

func main() {
	lines := utils.ReadFileLines("./data/day6.txt")

	var calculations [][]string
	for _, line := range lines {
		for i, element := range strings.Fields(line) {
			if i == len(calculations) {
				calculations = append(calculations, []string{element})
			} else {
				calculations[i] = append(calculations[i], element)
			}
		}
	}

	partA(calculations)
}

func partA(calculations [][]string) {
	total := 0
	for _, calculation := range calculations {
		switch calculation[4] {
		case "+":
			result := 0
			for _, num := range calculation[:4] {
				result += utils.ToInt(num)
			}
			total += result
		case "*":
			result := 1
			for _, num := range calculation[:4] {
				result *= utils.ToInt(num)
			}
			total += result
		}
	}

	fmt.Printf("Part A result: %d\n", total)
}
