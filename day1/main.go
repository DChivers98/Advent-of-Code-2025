package main

import (
	"fmt"

	"github.com/DChivers98/Advent-of-Code-2025/utils"
)

func main() {
	lines := utils.ReadFileLines("./data/day1.txt")

	var directions []string
	var moveAmmounts []int
	for _, line := range lines {
		directions = append(directions, string(line[0]))
		moveAmmounts = append(moveAmmounts, utils.ToInt(line[1:]))
	}

	part1(directions, moveAmmounts)
	part2(directions, moveAmmounts)
}

func part1(directions []string, moveAmmounts []int) {
	total := 0
	currentPosition := 50

	for i := 0; i < len(directions) && i < len(moveAmmounts); i++ {
		moveAmmount := moveAmmounts[i] % 100

		if directions[i] == "L" {
			moveAmmount = -moveAmmount
		}

		currentPosition = (currentPosition + moveAmmount + 100) % 100

		if currentPosition == 0 {
			total++
		}
	}

	fmt.Printf("Part 1 result: %d\n", total)
}

func part2(directions []string, moveAmounts []int) {
	total := 0
	currentPosition := 50

	for i := 0; i < len(directions) && i < len(moveAmounts); i++ {
		originalPosition := moveAmounts[i]
		step := originalPosition % 100
		if directions[i] == "L" {
			step = -step
		}

		total += utils.AbsVal(originalPosition) / 100

		oldPosition := currentPosition
		currentPosition = (currentPosition + step + 100) % 100

		if currentPosition == 0 && oldPosition != 0 {
			total++
		}

		newPosition := oldPosition + step
		if (newPosition >= 100 || newPosition < 0) && oldPosition != 0 && currentPosition != 0 {
			total++
		}
	}

	fmt.Printf("Part 2 result: %d\n", total)
}
