package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func boundry(x, y int, matrix [][]int) bool {
	if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) {
		return false
	}
	return true
}

func keygen(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

var directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func recur(x, y int, matrix [][]int, n int, visited map[string]bool) int {
	if n == 9 {
		if visited == nil {
			return 1
		}

		key := keygen(x, y)
		if _, ok := visited[key]; ok {
			return 0
		} else {
			visited[key] = true
			return 1
		}
	}

	running := 0
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if boundry(nx, ny, matrix) && matrix[nx][ny] == n+1 {
			running += recur(nx, ny, matrix, n+1, visited)
		}
	}

	return running
}

func main() {
	data, _ := os.ReadFile("data.txt")
	text := strings.Split(string(data), "\n")
	matrix := [][]string{}
	for _, line := range text {
		matrix = append(matrix, strings.Split(line, ""))
	}

	cache := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			x, _ := strconv.Atoi(matrix[i][j])
			cache[i] = append(cache[i], x)
		}
	}

	one := 0
	two := 0
	for x, row := range cache {
		for y, c := range row {
			if c == 0 {
				visited := map[string]bool{}
				one += recur(x, y, cache, 0, visited)
				two += recur(x, y, cache, 0, nil)
			}
		}
	}

	fmt.Println(one)
	fmt.Println(two)
}
