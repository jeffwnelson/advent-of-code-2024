package main

import (
	"fmt"
	"github.com/jeffwnelson/advent-of-code-2024/functions"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("day07/input.txt")
	input := string(inputBytes)
	fmt.Println("Part 1 solution: ", part1(input))
	fmt.Println("Part 2 solution: ", part2(input))
}

func part1(input string) int {
	equations := strings.Split(input, "\n")
	//var trueEquations []string

	solvableCounts := 0
	for _, equation := range equations {
		sections := strings.Split(equation, ":")
		result, _ := strconv.Atoi(sections[0])
		ints := functions.StringToIntsSlice(sections[1])

		if isSolveable(result, 0, ints) {
			solvableCounts += result
		}
	}

	return solvableCounts
}

func part2(input string) int {
	equations := strings.Split(input, "\n")
	//var trueEquations []string

	solvableCounts := 0
	for _, equation := range equations {
		sections := strings.Split(equation, ":")
		result, _ := strconv.Atoi(sections[0])
		ints := functions.StringToIntsSlice(sections[1])

		if isSolveable2(result, 0, ints) {
			solvableCounts += result
		}
	}

	return solvableCounts
}

func isSolveable(result int, current int, ints []int) bool {
	// Check if we are starting at the beginning of our list of ints
	if current == 0 {
		return isSolveable(result, ints[0], ints[1:])
	}

	// We are at the end of our list of ends
	if len(ints) == 0 {
		return result == current
	}

	// If are current running total is > than our result, not solvable
	if current > result {
		return false
	}

	// Run the check again adding the int we are at to our current running total
	canAdd := isSolveable(result, current+ints[0], ints[1:])

	// Run the check again multiplying the int we are at to our current running total
	canMultiply := isSolveable(result, current*ints[0], ints[1:])
	return canAdd || canMultiply
}

func isSolveable2(result int, current int, ints []int) bool {
	// Check if we are starting at the beginning of our list of ints
	if current == 0 {
		return isSolveable2(result, ints[0], ints[1:])
	}

	// We are at the end of our list of ends
	if len(ints) == 0 {
		return result == current
	}

	// If are current running total is > than our result, not solvable
	if current > result {
		return false
	}

	// Run the check again - adding
	canAdd := isSolveable2(result, current+ints[0], ints[1:])

	// Run the check again - multiplying
	canMultiply := isSolveable2(result, current*ints[0], ints[1:])

	// Run the check again - concatenating
	stringNumber, _ := strconv.Atoi(strconv.Itoa(current) + strconv.Itoa(ints[0]))
	canConcatenate := isSolveable2(result, stringNumber, ints[1:])

	return canAdd || canMultiply || canConcatenate
}
