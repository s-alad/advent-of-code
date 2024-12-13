package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Equation struct {
	a     int
	b     int
	total int
}

type System struct {
	eq1 Equation
	eq2 Equation
}

func solve(system System) (int, int) {
	a1, b1, c1 := system.eq1.a, system.eq1.b, system.eq1.total
	a2, b2, c2 := system.eq2.a, system.eq2.b, system.eq2.total

	det := a1*b2 - a2*b1
	if det == 0 {
		return -1, -1
	}

	x := (c1*b2 - c2*b1) / det
	y := (a1*c2 - a2*c1) / det

	if a1*x+b1*y != c1 || a2*x+b2*y != c2 {
		return -1, -1
	}

	return x, y
}

func compute(systems []System) {
	var total int
	for _, system := range systems {
		x, y := solve(system)
		if x == -1 || y == -1 {
			continue
		}

		total += (x * 3) + y
	}

	fmt.Println(total)
}

func main() {
	data, _ := os.ReadFile("data.txt")
	graph := lo.Map(strings.Split(string(data), "\n\n"), func(line string, _ int) []string { return strings.Split(line, "\n") })

	var systems []System
	for _, group := range graph {
		re := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
		ma := re.FindStringSubmatch(group[0])
		mb := re.FindStringSubmatch(group[1])

		rp := regexp.MustCompile(`X=(\d+), Y=(\d+)`)
		pr := rp.FindStringSubmatch(group[2])

		system := System{
			eq1: Equation{
				a:     func() int { v, _ := strconv.Atoi(ma[1]); return v }(),
				b:     func() int { v, _ := strconv.Atoi(mb[1]); return v }(),
				total: func() int { v, _ := strconv.Atoi(pr[1]); return v }() + 10000000000000,
			},
			eq2: Equation{
				a:     func() int { v, _ := strconv.Atoi(ma[2]); return v }(),
				b:     func() int { v, _ := strconv.Atoi(mb[2]); return v }(),
				total: func() int { v, _ := strconv.Atoi(pr[2]); return v }() + 10000000000000,
			},
		}

		systems = append(systems, system)
	}

	compute(systems)

}
