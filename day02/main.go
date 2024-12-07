package main

import (
	"aoc-2024/utils"
	"math"
	"strconv"
	"strings"
)

func partOne(line string) bool {
	chars := strings.Split(line, " ")

	var levels []int
	for _, char := range chars {
		level, _ := strconv.Atoi(char)
		levels = append(levels, level)
	}

	allIncreasing := true
	allDecreasing := true
	var maxDiff float64

	for i := 0; i < len(levels)-1; i++ {
		var diff float64
		if levels[i] >= levels[i+1] {
			allIncreasing = false
		}
		if levels[i] <= levels[i+1] {
			allDecreasing = false
		}
		diff = math.Abs(float64(levels[i] - levels[i+1]))
		maxDiff = max(diff, maxDiff)
	}
	differenceIsAtleastOneAndAtMostThree := maxDiff >= 1 && maxDiff <= 3
	isSafe := allIncreasing != allDecreasing && differenceIsAtleastOneAndAtMostThree
	return isSafe
}

func main() {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	countSafe := 0
	for _, line := range lines {
		if partOne(line) {
			countSafe++
		}
	}

	println(countSafe)
}
