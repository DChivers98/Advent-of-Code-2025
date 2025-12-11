package main

import (
	"fmt"
	"strings"

	"github.com/DChivers98/Advent-of-Code-2025/utils"
)

func main() {
	lines := utils.ReadFileLines("./data/day6.txt")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
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

	fmt.Printf("Part 1 result: %d\n", total)
}

func part2(lines []string) {
	rows, cols := len(lines), len(lines[0])
	total, blockStart := 0, 0

	processBlock := func(end int) {
		operator := byte(' ')
		for col := blockStart; col < end; col++ {
			if char := lines[rows-1][col]; char != ' ' {
				operator = char
				break
			}
		}

		value := 0
		if operator == '+' {
			value = 0
		} else {
			value = 1
		}
		for col := blockStart; col < end; col++ {
			num := 0
			for row := 0; row < rows-1; row++ {
				if lines[row][col] != ' ' {
					num = num*10 + int(lines[row][col]-'0')
				}
			}
			if operator == '+' {
				value += num
			} else {
				value *= num
			}
		}

		total += value
	}

	for col := range cols {
		isBlankCol := true
		for row := range rows {
			if lines[row][col] != ' ' {
				isBlankCol = false
				break
			}
		}
		if isBlankCol {
			processBlock(col)
			blockStart = col + 1
		}
	}

	if blockStart < cols {
		processBlock(cols)
	}

	fmt.Printf("Part 2 result: %d\n", total)
}
