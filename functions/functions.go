package functions

import (
	"strconv"
)

// AbsInt Return absolute value
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func StringToInts(s []string) []int {
	ints := make([]int, len(s))

	for i, s := range s {
		num, _ := strconv.Atoi(s)
		ints[i] = num
	}
	return ints
}
