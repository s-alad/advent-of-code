package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func one(rules map[int][]int, pages [][]int) [][]int {
	acceptable := [][]int{}
	unacceptable := [][]int{}

	for _, page := range pages {
		valid := true
		for i, num := range page {
			rule := rules[num]

			before := page[:i]
			for _, n := range before {
				if slices.Contains(rule, n) {
					valid = false
					break
				}
			}
		}

		if valid {
			acceptable = append(acceptable, page)
		} else {
			unacceptable = append(unacceptable, page)
		}
	}

	//fmt.Println(acceptable)
	//fmt.Println(unacceptable)

	total := 0
	for _, page := range acceptable {
		total += page[len(page)/2]
	}

	fmt.Println(total)

	return unacceptable
}

func two(rules map[int][]int, pages [][]int) {
	for _, page := range pages {
		sort.Slice(page, func(i, j int) bool {
			return !slices.Contains(rules[page[j]], page[i])
		})
	}

	total := 0
	for _, page := range pages {
		total += page[len(page)/2]
	}

	fmt.Println(total)
}

func main() {
	data, _ := os.ReadFile("data.txt")
	lines := strings.Split(string(data), "\n\n")

	rules := make(map[int][]int)
	pages := [][]int{}

	for _, rule := range strings.Split(lines[0], "\n") {
		parts := strings.Split(rule, "|")
		prior, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])
		rules[prior] = append(rules[prior], after)
	}

	//fmt.Println(rules)

	for _, page := range strings.Split(lines[1], "\n") {
		nums := []int{}
		for _, num := range strings.Split(page, ",") {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}
		pages = append(pages, nums)
	}

	//fmt.Println(pages)

	two(rules, one(rules, pages))
}
