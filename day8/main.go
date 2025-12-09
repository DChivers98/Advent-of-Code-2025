package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"

	"github.com/DChivers98/Advent-of-Code-2025/utils"
)

type JunctionBox struct {
	X, Y, Z int
}

type Distance struct {
	firstJunctionBox, secondJunctionBox int
	distance                            float64
}

func main() {
	lines := utils.ReadFileLines("./data/day8.txt")

	var junctionBoxes []JunctionBox
	for _, line := range lines {
		parts := strings.Split(line, ",")

		junctionBox := JunctionBox{
			X: utils.ToInt(parts[0]),
			Y: utils.ToInt(parts[1]),
			Z: utils.ToInt(parts[2]),
		}

		junctionBoxes = append(junctionBoxes, junctionBox)
	}

	partA(junctionBoxes)
	partB(junctionBoxes)
}

func partA(junctionBoxes []JunctionBox) {
	var distances []Distance
	for i := range junctionBoxes {
		for j := i + 1; j < len(junctionBoxes); j++ {
			dx := float64(junctionBoxes[i].X - junctionBoxes[j].X)
			dy := float64(junctionBoxes[i].Y - junctionBoxes[j].Y)
			dz := float64(junctionBoxes[i].Z - junctionBoxes[j].Z)
			dist := math.Sqrt(dx*dx + dy*dy + dz*dz)

			distances = append(distances, Distance{
				firstJunctionBox:  i,
				secondJunctionBox: j,
				distance:          dist,
			})
		}
	}

	sort.Slice(distances, func(a, b int) bool {
		return distances[a].distance < distances[b].distance
	})

	var circuits [][]int
	for _, connection := range distances[:1000] {
		junctionBoxIndex := slices.IndexFunc(circuits, func(circuit []int) bool {
			return slices.Contains(circuit, connection.firstJunctionBox)
		})

		connectingJunctionBoxIndex := slices.IndexFunc(circuits, func(circuit []int) bool {
			return slices.Contains(circuit, connection.secondJunctionBox)
		})

		switch {
		case junctionBoxIndex >= 0 && connectingJunctionBoxIndex >= 0:
			if junctionBoxIndex != connectingJunctionBoxIndex {
				circuits[junctionBoxIndex] = append(circuits[junctionBoxIndex], circuits[connectingJunctionBoxIndex]...)
				circuits = slices.Delete(circuits, connectingJunctionBoxIndex, connectingJunctionBoxIndex+1)
			}
		case junctionBoxIndex >= 0:
			circuits[junctionBoxIndex] = append(circuits[junctionBoxIndex], connection.secondJunctionBox)
		case connectingJunctionBoxIndex >= 0:
			circuits[connectingJunctionBoxIndex] = append(circuits[connectingJunctionBoxIndex], connection.firstJunctionBox)
		default:
			circuits = append(circuits, []int{connection.firstJunctionBox, connection.secondJunctionBox})
		}
	}

	slices.SortFunc(circuits, func(a, b []int) int {
		switch {
		case len(a) < len(b):
			return 1
		case len(a) > len(b):
			return -1
		default:
			return 0
		}
	})

	total := 1
	for _, circuit := range circuits[:3] {
		total *= len(circuit)
	}

	fmt.Printf("Part A result: %d\n", total)
}

func partB(junctionBoxes []JunctionBox) {
	var distances []Distance
	for i := range junctionBoxes {
		for j := i + 1; j < len(junctionBoxes); j++ {
			dx := float64(junctionBoxes[i].X - junctionBoxes[j].X)
			dy := float64(junctionBoxes[i].Y - junctionBoxes[j].Y)
			dz := float64(junctionBoxes[i].Z - junctionBoxes[j].Z)
			dist := math.Sqrt(dx*dx + dy*dy + dz*dz)

			distances = append(distances, Distance{
				firstJunctionBox:  i,
				secondJunctionBox: j,
				distance:          dist,
			})
		}
	}

	sort.Slice(distances, func(a, b int) bool {
		return distances[a].distance < distances[b].distance
	})

	var circuits [][]int
	var lastJunctionBox, lastConnectingJunctionBox JunctionBox
	for _, connection := range distances {
		junctionBoxIndex := slices.IndexFunc(circuits, func(circuit []int) bool {
			return slices.Contains(circuit, connection.firstJunctionBox)
		})

		connectingJunctionBoxIndex := slices.IndexFunc(circuits, func(circuit []int) bool {
			return slices.Contains(circuit, connection.secondJunctionBox)
		})

		changed := false

		switch {
		case junctionBoxIndex >= 0 && connectingJunctionBoxIndex >= 0 && junctionBoxIndex != connectingJunctionBoxIndex:
			circuits[junctionBoxIndex] = append(circuits[junctionBoxIndex], circuits[connectingJunctionBoxIndex]...)
			circuits = slices.Delete(circuits, connectingJunctionBoxIndex, connectingJunctionBoxIndex+1)
			changed = true

		case junctionBoxIndex >= 0 && connectingJunctionBoxIndex < 0:
			circuits[junctionBoxIndex] = append(circuits[junctionBoxIndex], connection.secondJunctionBox)
			changed = true

		case junctionBoxIndex < 0 && connectingJunctionBoxIndex >= 0:
			circuits[connectingJunctionBoxIndex] = append(circuits[connectingJunctionBoxIndex], connection.firstJunctionBox)
			changed = true

		case junctionBoxIndex < 0 && connectingJunctionBoxIndex < 0:
			circuits = append(circuits, []int{connection.firstJunctionBox, connection.secondJunctionBox})
			changed = true
		}

		if changed {
			lastJunctionBox = junctionBoxes[connection.firstJunctionBox]
			lastConnectingJunctionBox = junctionBoxes[connection.secondJunctionBox]
		}
	}

	fmt.Printf("Part B result: %d\n", lastJunctionBox.X*lastConnectingJunctionBox.X)
}
