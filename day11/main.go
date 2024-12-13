package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("day11/input.txt")
	input := string(inputBytes)
	fmt.Println("Part 1 solution: ", part1(input))
	fmt.Println("Part 2 solution: ", part2(input))
}

func part1(input string) int {
	stones := strings.Fields(input)
	newStones := make(map[int]int)

	for _, stone := range stones {
		value, _ := strconv.Atoi(stone)
		newStones[value]++
	}
	// fmt.Println(newStones)
	// map[0:1 1:1 22:1 528:1 2790:1 10725:1 572556:1 4679021:1]

	for i := 0; i < 25; i++ {
		newStones = solve(newStones)
	}

	return totalStoneCount(newStones)
}

func part2(input string) int {
	// Learning about memoization
	// https://fullstackdojo.medium.com/memoization-in-golang-a946acd10829
	stones := strings.Fields(input)
	newStones := make(map[int]int)

	for _, stone := range stones {
		value, _ := strconv.Atoi(stone)
		newStones[value]++
	}

	// fmt.Println(newStones)
	// map[17:1 125:1]
	for i := 0; i < 75; i++ {
		newStones = solve(newStones)
		//fmt.Println(newStones)
		//map[1:1 7:1 253000:1]
		//map[0:1 253:1 2024:1 14168:1]
		//map[1:1 20:1 24:1 512072:1 28676032:1]
		//map[0:1 2:2 4:1 72:1 512:1 2024:1 2867:1 6032:1]
		//map[1:1 2:1 7:1 20:1 24:1 28:1 32:1 60:1 67:1 4048:2 8096:1 1036288:1]
		//map[0:2 2:4 3:1 4:1 6:2 7:1 8:1 40:2 48:2 80:1 96:1 2024:1 4048:1 14168:1 2097446912:1]
	}

	return totalStoneCount(newStones)
}

func totalStoneCount(stoneCount map[int]int) int {
	count := 0
	for _, value := range stoneCount {
		count += value
		//fmt.Printf("Count / Value [%v] [%v,%v]\n", i, count, value)
	}
	return count
}

func solve(stoneCount map[int]int) map[int]int {
	updatedStoneCount := make(map[int]int)

	//stoneCounts coming in as "map[int]int"
	//map[1:1 7:1 253000:1]
	//map[0:1 253:1 2024:1 14168:1]
	//map[1:1 20:1 24:1 512072:1 28676032:1]
	//map[0:1 2:2 4:1 72:1 512:1 2024:1 2867:1 6032:1]
	//map[1:1 2:1 7:1 20:1 24:1 28:1 32:1 60:1 67:1 4048:2 8096:1 1036288:1]
	//map[0:2 2:4 3:1 4:1 6:2 7:1 8:1 40:2 48:2 80:1 96:1 2024:1 4048:1 14168:1 2097446912:1]
	for stone, count := range stoneCount {
		stones := blink(stone) // blink the stone
		for _, value := range stones {
			//fmt.Printf("[%v] Blinking stone %v becomes %v\n", i, stone, value)
			updatedStoneCount[value] += count
		}
		//fmt.Println("")
	}
	return updatedStoneCount
}

func blink(value int) []int {
	sStone := strconv.Itoa(value)

	// Checking conditions
	switch {
	// If our value is 0, return 1
	case value == 0:
		return []int{1}

	// If our value is event, split left and right sides
	case len(sStone)%2 == 0:
		left, _ := strconv.Atoi(sStone[:len(sStone)/2])
		right, _ := strconv.Atoi(sStone[len(sStone)/2:])
		return []int{left, right}

	// If the other cases didn't match, then our default is to multiply value by 2024
	default:
		return []int{value * 2024}
	}
}
