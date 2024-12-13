package main

import (
	"fmt"
	"os"
	"strings"
)

var grid [][]string
var gridOriginal [][]string
var loopCounter int = 0
var obstacleTotalCount int = 0

func main() {
	inputBytes, _ := os.ReadFile("day06/input.txt")
	input := string(inputBytes)
	fmt.Println("Part 1 solution: ", part1(input))
	//fmt.Println("Part 2 solution: ", part2(input))
}

func part1(input string) int {
	steps := 0
	getGrid(input)

	fmt.Println("Starting Grid")
	showGrid()

	for !edgeCheck() {
		if canGuardMoveForward() {
			moveGuardForward()
		} else {
			rotateGuardRight()
		}
		steps++
	}

	fmt.Println("Resulting Grid")
	showGrid()
	fmt.Printf("We took %v steps before we left the map!\n", steps)
	return calculatePositions()
}

func part2(input string) int {
	steps := 0
	obstacleChecks := 0
	getGrid(input)
	copyGrid()

	for y, row := range grid {
		for x := range row {
			loopCounter = 0
			obstacleChecks++

			if grid[y][x] != "^" {
				grid[y][x] = "#"
			}
			//showGrid()

			for !edgeCheck() && !areWeLooping() {
				if canGuardMoveForward() {
					moveGuardForward()
				} else {
					rotateGuardRight()
				}
				steps++
				//showGrid()
			}
			resetGrid()
			fmt.Printf("Checked @ [%v, %v] [%v]\n", y, x, obstacleChecks)
		}
	}

	fmt.Printf("We found %v obstacles!\n", obstacleTotalCount)
	return obstacleTotalCount
}

func getGrid(input string) {
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		columns := strings.Split(row, "")
		grid = append(grid, columns)
	}
}

func areWeLooping() bool {
	x, y := findGuard()

	if loopCounter > 1000 {
		obstacleTotalCount++
		fmt.Println("Loop detected!")
		return true
	}

	if grid[y][x] == "^" {
		if grid[y-1][x] == "X" {
			loopCounter++
		}
	}

	if grid[y][x] == ">" {
		if grid[y][x+1] == "X" {
			loopCounter++
		}
	}

	if grid[y][x] == "v" {
		if grid[y+1][x] == "X" {
			loopCounter++
		}
	}

	if grid[y][x] == "<" {
		if grid[y][x-1] == "X" {
			loopCounter++
		}
	}

	return false
}

func canGuardMoveForward() bool {
	x, y := findGuard()

	if grid[y][x] == "^" {
		if grid[y-1][x] != "#" {
			return true
		}
	}

	if grid[y][x] == ">" {
		if grid[y][x+1] != "#" {
			return true
		}
	}

	if grid[y][x] == "v" {
		if grid[y+1][x] != "#" {
			return true
		}
	}

	if grid[y][x] == "<" {
		if grid[y][x-1] != "#" {
			return true
		}
	}

	return false
}

func moveGuardForward() {
	x, y := findGuard()

	if grid[y][x] == "^" {
		grid[y-1][x] = "^"
		grid[y][x] = "X"
	}

	if grid[y][x] == ">" {
		grid[y][x+1] = ">"
		grid[y][x] = "X"
	}

	if grid[y][x] == "v" {
		grid[y+1][x] = "v"
		grid[y][x] = "X"
	}

	if grid[y][x] == "<" {
		grid[y][x-1] = "<"
		grid[y][x] = "X"
	}
}

func rotateGuardRight() {
	x, y := findGuard()

	if grid[y][x] == "^" { // If we are facing up, rotate right (now we are facing right)
		grid[y][x] = ">"
	} else if grid[y][x] == ">" { // If we are facing right, rotate right (now we are facing down)
		grid[y][x] = "v"
	} else if grid[y][x] == "v" { // If we are facing down, rotate right (now we are facing left)
		grid[y][x] = "<"
	} else if grid[y][x] == "<" { // If we are facing left, rotate right (now we are facing up)
		grid[y][x] = "^"
	}
}

func findGuard() (x, y int) {
	var xPos, yPos int
	for y, row := range grid {
		for x, column := range row {
			if strings.Contains(column, "^") || strings.Contains(column, ">") || strings.Contains(column, "v") || strings.Contains(column, "<") {
				xPos = x
				yPos = y
			}
		}
	}
	return xPos, yPos
}

func edgeCheck() bool {
	x, y := findGuard()

	if grid[y][x] == "^" && y == 0 {
		grid[y][x] = "X"
		return true
	}

	if grid[y][x] == ">" && x == len(grid[0])-1 {
		grid[y][x] = "X"
		return true
	}

	if grid[y][x] == "v" && y == len(grid)-1 {
		grid[y][x] = "X"
		return true
	}

	if grid[y][x] == "<" && x == 0 {
		grid[y][x] = "X"
		return true
	}

	return false
}

func showGrid() {
	for y, row := range grid {
		fmt.Println(y, row)
	}
	fmt.Println("")
}

func calculatePositions() int {
	count := 0
	for _, y := range grid {
		for _, x := range y {
			if x == "X" {
				count++
			}
		}
	}
	return count
}

//func placeObstacle(x int, y int) {
//	// Don't replace our guard!
//	if grid[y][x] != "^" || grid[y][x] != ">" || grid[y][x] != "v" || grid[y][x] != "<" {
//		grid[y][x] = "#"
//	}
//}

func resetGrid() {
	for i := range grid {
		copy(grid[i], gridOriginal[i])
	}
}

func copyGrid() {
	gridOriginal = make([][]string, len(grid))
	for i := range grid {
		gridOriginal[i] = make([]string, len(grid[i]))
		copy(gridOriginal[i], grid[i])
	}
}
