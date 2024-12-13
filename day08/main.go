package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readFileAsMatrix("day08/input.txt")

	fmt.Println("Part 1 solution: ", part1(input))
	// Attempt #1: 380 (too high)
	// Attempt #2 was correct. Forgot that my grid was 0 through 49, not 0 through 50.
	// So several antinode locations were being counted when they should not have been.

	fmt.Println("Part 2 solution: ", part2(input))
}

func readFileAsMatrix(filename string) [][]rune {
	file, _ := os.Open(filename)
	defer file.Close()
	var matrix [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		matrix = append(matrix, row)
	}

	return matrix
}

func part1(file [][]rune) int {
	charMap := createMap(file)
	antinodes := createAntinodeMatrix(len(file), len(file[0]))
	result := 0
	for _, val := range charMap {
		result += placeUniqueAntinodes(&val, &antinodes)
	}
	return result
}

func part2(file [][]rune) int {
	charMap := createMap(file)
	antinodes := createAntinodeMatrix(len(file), len(file[0]))
	result := 0
	for _, val := range charMap {
		for _, node := range val {
			antinodes[node.row][node.col] = true
			result++
		}
	}
	for _, val := range charMap {
		result += placeAntinodeLines(&val, &antinodes)
	}
	return result
}

func createMap(file [][]rune) map[rune][]Coordinates {
	result := make(map[rune][]Coordinates)
	for i, line := range file {
		for j, char := range line {
			if char != '.' {
				result[char] = append(result[char], Coordinates{row: i, col: j})
			}
		}
	}
	return result
}

func createAntinodeMatrix(rows int, cols int) [][]bool {
	var result [][]bool
	for i := 0; i < rows; i++ {
		result = append(result, make([]bool, cols))
	}
	return result
}

func placeUniqueAntinodes(places *[]Coordinates, antinodes *[][]bool) int {
	size_r := len(*antinodes)
	size_c := len((*antinodes)[0])
	result := 0
	for i := 0; i < len(*places)-1; i++ {
		for j := i + 1; j < len(*places); j++ {
			coordinate_1 := calcAntinode((*places)[i], (*places)[j])
			if isNewResult(coordinate_1, size_r, size_c, antinodes) {
				result++
			}
			coordinate_2 := calcAntinode((*places)[j], (*places)[i])
			if isNewResult(coordinate_2, size_r, size_c, antinodes) {
				result++
			}
		}
	}
	return result
}

func calcAntinode(antenna1 Coordinates, antenna2 Coordinates) Coordinates {
	row_dif := antenna1.row - antenna2.row
	col_dif := antenna1.col - antenna2.col
	return Coordinates{row: antenna1.row + row_dif, col: antenna1.col + col_dif}
}

func isNewResult(coordinate Coordinates, size_r int, size_c int, antinodes *[][]bool) bool {
	if coordinate.row < 0 || coordinate.col < 0 {
		return false
	}
	if coordinate.row >= size_r || coordinate.col >= size_c {
		return false
	}
	if (*antinodes)[coordinate.row][coordinate.col] {
		return false
	}
	(*antinodes)[coordinate.row][coordinate.col] = true
	return true
}

func placeAntinodeLines(places *[]Coordinates, antinodes *[][]bool) int {
	result := 0
	for i := 0; i < len(*places)-1; i++ {
		for j := i + 1; j < len(*places); j++ {
			result += calcAntinodeLines((*places)[i], (*places)[j], antinodes)
			result += calcAntinodeLines((*places)[j], (*places)[i], antinodes)
		}
	}
	return result
}

func calcAntinodeLines(antenna1 Coordinates, antenna2 Coordinates, antinodes *[][]bool) int {
	row_dif := antenna1.row - antenna2.row
	col_dif := antenna1.col - antenna2.col
	curr_row := antenna1.row
	curr_col := antenna1.col
	uniques := 0
	size_r := len((*antinodes))
	size_c := len((*antinodes)[0])
	for validPositon(curr_row+row_dif, curr_col+col_dif, size_r, size_c) {
		curr_row += row_dif
		curr_col += col_dif
		if !(*antinodes)[curr_row][curr_col] {
			uniques++
			(*antinodes)[curr_row][curr_col] = true
		}
	}
	return uniques
}

func validPositon(new_row int, new_col int, size_r int, size_c int) bool {
	if new_row < 0 || new_col < 0 {
		return false
	}
	if new_row >= size_r || new_col >= size_c {
		return false
	}
	return true
}

type Coordinates struct {
	row int
	col int
}
