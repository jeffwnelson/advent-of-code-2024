package main

import (
	"fmt"
	"github.com/jeffwnelson/advent-of-code-2024/functions"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputBytes, _ := os.ReadFile("day05/test.txt")
	input := string(inputBytes)
	fmt.Println("Part 1 solution: ", part1(input))
	fmt.Println("Part 2 solution: ", part2(input))
}

func part1(input string) int {
	sections := strings.Split(input, "\n\n")
	rules := strings.Split(sections[0], "\n")
	pages := strings.Split(sections[1], "\n")

	total := 0
	for i, _ := range pages {
		check, value := checkRules(pages[i], rules)
		if check {
			total += value
		}
	}
	return total
}

func part2(input string) int {
	sections := strings.Split(input, "\n\n")
	rules := strings.Split(sections[0], "\n")
	pages := strings.Split(sections[1], "\n")

	total := 0
	for i, _ := range pages {
		check, value := fixedRules(pages[i], rules)
		if check {
			total += value
		}
	}
	return total
}

func checkRules(pages string, rules []string) (bool, int) {
	pageArray := strings.Split(pages, ",")

	check := false
	middleValue := 0
	for _, value := range rules {
		ruleArray := strings.Split(value, "|")
		x := ruleArray[0]
		y := ruleArray[1]

		if functions.StringContains(x, pageArray) && functions.StringContains(y, pageArray) {
			xPos := functions.GetStringPosition(x, pageArray)
			yPos := functions.GetStringPosition(y, pageArray)

			if xPos != -1 && yPos != -1 && xPos < yPos {
				check = true
			} else {
				return false, 0
			}
		}
	}

	if check {
		middleValue = getMiddleValue(pageArray)
	}

	return check, middleValue
}

func getMiddleValue(string []string) int {
	middleValue, _ := strconv.Atoi(string[len(string)/2])
	return middleValue
}

func fixedRules(pages string, rules []string) (bool, int) {
	check := true
	for check {
		pageArray := strings.Split(pages, ",")

		for i, value := range rules {
			ruleArray := strings.Split(value, "|")
			x := ruleArray[0]
			y := ruleArray[1]

			if functions.StringContains(x, pageArray) && functions.StringContains(y, pageArray) {
				xPos := functions.GetStringPosition(x, pageArray)
				yPos := functions.GetStringPosition(y, pageArray)
				if xPos != -1 && yPos != -1 && i < len(rules)-1 {
					rules[i] = rules[i+1]
					rules[i+1] = rules[i]
				} else {
					check = false
				}

				check = false
			}
		}
	}
	return checkRules(pages, rules)
}
