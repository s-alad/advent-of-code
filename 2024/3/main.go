package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func one(s string) {
	matches := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`).FindAllStringSubmatch(s, -1)
	total := 0

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		total += x * y
	}

	fmt.Println(total)
}

func two(s string) {
	matches := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`).FindAllStringSubmatch(s, -1)
	total := 0
	enabled := true

	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			total += x * y
		}
	}

	fmt.Println(total)
}

func main() {
	data, _ := os.ReadFile("data.txt")
	stringdata := string(data)

	one(stringdata)
	two(stringdata)
}
