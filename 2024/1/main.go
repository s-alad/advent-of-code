package main

import (
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func one(left []int, right []int, length int) {
	sum := 0

	for n := 0; n < length; n++ {
		distance := int(math.Abs(float64(left[n] - right[n])))
		sum += distance
	}

	println(sum)
}

func two(left []int, right []int) {
	sum := 0
	frequencies := map[int]int{}

	for _, r := range right {
		if _, ok := frequencies[r]; ok {
			frequencies[r]++
		} else {
			frequencies[r] = 1
		}
	}

	for _, l := range left {
		if _, ok := frequencies[l]; ok {
			sum += (l * frequencies[l])
		}
	}

	println(sum)
}

func main() {
	data, _ := os.ReadFile("data.txt")
	lines := strings.Split(string(data), "\n")
	length := len(lines)
	left := []int{}
	right := []int{}

	for _, line := range lines {
		parts := strings.Split(line, "   ")
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	one(left, right, length)
	two(left, right)
}
