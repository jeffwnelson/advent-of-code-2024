package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, _ := os.Open("day03/input.txt")
	defer input.Close()

	fmt.Println("Part 1 solution: ", part1(input))
	input.Seek(0, 0) // Reset our scanner to start of file
	fmt.Println("Part 2 solution: ", part2(input))
}

func part1(input *os.File) int {
	scanner := bufio.NewScanner(input)
	total := 0

	for scanner.Scan() { // We are reading one line at a time by now...
		memory := scanner.Text()
		mulArray := findMuls(memory)

		for _, value := range mulArray {
			total += muls(value)
		}
	}
	return total
}

func part2(input *os.File) int {
	scanner := bufio.NewScanner(input)
	total := 0
	keepMul := true // Start off keeping matches

	for scanner.Scan() { // We are reading one line at a time by now...
		memory := scanner.Text()
		mulArray := findMulsWithConditions(memory)

		for _, value := range mulArray {
			if strings.Contains(value, "don't()") { // If we find a "don't()" string, then any muls found after this are ignored
				keepMul = false
				continue
			} else if strings.Contains(value, "do()") { // If we find a "do()" string, then any muls found after this should be kept
				keepMul = true
				continue
			} else { // This string should be a "mul(x,y)" string...
				if keepMul {
					total += muls(value)
				}
			}
		}
	}
	return total
}

// Convert "mul(x,y)" to result of x*y
func muls(string string) int {
	var x, y int
	fmt.Sscanf(string, "mul(%d,%d)", &x, &y)
	return x * y
}

// Create an array of matches where the pattern seen is "mul(x,y)"
func findMuls(string string) []string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(string, -1)
	return matches
}

// Create an array of matches where the pattern seen is "mul(x,y)", "do()", and "don't()"
func findMulsWithConditions(string string) []string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	matches := re.FindAllString(string, -1)
	return matches
}
