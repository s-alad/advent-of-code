package main

import (
	"fmt"
	"os"
	"strings"
)

func one(matrix [][]string) {
	total := 0
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, 1}, {1, -1}, {-1, -1}}

	for i, row := range matrix {
		for j, cell := range row {
			if cell == "X" {
				for _, dir := range directions {
					dx, dy := dir[0], dir[1]
					if check(matrix, i, j, dx, dy) {
						total++
					}
				}
			}
		}
	}

	fmt.Println(total)
}

func check(matrix [][]string, x, y, dx, dy int) bool {
	for k, letter := range []string{"M", "A", "S"} {
		x, y := x+(k+1)*dx, y+(k+1)*dy
		if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0]) || matrix[x][y] != letter {
			return false
		}
	}
	return true
}

func two(matrix [][]string) {
	total := 0
	directions := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	for i, row := range matrix {
		for j, cell := range row {
			if cell == "A" {
				M, S := 0, 0

				for _, dir := range directions {
					x, y := i+dir[0], j+dir[1]
					if x >= 0 && y >= 0 && x < len(matrix) && y < len(matrix[0]) {
						if matrix[x][y] == "M" {
							M++
						} else if matrix[x][y] == "S" {
							S++
						}
					}
				}

				if M == 2 && S == 2 && matrix[i-1][j-1] != matrix[i+1][j+1] && matrix[i-1][j+1] != matrix[i+1][j-1] {
					total++
				}
			}
		}
	}

	fmt.Println(total)
}

func main() {
	data, _ := os.ReadFile("data.txt")
	lines := strings.Split(string(data), "\n")
	matrix := make([][]string, len(lines))
	for i, line := range lines {
		matrix[i] = strings.Split(line, "")
	}

	one(matrix)
	two(matrix)
}
