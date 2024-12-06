package main

import (
	"fmt"
	"os"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

type Coordinate struct {
	x int
	y int
}

type State struct {
	coordinate Coordinate
	direction  string
}

func (p Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func boundry(x, y int, matrix [][]string) bool {
	if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0]) {
		return false
	}
	return true
}

func starting(matrix [][]string) (int, int) {
	for i, row := range matrix {
		for j, char := range row {
			if char == "^" {
				return i, j
			}
		}
	}
	return 0, 0
}

func check(n int, m int, matrix [][]string, potential *Coordinate) bool {
	if boundry(n, m, matrix) {
		if matrix[n][m] == "#" {
			return true
		}
		if potential != nil && (Coordinate{x: n, y: m} == *potential) {
			return true
		}
	}
	return false
}

func one(matrix [][]string) {
	startX, startY := starting(matrix)
	direction := "up"

	for boundry(startX, startY, matrix) {
		switch direction {
		case "up":
			if check(startX-1, startY, matrix, nil) {
				direction = "right"
			} else {
				matrix[startX][startY] = "x"
				startX--
			}
		case "right":
			if check(startX, startY+1, matrix, nil) {
				direction = "down"
			} else {
				matrix[startX][startY] = "x"
				startY++
			}
		case "down":
			if check(startX+1, startY, matrix, nil) {
				direction = "left"
			} else {
				matrix[startX][startY] = "x"
				startX++
			}
		case "left":
			if check(startX, startY-1, matrix, nil) {
				direction = "up"
			} else {
				matrix[startX][startY] = "x"
				startY--
			}
		}
	}

	xs := 0
	for _, row := range matrix {
		for _, char := range row {
			if char == "x" {
				xs++
			}
		}
	}

	fmt.Println(xs)
}

func two(matrix [][]string) {
	potentials := []Coordinate{}

	for i, row := range matrix {
		for j, char := range row {
			if char == "." {
				potentials = append(potentials, Coordinate{x: i, y: j})
			}
		}
	}

	total := 0
	for _, potential := range potentials {
		startX, startY := starting(matrix)
		direction := "up"
		visited := mapset.NewSet[State]()

		for boundry(startX, startY, matrix) {
			current := State{coordinate: Coordinate{x: startX, y: startY}, direction: direction}
			if visited.Contains(current) {
				total++
				break
			}
			visited.Add(current)

			switch direction {
			case "up":
				if check(startX-1, startY, matrix, &potential) {
					direction = "right"
				} else {
					startX--
				}
			case "right":
				if check(startX, startY+1, matrix, &potential) {
					direction = "down"
				} else {
					startY++
				}
			case "down":
				if check(startX+1, startY, matrix, &potential) {
					direction = "left"
				} else {
					startX++
				}
			case "left":
				if check(startX, startY-1, matrix, &potential) {
					direction = "up"
				} else {
					startY--
				}
			}
		}
	}

	fmt.Println(total)
}

func main() {
	data, _ := os.ReadFile("data.txt")
	lines := strings.Split(string(data), "\n")

	m1 := make([][]string, len(lines))
	for i, line := range lines {
		m1[i] = strings.Split(line, "")
	}
	m2 := make([][]string, len(m1))
	for i := range m2 {
		m2[i] = make([]string, len(m1[0]))
		copy(m2[i], m1[i])
	}

	one(m1)
	two(m2)
}
