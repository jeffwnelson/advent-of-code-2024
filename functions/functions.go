package functions

import (
	"strconv"
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
