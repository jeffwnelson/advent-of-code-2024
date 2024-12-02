package main

import (
	"os"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	input, _ := os.Open("day02/input.txt")
	defer input.Close()

	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	input, _ := os.Open("day02/input.txt")
	defer input.Close()

	for i := 0; i < b.N; i++ {
		part2(input)
	}
}
