package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func one(memory []string) {
	left := 0
	right := len(memory) - 1

	for left < right {
		if memory[left] != "." {
			left++
			continue
		}
		if memory[right] == "." {
			right--
			continue
		}

		memory[left] = memory[right]
		memory[right] = "."

		left++
		right--
	}

	total := 0
	for i, m := range memory {
		m, _ := strconv.Atoi(m)
		total += m * i
	}

	fmt.Println(total)
}

func delete(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
func two(memory []string) {

	right := len(memory) - 1

	cache := make([]string, len(memory))
	copy(cache, memory)

	for right > 0 {
		if strings.Contains(cache[right], ".") {
			right--
			continue
		}

		for i := 0; i < right; i++ {
			amount := (strings.Count(cache[right], "-") + 1)
			if strings.Contains(cache[i], ".") && len(cache[i]) >= amount {
				temp := cache[right]

				if right-1 >= 0 && strings.Contains(cache[right-1], ".") &&
					right+1 < len(cache) && strings.Contains(cache[right+1], ".") {
					newdots := strings.Repeat(".", len(cache[right-1])+len(cache[right+1])+amount)
					cache = delete(cache, right+1)
					cache = delete(cache, right)
					cache[right-1] = newdots

				} else if right-1 >= 0 && strings.Contains(cache[right-1], ".") {
					cache[right-1] = cache[right-1] + strings.Repeat(".", amount)
					cache = delete(cache, right)
				} else if right+1 < len(cache) && strings.Contains(cache[right+1], ".") {
					cache[right+1] = cache[right+1] + strings.Repeat(".", amount)
					cache = delete(cache, right)
				} else {
					cache[right] = strings.Repeat(".", amount)
				}

				leftover := len(cache[i]) - amount
				if right-i == 1 {
					total := strings.Repeat(".", leftover)
					cache[i] = temp
					cache = slices.Insert(cache, i+1, total)
				} else {

					cache[i] = temp

					if leftover > 0 {
						cache = slices.Insert(cache, i+1, strings.Repeat(".", leftover))
						right++
					}
				}

				break
			}
		}

		right--
	}

	id := 0
	total := 0
	for _, c := range cache {
		var splits []string
		if strings.Contains(c, "-") {
			splits = strings.Split(c, "-")
		} else if strings.Contains(c, ".") {
			splits = strings.Split(c, "")
		} else {
			splits = []string{c}
		}
		for _, s := range splits {
			if s != "." {
				digit, _ := strconv.Atoi(s)
				total += digit * id
			}
			id++
		}

	}

	//fmt.Println(cache)
	fmt.Println(total)
}

func main() {
	data, _ := os.ReadFile("data.txt")
	text := string(data)
	//fmt.Println(text)

	iter := 0
	m1 := []string{}
	m2 := []string{}
	for index, c := range text {
		r := lo.Ternary(index%2 == 0, strconv.Itoa(iter), ".")
		for i := 0; i < int(c-'0'); i++ {
			m1 = append(m1, r)
		}
		ns := strings.Repeat(r+"-", int(c-'0'))
		co := lo.Ternary(strings.LastIndex(ns, "-") != -1, strings.LastIndex(ns, "-"), 0)
		entry := lo.Ternary(index%2 == 0, ns[:co], strings.Repeat(".", int(c-'0')))
		if entry != "" {
			m2 = append(m2, entry)
		}
		iter += lo.Ternary(index%2 == 0, 1, 0)
	}

	// merge dots
	for i := 0; i < len(m2)-1; i++ {
		if m2[i] == "." && m2[i+1] == "." {
			m2[i] = strings.Repeat(".", len(m2[i])+len(m2[i+1]))
			m2 = append(m2[:i+1], m2[i+2:]...)
		}
	}

	//fmt.Println(m1)
	//fmt.Println(m2)

	one(m1)
	two(m2)
}
