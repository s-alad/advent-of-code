package main

import (
	"fmt"
	"os"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
)

type Coordinate struct {
	x, y int
}

type Pair struct {
	c1, c2 Coordinate
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(c1, c2 Coordinate) int {
	return abs(c1.x-c2.x) + abs(c1.y-c2.y)
}

func slope(c1, c2 Coordinate) (int, int) {
	// return the rise and the run between two coordinates
	return c2.x - c1.x, c2.y - c1.y
}

func bound(c Coordinate, graph [][]string) bool {
	return c.x >= 0 && c.y >= 0 && c.x < len(graph) && c.y < len(graph[0])
}

func one(nodes map[string][]Coordinate, graph [][]string) {
	total := 0
	antis := mapset.NewSet[Coordinate]()
	for _, coords := range nodes {
		for c1 := range coords {
			for c2 := c1; c2 < len(coords); c2++ {
				if c1 == c2 {
					continue
				}

				sx, sy := slope(coords[c1], coords[c2])

				v1 := Coordinate{coords[c1].x - sx, coords[c1].y - sy}
				v2 := Coordinate{coords[c2].x + sx, coords[c2].y + sy}

				if bound(v1, graph) {
					total++
					antis.Add(v1)
				}
				if bound(v2, graph) {
					total++
					antis.Add(v2)
				}
			}
		}
	}

	fmt.Println(antis.Cardinality())
}

func two(nodes map[string][]Coordinate, graph [][]string) {
	total := 0
	antis := mapset.NewSet[Coordinate]()
	for _, coords := range nodes {
		for c1 := range coords {
			for c2 := c1; c2 < len(coords); c2++ {
				if c1 == c2 {
					continue
				}

				sx, sy := slope(coords[c1], coords[c2])

				bounded1, bounded2 := false, false
				multiplier := 1

				antis.Add(coords[c1])
				antis.Add(coords[c2])

				for !bounded1 || !bounded2 {
					v1 := Coordinate{coords[c1].x - sx*multiplier, coords[c1].y - sy*multiplier}
					v2 := Coordinate{coords[c2].x + sx*multiplier, coords[c2].y + sy*multiplier}

					// /fmt.Println(v1, v2)

					if bound(v1, graph) {
						total++
						antis.Add(v1)
					} else {
						bounded1 = true
					}
					if bound(v2, graph) {
						total++
						antis.Add(v2)
					} else {
						bounded2 = true
					}

					multiplier++
				}
			}
		}
	}

	fmt.Println(antis.Cardinality())

}

func main() {
	data, _ := os.ReadFile("data.txt")
	graph := lo.Map(strings.Split(string(data), "\n"), func(line string, _ int) []string { return strings.Split(line, "") })
	nodes := map[string][]Coordinate{}
	for x, row := range graph {
		for y, cell := range row {
			if cell != "." {
				nodes[cell] = append(nodes[cell], Coordinate{x, y})
			}
		}
	}

	fmt.Println(nodes)
	//one(nodes, graph)
	two(nodes, graph)
}
