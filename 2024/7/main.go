package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func operate(matrix []map[int][]int, operations []string) {
	acceptable := []int{}

	for _, row := range matrix {
		for start, numbers := range row {
			n := len(numbers) - 1

			combinations := int(math.Pow(float64(len(operations)), float64(n)))
			for combination := 0; combination < combinations; combination++ {
				ordering := []string{}
				temp := combination
				for j := 0; j < len(numbers)-1; j++ {
					ordering = append(ordering, operations[temp%len(operations)])
					temp /= len(operations)
				}

				result := numbers[0]
				for k, op := range ordering {
					switch op {
					case "+":
						result += numbers[k+1]
					case "*":
						result *= numbers[k+1]
					case "||":
						x, _ := strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(numbers[k+1]))
						result = x
					}
				}

				if result == start {
					acceptable = append(acceptable, start)
					break
				}
			}
		}
	}

	total := 0
	for _, number := range acceptable {
		total += number
	}

	fmt.Println(total)
}

func one(matrix []map[int][]int) {
	operate(matrix, []string{"+", "*"})
}

func two(matrix []map[int][]int) {
	operate(matrix, []string{"+", "*", "||"})
}

func main() {
	data, _ := os.ReadFile("data.txt")
	lines := strings.Split(string(data), "\n")
	matrix := make([]map[int][]int, len(lines))

	for i, line := range lines {
		matrix[i] = make(map[int][]int)
		values := strings.Split(line, ": ")
		start, _ := strconv.Atoi(values[0])
		posterior := strings.Split(values[1], " ")
		numbers := make([]int, len(posterior))
		for j, number := range posterior {
			numbers[j], _ = strconv.Atoi(number)
		}

		matrix[i][start] = numbers
	}

	one(matrix)
	two(matrix)
}
