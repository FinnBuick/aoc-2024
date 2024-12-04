package main

import (
	"aoc-2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isReportSafe(line string) bool {
	chars := strings.Split(line, " ")

	var levels []int
	for _, char := range chars {
		level, _ := strconv.Atoi(char)
		levels = append(levels, level)
	}

	allIncreasing := true
	allDecreasing := true
	differenceGreaterThanThree := true

	for i := 0; i < len(levels)-1; i++ {
		if levels[i] > levels[i+1] {
			allIncreasing = false
		}

		if levels[i] < levels[i+1] {
			allDecreasing = false
		}

		diff := math.Abs(float64(levels[i] - levels[i+1]))

		// diff is at least 1 but at most 3
		if diff > 3 && diff < 1 {
			differenceGreaterThanThree = false
		}
	}

	isSafe := allIncreasing || allDecreasing && !differenceGreaterThanThree

	fmt.Printf("line: %v isSafe: %b\n", line, isSafe)

	return isSafe
}

func main() {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	countSafe := 0
	for _, line := range lines {
		if isReportSafe(line) {
			countSafe++
		}
	}

	println(countSafe)
}
