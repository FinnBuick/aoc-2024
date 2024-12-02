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

func main() {
	lines, err := utils.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

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
