package main

import (
	"fmt"

	"github.com/DChivers98/Advent-of-Code-2025/utils"
)

func main() {
	banks := utils.ReadFileLinesIntoIntGrid("./data/day3.txt")

	part1(banks)
	part2(banks)
}

func part1(banks [][]int) {
	total := 0

	for _, bank := range banks {
		maxVoltage := 0
		for i := range bank {
			for j := i + 1; j < len(bank); j++ {
				voltage := bank[i]*10 + bank[j]
				if voltage > maxVoltage {
					maxVoltage = voltage
				}
			}
		}
		total += maxVoltage
	}

	fmt.Printf("Part 1 result: %d\n", total)
}

func part2(banks [][]int) {
	total := 0

	for _, bank := range banks {
		selected := make([]int, 0, 12)

		for i, value := range bank {
			for len(selected) > 0 && selected[len(selected)-1] < value && (len(bank)-i-1)+len(selected) >= 12 {
				selected = selected[:len(selected)-1]
			}

			if len(selected) < 12 {
				selected = append(selected, value)
			}
		}

		voltage := 0
		for _, value := range selected {
			voltage = voltage*10 + value
		}
		total += voltage
	}

	fmt.Printf("Part 2 result: %d\n", total)
}
