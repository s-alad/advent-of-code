package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func slope(p []string, length int) bool {
	increasing := true
	decreasing := true
	for n := 1; n < length; n++ {
		last, _ := strconv.Atoi(p[n-1])
		current, _ := strconv.Atoi(p[n])
		if current > last {
			decreasing = false
		} else {
			increasing = false
		}
	}

	return increasing || decreasing
}

func compare(p []string) bool {
	length := len(p)
	for n := 1; n < length; n++ {
		last, _ := strconv.Atoi(p[n-1])
		current, _ := strconv.Atoi(p[n])
		if !(int(math.Abs(float64(current-last))) >= 1 && int(math.Abs(float64(current-last))) <= 3) {
			return false
		}
	}

	return true
}

func one(lines []string) {
	safe := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if compare(parts) && slope(parts, len(parts)) {
			safe++
		}
	}

	fmt.Println(safe)
}

func two(lines []string) {
	safe := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		length := len(parts)

		permutations := [][]string{parts}
		for i := 0; i < length; i++ {
			permutation := make([]string, length-1)
			copy(permutation, parts[:i])
			copy(permutation[i:], parts[i+1:])
			permutations = append(permutations, permutation)
		}

		for _, permutation := range permutations {
			if compare(permutation) && slope(permutation, len(permutation)) {
				safe++
				break
			}
		}
	}

	fmt.Println(safe)
}

func main() {
	data, _ := os.ReadFile("data.txt")
	lines := strings.Split(string(data), "\n")

	one(lines)
	two(lines)
}
