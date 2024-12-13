package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

var (
	store = make(map[string]int)
)

func keygen(stone int, blinks int) string {
	return strconv.Itoa(stone) + "-" + strconv.Itoa(blinks)
}

func count(stone int, blinks int) int {
	if blinks == 0 {
		return 1
	}

	key := keygen(stone, blinks)
	if val, ok := store[key]; ok {
		return val
	}

	var result int
	switch {
	case stone == 0:
		result = count(1, blinks-1)
	case len(strconv.Itoa(stone))%2 == 0:
		stoned := strconv.Itoa(stone)
		mid := len(stoned) / 2
		left, _ := strconv.Atoi(stoned[:mid])
		right, _ := strconv.Atoi(stoned[mid:])
		result = count(left, blinks-1) + count(right, blinks-1)
	default:
		result = count(stone*2024, blinks-1)
	}

	store[key] = result
	return result
}

func main() {
	data, _ := os.ReadFile("data.txt")
	text := string(data)
	stones := lo.Map(strings.Split(text, " "), func(p string, index int) int {
		n, _ := strconv.Atoi(p)
		return n
	})
	total := lo.SumBy(stones, func(stone int) int {
		return count(stone, 75)
	})

	fmt.Println(total)
}
