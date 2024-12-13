package functions

import (
	"strconv"
	"strings"
)

// AbsInt Converts a negative int to a positive int for absolute value needs
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// StringToInts Takes an array of strings and converts to an array of ints
func StringToInts(s []string) []int {
	ints := make([]int, len(s))

	for i, s := range s {
		num, _ := strconv.Atoi(s)
		ints[i] = num
	}
	return ints
}

func StringContains(str string, slice []string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func GetStringPosition(string string, slice []string) int {
	for i, v := range slice {
		if v == string {
			return i
		}
	}
	return -1 // Return -1 if the string is not found
}

func StringToIntsSlice(s string) []int {
	strs := strings.Split(strings.TrimSpace(s), " ")
	array := make([]int, len(strs))
	for i := range array {
		array[i], _ = strconv.Atoi(strs[i])
	}
	return array
}
