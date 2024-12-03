package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var file = "day02/input.txt"

func TestPart1(t *testing.T) {
	start := time.Now() // Record start time

	input, _ := os.Open(file)
	defer input.Close()
	part1(input)

	duration := time.Since(start) // Calculate the duration
	fmt.Printf("Execution time: %d µs\n", duration.Microseconds())
}

func TestPart2(t *testing.T) {
	start := time.Now() // Record start time

	input, _ := os.Open(file)
	defer input.Close()
	part2(input)

	duration := time.Since(start) // Calculate the duration
	fmt.Printf("Execution time: %d µs\n", duration.Microseconds())
}
