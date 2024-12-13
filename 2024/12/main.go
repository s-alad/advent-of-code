package main

import (
	"fmt"
	"os"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
)

type Direction string

const (
	Up    Direction = "up"
	Down  Direction = "down"
	Left  Direction = "left"
	Right Direction = "right"
)

type Coordinate struct {
	x, y int
}

type Vector struct {
	coordinate Coordinate
	direction  Direction
}

var directions = []Vector{
	{Coordinate{-1, 0}, Up},
	{Coordinate{1, 0}, Down},
	{Coordinate{0, 1}, Right},
	{Coordinate{0, -1}, Left},
}

func unbounded(c Coordinate, graph [][]string) bool {
	return c.x < 0 || c.y < 0 || c.x >= len(graph) || c.y >= len(graph[0])
}

func boundry(c Coordinate, graph [][]string) bool {
	return c.x >= 0 && c.y >= 0 && c.x < len(graph) && c.y < len(graph[0])
}

func one(graph [][]string) {
	visited := mapset.NewSet[Coordinate]()
	price := 0

	for x := range graph {
		for y := range graph[x] {
			point := Coordinate{x, y}
			if visited.Contains(point) {
				continue
			} else {
				area := 1
				perimeter := 0

				var dfs func(vertex Coordinate, graph [][]string, crop string)
				dfs = func(vertex Coordinate, graph [][]string, crop string) {
					visited.Add(vertex)
					for _, direction := range directions {
						nv := Coordinate{vertex.x + direction.coordinate.x, vertex.y + direction.coordinate.y}
						if boundry(nv, graph) && !visited.Contains(nv) && graph[nv.x][nv.y] == crop {
							area++
							dfs(nv, graph, crop)
						} else if !boundry(nv, graph) || graph[nv.x][nv.y] != crop {
							perimeter++
						}
					}
				}
				crop := graph[x][y]
				dfs(point, graph, crop)
				price += area * perimeter
			}
		}
	}

	fmt.Println(price)
}

func flow(graph [][]string, shape mapset.Set[Coordinate]) int {
	perimeter := 0
	land := make([][]string, len(graph))
	for i := range graph {
		land[i] = make([]string, len(graph[i]))
		for j := range land[i] {
			land[i][j] = "#"
		}
	}
	for _, p := range shape.ToSlice() {
		land[p.x][p.y] = "X"
	}

	for _, direction := range directions {
		switch direction.direction {
		case Up:
			startX := len(land)
			startY := 0

			for start := startX; start > 0; start-- {
				neighbors := mapset.NewSet[Coordinate]()
				for i := startY; i < len(land[0]); i++ {
					if (unbounded(Coordinate{start, i}, land) || land[start][i] == "#") && land[start-1][i] == "X" {
						valid := true
						for _, n := range neighbors.ToSlice() {
							if n.x == start && (n.y == i-1 || n.y == i+1) {
								valid = false
								break
							}
						}

						if valid {
							perimeter++
						}

						neighbors.Add(Coordinate{start, i})
					}
				}
			}

		case Down:
			startX := -1
			startY := 0

			for start := startX; start < len(land)-1; start++ {
				neighbors := mapset.NewSet[Coordinate]()
				for i := startY; i < len(land[0]); i++ {
					if (unbounded(Coordinate{start, i}, land) || land[start][i] == "#") && land[start+1][i] == "X" {
						valid := true
						for _, n := range neighbors.ToSlice() {
							if n.x == start && (n.y == i-1 || n.y == i+1) {
								valid = false
								break
							}
						}

						if valid {
							perimeter++
						}

						neighbors.Add(Coordinate{start, i})
					}
				}
			}
		case Left:
			startY := -1
			startX := 0

			for start := startY; start < len(land[0])-1; start++ {
				neighbors := mapset.NewSet[Coordinate]()
				for i := startX; i < len(land); i++ {
					if (unbounded(Coordinate{i, start}, land) || land[i][start] == "#") && land[i][start+1] == "X" {
						valid := true
						for _, n := range neighbors.ToSlice() {
							if n.y == start && (n.x == i-1 || n.x == i+1) {
								valid = false
								break
							}
						}

						if valid {
							perimeter++
						}

						neighbors.Add(Coordinate{i, start})
					}
				}
			}
		case Right:
			startY := len(land[0])
			startX := 0

			for start := startY; start > 0; start-- {
				neighbors := mapset.NewSet[Coordinate]()
				for i := startX; i < len(land); i++ {
					if (unbounded(Coordinate{i, start}, land) || land[i][start] == "#") && land[i][start-1] == "X" {
						valid := true
						for _, n := range neighbors.ToSlice() {
							if n.y == start && (n.x == i-1 || n.x == i+1) {
								valid = false
								break
							}
						}

						if valid {
							perimeter++
						}

						neighbors.Add(Coordinate{i, start})
					}
				}
			}
		}
	}

	return perimeter
}

func two(graph [][]string) {
	visited := mapset.NewSet[Coordinate]()
	price := 0

	for x := range graph {
		for y := range graph[x] {
			point := Coordinate{x, y}
			if visited.Contains(point) {
				continue
			} else {
				area := 1
				shape := mapset.NewSet[Coordinate]()

				var dfs func(vertex Coordinate, graph [][]string, crop string)
				dfs = func(vertex Coordinate, graph [][]string, crop string) {
					visited.Add(vertex)
					shape.Add(vertex)
					for _, direction := range directions {
						nv := Coordinate{vertex.x + direction.coordinate.x, vertex.y + direction.coordinate.y}
						if boundry(nv, graph) && !visited.Contains(nv) && graph[nv.x][nv.y] == crop {
							area++
							dfs(nv, graph, crop)
						}
					}
				}
				crop := graph[x][y]
				dfs(point, graph, crop)
				perimeter := flow(graph, shape)
				price += area * perimeter
			}
		}
	}

	fmt.Println(price)
}

func main() {
	data, _ := os.ReadFile("data.txt")
	graph := lo.Map(strings.Split(string(data), "\n"), func(line string, _ int) []string { return strings.Split(line, "") })

	/* for _, row := range graph {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	} */

	one(graph)
	two(graph)
}
