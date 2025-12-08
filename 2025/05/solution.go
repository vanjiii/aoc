package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/vanjiii/aoc/pkg"
)

func main() {
	rows, err := pkg.ReadStrings("input")
	if err != nil {
		log.Fatalf("fail to read input: %v", err)
		return
	}

	fmt.Printf("part one: %v\n", partone(rows))
	fmt.Printf("part two: %v\n", parttwo(rows))
}

func parttwo(rows []string) int {
	freshIng := [][]int{}

	res := 0
	for _, row := range rows {
		if strings.Contains(row, "-") {
			ranges := strings.Split(row, "-")
			a, _ := strconv.Atoi(ranges[0])
			b, _ := strconv.Atoi(ranges[1])

			freshIng = append(freshIng, []int{a, b})
		}
	}

	smushed := [][]int{}

	sort.Slice(freshIng, func(i, j int) bool {
		if freshIng[i][0] != freshIng[j][0] {
			return freshIng[i][0] < freshIng[j][0]
		}
		return freshIng[i][1] < freshIng[j][1]
	})

	for _, tuple := range freshIng {
		a, b := tuple[0], tuple[1]

		for _, r := range smushed {
			x, y := r[0], r[1]

			if a >= x && a <= y {
				a = y + 1
			}

			if b >= x && b <= y {
				b = x - 1
			}
		}

		if a <= b {
			smushed = append(smushed, []int{a, b})
		}
	}

	for _, tuple := range smushed {
		// fmt.Println(tuple)
		a, b := tuple[0], tuple[1]

		res += b - a + 1
	}

	return res
}

func partone(rows []string) int {
	freshIng := [][]int{}

	res := 0
	for _, row := range rows {
		if strings.Contains(row, "-") {
			ranges := strings.Split(row, "-")
			a, _ := strconv.Atoi(ranges[0])
			b, _ := strconv.Atoi(ranges[1])

			freshIng = append(freshIng, []int{a, b})
		} else {
			curr, _ := strconv.Atoi(row)
			for _, r := range freshIng {
				a, b := r[0], r[1]
				if a <= curr && curr <= b {
					res++
					break
				}
			}
		}
	}

	return res
}
