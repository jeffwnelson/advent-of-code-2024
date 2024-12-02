package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/jeffwnelson/advent-of-code-2024/functions"
)

func main() {
	input, _ := os.Open("day01/input.txt")
	defer input.Close()

	fmt.Println("Part 1 solution: ", part1(input))
	input.Seek(0, 0) // Reset our scanner to start of file
	fmt.Println("Part 2 solution: ", part2(input))
}

func part1(input *os.File) int {
	scanner := bufio.NewScanner(input)
	var ll, rl []int

	for scanner.Scan() {
		line := scanner.Text()
		lists := strings.Fields(line)

		num1, _ := strconv.Atoi(lists[0])
		num2, _ := strconv.Atoi(lists[1])

		ll = append(ll, num1)
		rl = append(rl, num2)
	}

	sort.Ints(ll)
	sort.Ints(rl)

	res := 0
	for i := range ll {
		res += functions.AbsInt(ll[i] - rl[i])
	}

	return res
}

func part2(input2 *os.File) int {
	scanner := bufio.NewScanner(input2)
	var ll, rl []int

	for scanner.Scan() {
		line := scanner.Text()
		lists := strings.Fields(line)

		num1, _ := strconv.Atoi(lists[0])
		num2, _ := strconv.Atoi(lists[1])

		ll = append(ll, num1)
		rl = append(rl, num2)
	}

	frequency := make(map[int]int)
	for _, v := range rl {
		frequency[v]++
	}

	res := 0
	for _, v := range ll {
		res += v * frequency[v]
	}

	return res
}
