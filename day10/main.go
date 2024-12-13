package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("day10/input.txt")
	input := string(inputBytes)
	fmt.Println("Part 1 solution: ", part1(input))
	fmt.Println("Part 2 solution: ", part2(input, 2))
}

func part1(input string) int {
	grid := createGrid(input)
	trailheads := findTrailHeads(grid)
	result := 0

	for _, trailhead := range trailheads {
		result += depthSearchFirst(grid, trailhead)
	}

	return result
}

func part2(input string, part int) int {
	grid := createGrid(input)
	trailheads := findTrailHeads(grid)
	result := 0

	for _, trailhead := range trailheads {
		visited := make(map[[2]int]bool)
		result += countPaths(grid, trailhead, visited)
	}

	return result
}

func createGrid(input string) [][]int {
	lines := strings.Split(input, "\n")
	var result [][]int

	for _, line := range lines {
		var row []int
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}
		result = append(result, row)
	}
	return result
}

func findTrailHeads(grid [][]int) [][2]int {
	var trailheads [][2]int

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				trailheads = append(trailheads, [2]int{i, j})
			}
		}
	}

	return trailheads
}

func breatheSearchFirst(grid [][]int, start [2]int) int {
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := [][2]int{start}
	visited := make(map[[2]int]bool)
	visited[start] = true
	count := 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, d := range directions {
			x, y := curr[0]+d[0], curr[1]+d[1]

			if inBoundary(grid, [2]int{x, y}) {
				next := [2]int{x, y}
				if !visited[next] && grid[x][y] == grid[curr[0]][curr[1]]+1 {
					visited[next] = true
					queue = append(queue, next)
					if grid[x][y] == 9 {
						count++
					}
				}
			}
		}
	}

	return count
}

func depthSearchFirst(grid [][]int, start [2]int) int {
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	stack := [][2]int{start}
	visited := make(map[[2]int]bool)
	visited[start] = true
	count := 0

	for len(stack) > 0 {
		// Pop the top element from the stack
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, d := range directions {
			x, y := curr[0]+d[0], curr[1]+d[1]

			if inBoundary(grid, [2]int{x, y}) {
				next := [2]int{x, y}
				if !visited[next] && grid[x][y] == grid[curr[0]][curr[1]]+1 {
					visited[next] = true
					stack = append(stack, next) // Push to stack
					if grid[x][y] == 9 {
						count++
					}
				}
			}
		}
	}

	return count
}

func inBoundary(grid [][]int, pos [2]int) bool {
	return pos[0] >= 0 && pos[0] < len(grid) && pos[1] >= 0 && pos[1] < len(grid[0])
}

func countPaths(grid [][]int, pos [2]int, visited map[[2]int]bool) int {
	if grid[pos[0]][pos[1]] == 9 {
		return 1
	}

	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited[pos] = true
	totalPaths := 0

	for _, d := range directions {
		x, y := pos[0]+d[0], pos[1]+d[1]
		next := [2]int{x, y}

		if inBoundary(grid, next) {
			if !visited[next] && grid[x][y] == grid[pos[0]][pos[1]]+1 {
				totalPaths += countPaths(grid, next, visited)
			}
		}
	}

	visited[pos] = false // backtrack
	return totalPaths
}
