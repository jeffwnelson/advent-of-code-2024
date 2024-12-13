package main

import (
	"fmt"
	"os"
)

func main() {
	inputBytes, _ := os.ReadFile("day07/test.txt")
	input := string(inputBytes)
	fmt.Println("Part 1 solution: ", part1(input))
	//fmt.Println("Part 2 solution: ", part2(input))
}

func part1(input string) int {

	return 0
}
