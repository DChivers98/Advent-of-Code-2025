package utils

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadFileLinesSplitOnBlank(filename string) ([]string, []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close() //nolint:errcheck

	scanner := bufio.NewScanner(file)

	var before, after []string
	part := &before
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" && part == &before {
			part = &after
			continue
		}
		*part = append(*part, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file %s", err)
	}

	return before, after
}

func ReadFileLines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close() //nolint:errcheck

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file %s", err)
	}

	return lines
}

func ReadFile(fileName string) string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic("Failed to open file: " + err.Error())
	}

	return strings.TrimSpace(string(data))
}

func ToInt(val string) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		panic("Error converting string to int: " + err.Error())
	}

	return num
}

func AbsVal(val int) int {
	return int(math.Abs(float64(val)))
}
