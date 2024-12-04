package main

import (
	"aoc-2024/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func part1(lines []string) {
	var firstList, secondList []int

	for _, line := range lines {
		fields := strings.Fields(line)
		num1, _ := strconv.Atoi(fields[0])
		num2, _ := strconv.Atoi(fields[1])
		firstList = append(firstList, num1)
		secondList = append(secondList, num2)
	}

	sort.Ints(firstList)
	sort.Ints(secondList)

	if len(firstList) != len(secondList) {
		fmt.Print("lists are not equal size")
	}

	var totalDistance int
	for i := range firstList {
		totalDistance += diff(firstList[i], secondList[i])
	}

	fmt.Printf("Answer: %d\n", totalDistance)
}

func part2(lines []string) {
	var firstList, secondList []int

	for _, line := range lines {
		fields := strings.Fields(line)
		num1, _ := strconv.Atoi(fields[0])
		num2, _ := strconv.Atoi(fields[1])
		firstList = append(firstList, num1)
		secondList = append(secondList, num2)
	}

	totalScore := 0
	for i := range firstList {
		current := firstList[i]
		for j := range secondList {
			if secondList[j] == current {
				totalScore += current
			}
		}
	}

	fmt.Printf("Answer: %d\n", totalScore)
}

func main() {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	part2(lines)
}
