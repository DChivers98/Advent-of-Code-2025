package main

import (
	"fmt"

	"github.com/DChivers98/Advent-of-Code-2025/utils"
)

func main() {
	banks := utils.ReadFileLines("./data/day3.txt")

	intBanks := make([][]int, 0, len(banks))
	for _, bank := range banks {
		bankVoltages := make([]int, 0, len(bank))
		for _, character := range bank {
			bankVoltages = append(bankVoltages, int(character-'0'))
		}
		intBanks = append(intBanks, bankVoltages)
	}

	partA(intBanks)
	partB(intBanks)
}

func partA(banks [][]int) {
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

	fmt.Printf("Part A result: %d\n", total)
}

func partB(banks [][]int) {
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

	fmt.Printf("Part B result: %d\n", total)
}
