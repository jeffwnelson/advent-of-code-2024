package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, _ := os.Open("day04/input.txt")
	defer grid.Close()

	fmt.Println("Part 1 solution: ", part1(grid))
	grid.Seek(0, 0) // Reset our scanner to start of file
	fmt.Println("Part 2 solution: ", part2(grid))
}

func part1(input *os.File) int {
	scanner := bufio.NewScanner(input)

	var grid []string

	for scanner.Scan() { // We are reading one line at a time by now...
		grid = append(grid, scanner.Text())
	}

	totalCount := findX(grid)
	return totalCount
}

func part2(input *os.File) int {
	scanner := bufio.NewScanner(input)

	var grid []string

	for scanner.Scan() { // We are reading one line at a time by now...
		grid = append(grid, scanner.Text())
	}

	totalCount := findA(grid)
	return totalCount
}

func findX(grid []string) int {
	count := 0

	for x := 0; x < len(grid); x++ {
		for y, character := range grid[x] {
			if character == 'X' {
				// We found an "X", so time to start branching
				count += checkHorizontal(grid, x, y)
				count += checkVertical(grid, x, y)
				count += checkDiagonally(grid, x, y)
			}
		}
	}

	return count
}

func checkHorizontal(grid []string, x int, y int) int {
	count := 0

	// Check forwards
	if y <= (len(grid[0]) - 1 - 3) {
		if grid[x][y] == byte('X') && grid[x][y+1] == byte('M') && grid[x][y+2] == byte('A') && grid[x][y+3] == byte('S') {
			count++
		}
	}

	// Check backwards
	if y >= 3 {
		if grid[x][y] == byte('X') && grid[x][y-1] == byte('M') && grid[x][y-2] == byte('A') && grid[x][y-3] == byte('S') {
			count++
		}
	}
	return count
}

func checkVertical(grid []string, x int, y int) int {
	count := 0

	// Check up
	if x >= 3 {
		if grid[x][y] == byte('X') && grid[x-1][y] == byte('M') && grid[x-2][y] == byte('A') && grid[x-3][y] == byte('S') {
			count++
		}
	}

	// Check down
	if x <= (len(grid[0]) - 1 - 3) {
		if grid[x][y] == byte('X') && grid[x+1][y] == byte('M') && grid[x+2][y] == byte('A') && grid[x+3][y] == byte('S') {
			count++
		}
	}
	return count
}

func checkDiagonally(grid []string, x int, y int) int {
	count := 0

	// Check up going right
	if x >= 3 && y <= (len(grid[0])-1-3) {
		if grid[x][y] == byte('X') &&
			grid[x-1][y+1] == byte('M') &&
			grid[x-2][y+2] == byte('A') &&
			grid[x-3][y+3] == byte('S') {
			count++
		}
	}

	// Check down going right
	if x <= (len(grid[0])-1-3) && y <= (len(grid[0])-1-3) {
		if grid[x][y] == byte('X') &&
			grid[x+1][y+1] == byte('M') &&
			grid[x+2][y+2] == byte('A') &&
			grid[x+3][y+3] == byte('S') {
			count++
		}
	}

	// Check up going left
	if x >= 3 && y >= 3 {
		if grid[x][y] == byte('X') &&
			grid[x-1][y-1] == byte('M') &&
			grid[x-2][y-2] == byte('A') &&
			grid[x-3][y-3] == byte('S') {
			count++
		}
	}

	// Check down going left
	if x <= (len(grid[0])-1-3) && y >= 3 {
		if grid[x][y] == byte('X') &&
			grid[x+1][y-1] == byte('M') &&
			grid[x+2][y-2] == byte('A') &&
			grid[x+3][y-3] == byte('S') {
			count++
		}
	}
	return count
}

func findA(grid []string) int {
	count := 0

	for x := 1; x < len(grid)-1; x++ {
		for y, character := range grid[x] {
			if character == 'A' {
				// We found an "A", so time to start branching
				count += checkPattern(grid, x, y)
			}
		}
	}

	return count
}

func checkPattern(grid []string, x, y int) int {
	count := 0

	// Check if we are out of bounds (or if we don't have another characters for a possible solution)
	if y < 1 || y > len(grid[x])-2 {
		return 0
	}

	// [M . M]
	// [. A .]
	// [S . S]
	if grid[x-1][y-1] == byte('M') &&
		grid[x-1][y+1] == byte('M') &&
		grid[x+1][y-1] == byte('S') &&
		grid[x+1][y+1] == byte('S') {
		count++
	}

	// [S . S]
	// [. A .]
	// [M . M]
	if grid[x-1][y-1] == byte('S') &&
		grid[x-1][y+1] == byte('S') &&
		grid[x+1][y-1] == byte('M') &&
		grid[x+1][y+1] == byte('M') {
		count++
	}

	// [M . S]
	// [. A .]
	// [M . S]
	if grid[x-1][y-1] == byte('M') &&
		grid[x-1][y+1] == byte('S') &&
		grid[x+1][y-1] == byte('M') &&
		grid[x+1][y+1] == byte('S') {
		count++
	}

	// [S . M]
	// [. A .]
	// [S . M]
	if grid[x-1][y-1] == byte('S') &&
		grid[x-1][y+1] == byte('M') &&
		grid[x+1][y-1] == byte('S') &&
		grid[x+1][y+1] == byte('M') {
		count++
	}
	return count
}
