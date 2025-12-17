package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DChivers98/Advent-of-Code-2025/utils"
)

func main() {
	lines := utils.ReadFileLines("./data/day11.txt")

	devices := make(map[string][]string)

	for _, device := range lines {
		parts := strings.Split(device, ":")
		devices[parts[0]] = strings.Fields(parts[1])
	}

	part1(devices)
	part2(devices)
}

func part1(devices map[string][]string) {
	var walkPath func(device string) int
	walkPath = func(device string) int {
		paths := 0
		for _, connection := range devices[device] {
			switch connection {
			case "out":
				paths++
			default:
				paths += walkPath(connection)
			}
		}
		return paths
	}

	fmt.Printf("Part 1 result: %d\n", walkPath("you"))
}

func part2(devices map[string][]string) {
	paths := make(map[string]int)
	var walkPath func(device string, found int) int
	walkPath = func(device string, found int) int {
		if device == "out" {
			if found == 2 {
				return 1
			}
			return 0
		}

		key := device + ":" + strconv.Itoa(found)
		if value, ok := paths[key]; ok {
			return value
		}

		total := 0
		for _, connection := range devices[device] {
			if connection == "fft" || connection == "dac" {
				found++
			}

			total += walkPath(connection, found)
		}

		paths[key] = total
		return total
	}

	fmt.Printf("Part 2 result: %d\n", walkPath("svr", 0))
}
