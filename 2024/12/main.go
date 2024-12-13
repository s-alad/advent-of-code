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

func two(graph [][]string) {
	visited := mapset.NewSet[Coordinate]()
	price := 0

	for x := range graph {
		for y := range graph[x] {
			point := Coordinate{x, y}
			if visited.Contains(point) {
				continue
			} else {
				area := 0
				perimeter := 0
				vectors := mapset.NewSet[Vector]()

				bfs := func(start Coordinate, graph [][]string, crop string) {
					queue := []Coordinate{start}
					visited.Add(start)

					for len(queue) > 0 {
						current := queue[0]
						queue = queue[1:]
						area++

						for _, direction := range directions {
							nv := Coordinate{current.x + direction.coordinate.x, current.y + direction.coordinate.y}
							if boundry(nv, graph) && !visited.Contains(nv) && graph[nv.x][nv.y] == crop {
								queue = append(queue, nv)
								visited.Add(nv)
							} else if !boundry(nv, graph) || graph[nv.x][nv.y] != crop {
								orientation := direction.direction
								valid := true
								switch orientation {
								case Up:
									for _, v := range vectors.ToSlice() {
										if v.direction == Up &&
											(v.coordinate.y == nv.y-1 || v.coordinate.y == nv.y+1) &&
											(v.coordinate.x == nv.x) {
											valid = false
										}
									}
								case Down:
									for _, v := range vectors.ToSlice() {
										if v.direction == Down &&
											(v.coordinate.y == nv.y-1 || v.coordinate.y == nv.y+1) &&
											(v.coordinate.x == nv.x) {
											valid = false
										}
									}
								case Left:
									for _, v := range vectors.ToSlice() {
										if v.direction == Left &&
											(v.coordinate.x == nv.x-1 || v.coordinate.x == nv.x+1) &&
											(v.coordinate.y == nv.y) {
											valid = false
										}
									}
								case Right:
									for _, v := range vectors.ToSlice() {
										if v.direction == Right &&
											(v.coordinate.x == nv.x-1 || v.coordinate.x == nv.x+1) &&
											(v.coordinate.y == nv.y) {
											valid = false
										}
									}
								}

								if valid {
									perimeter++
								}

								vectors.Add(Vector{
									Coordinate{nv.x, nv.y},
									orientation,
								})
							}
						}
					}
				}

				crop := graph[x][y]
				//fmt.Println("starting with", crop)
				bfs(point, graph, crop)
				//fmt.Printf("ended with %d area and %d perimeter for %s\n", area, perimeter, crop)
				price += area * perimeter
			}
		}
	}

	fmt.Println(price)
}

func main() {
	data, _ := os.ReadFile("data.txt")
	graph := lo.Map(strings.Split(string(data), "\n"), func(line string, _ int) []string { return strings.Split(line, "") })

	for _, row := range graph {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}

	//one(graph)
	two(graph)
}
