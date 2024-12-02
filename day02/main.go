package main

import (
	"bufio"
	"fmt"
	"github.com/jeffwnelson/advent-of-code-2024/functions"
	"os"
	"strings"
)

func main() {
	input, _ := os.Open("day02/input.txt")
	defer input.Close()

	fmt.Println("Part 1 solution: ", part1(input))
	input.Seek(0, 0) // Reset our scanner to start of file
	fmt.Println("Part 2 solution: ", part2(input))
}

func part1(input *os.File) int {
	reports := bufio.NewScanner(input)
	safeLevels := 0

	for reports.Scan() { // We are reading one line at a time by now...
		report := reports.Text()
		sLevels := strings.Split(report, " ")
		iLevels := functions.StringToInts(sLevels)

		if isSafe(iLevels) {
			safeLevels++
		}
	}

	return safeLevels
}

func isSafe(levels []int) bool {
	previous := levels[0]
	decreasing := true

	// Our first and second level are matching, cannot be safe
	if levels[0] == levels[1] {
		return false
	}

	// Our second level is higher than our first level, therefore we are increasing
	if levels[0] < levels[1] {
		decreasing = false
	}

	// Comparing each level to the previous
	for i := 1; i < len(levels); i++ {
		current := levels[i]

		// If we are decreasing, yet our previous number is smaller than our current number, that cannot be safe
		if decreasing && previous < current {
			return false
		}

		// If we are increasing, yet our previous number is larger than our current number, that cannot be safe
		if !decreasing && previous > current {
			return false
		}

		// Check if we are within our threshold
		threshold := functions.AbsInt(previous - current)
		// If our threshold is less than 1 (previous and current numbers match) or if our threshold is greater than 3, that cannot be safe
		if threshold < 1 || threshold > 3 {
			return false
		}

		// Our comparison is still valid, move to next set
		previous = current
	}
	// All checks passed, levels are safe
	return true
}

func part2(input *os.File) int {
	reports := bufio.NewScanner(input)
	safeLevels := 0

	for reports.Scan() { // We are reading one line at a time by now...
		report := reports.Text()
		sLevels := strings.Split(report, " ")
		iLevels := functions.StringToInts(sLevels)

		if isSafeWithProblemDampener(iLevels) {
			safeLevels++
		}
	}

	return safeLevels
}

func isSafeWithProblemDampener(levels []int) bool {
	// Check if the report is safe from the start...
	if isSafe(levels) {
		return true
	}

	// Loop through each level...
	for i := 0; i < len(levels); i++ {

		// Remove each element from the array/slice and retest each time
		for j := 0; j < len(levels); j++ {
			testSlice := make([]int, 0, len(levels)-1)     // Create a slice that's 1 size smaller than our original array
			testSlice = append(testSlice, levels[:j]...)   // Add all elements before our j
			testSlice = append(testSlice, levels[j+1:]...) // Add all elements after our j

			// If removing our single element fixed our report to make it safe, we found another safe one
			if isSafe(testSlice) {
				return true
			}
		}
	}

	return false
}
