package main

import (
	"fmt"
	"os"
)

func main() {
	inputBytes, _ := os.ReadFile("day12/test2.txt")
	input := string(inputBytes)
	fmt.Println("Part 1 solution: ", part1(input))
	//fmt.Println("Part 2 solution: ", part2(input))
}

func part1(input string) int {
	count := 0
	return count
}
