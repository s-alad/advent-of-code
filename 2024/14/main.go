package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

var (
	HEIGHT = 103
	WIDTH  = 101
)

type Vector struct {
	x int
	y int
}

type Robot struct {
	position Vector
	velocity Vector
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printer(grid [][]string) {
	for _, g := range grid {
		fmt.Println(strings.Join(g, ""))
	}
	fmt.Println()
}

func robotprinter(robots []Robot, grid [][]string) {
	for _, robot := range robots {
		grid[robot.position.y][robot.position.x] = "#"
	}
	printer(grid)
	for _, robot := range robots {
		grid[robot.position.y][robot.position.x] = "."
	}
}

func one(robots []Robot, grid [][]string) {

	for i, robot := range robots {
		for seconds := 0; seconds < 100; seconds++ {
			potentialX := robot.position.x + robot.velocity.x
			potentialY := robot.position.y + robot.velocity.y
			potentialX = (potentialX + WIDTH) % WIDTH
			potentialY = (potentialY + HEIGHT) % HEIGHT
			robot.position = Vector{potentialX, potentialY}
		}
		robots[i] = robot
	}

	quadrants := make([]int, 4)
	mx := WIDTH / 2
	my := HEIGHT / 2

	for _, robot := range robots {
		x, y := robot.position.x, robot.position.y

		if x == mx || y == my {
			continue
		}

		quadrant := lo.Ternary(y < my, lo.Ternary(x < mx, 0, 1), lo.Ternary(x < mx, 2, 3))
		quadrants[quadrant]++
	}

	total := lo.Reduce(quadrants, func(acc, val, _ int) int { return acc * val }, 1)
	fmt.Println(total)
}

func heuristic(robots []Robot) int {
	score := 0
	positions := make(map[Vector]int)

	for _, robot := range robots {
		positions[robot.position]++
	}

	for i := 0; i < len(robots); i++ {
		for j := i + 1; j < len(robots); j++ {
			dx := abs(robots[i].position.x - robots[j].position.x)
			dy := abs(robots[i].position.y - robots[j].position.y)
			score += dx + dy
		}
	}

	for _, count := range positions {
		if count > 1 {
			score += (count * count) * 10
		}
	}

	return score
}

func two(robots []Robot, grid [][]string) {
	minimum := 1000000000
	best := 0

	for seconds := 0; seconds < 10000; seconds++ {
		for i, robot := range robots {
			potentialX := robot.position.x + robot.velocity.x
			potentialY := robot.position.y + robot.velocity.y
			potentialX = (potentialX + WIDTH) % WIDTH
			potentialY = (potentialY + HEIGHT) % HEIGHT
			robot.position = Vector{potentialX, potentialY}
			robots[i] = robot
		}

		score := heuristic(robots)
		if score < minimum {
			minimum = score
			best = seconds
			robotprinter(robots, grid)
		}
	}

	fmt.Println(best)
}

func main() {
	raw, _ := os.ReadFile("data.txt")
	data := lo.Map(strings.Split(string(raw), "\n"), func(line string, _ int) string { return line })

	var robots []Robot
	for _, row := range data {
		matches := regexp.
			MustCompile(`p=([+-]?\d+),([+-]?\d+)\s+v=([+-]?\d+),([+-]?\d+)`).
			FindStringSubmatch(row)
		px, _ := strconv.Atoi(matches[1])
		py, _ := strconv.Atoi(matches[2])
		vx, _ := strconv.Atoi(matches[3])
		vy, _ := strconv.Atoi(matches[4])

		robots = append(robots, Robot{Vector{px, py}, Vector{vx, vy}})
	}

	grid := make([][]string, HEIGHT)
	for y := range grid {
		grid[y] = make([]string, WIDTH)
		for x := range grid[y] {
			grid[y][x] = "."
		}
	}

	// one(robots, grid)
	// two(robots, grid)
}
