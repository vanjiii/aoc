package main

import (
	"fmt"
	"log"
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
	// fmt.Printf("part two: %v\n", parttwo(rows))
}

func partone(rows []string) int {
	data := [][]string{}
	for _, row := range rows {
		parts := strings.Fields(row)

		data = append(data, parts)
	}

	// for _, row := range data {
	// 	for _, val := range row {
	// 		fmt.Printf("%s ", val)
	// 	}
	// 	fmt.Println()
	// }

	res := 0

	for i := range len(data[0]) {
		nums := []int{}
		for j := range len(data) - 1 {
			tmp := strings.Trim(data[j][i], " ")
			fmt.Println(data[j][i])
			a, err := strconv.Atoi(tmp)
			if err != nil {
				panic(err)
			}

			nums = append(nums, a)

		}

		op := data[len(data)-1][i]

		curr := 0

		switch op {
		case "+":
			for _, v := range nums {
				curr += v
			}
		case "*":
			m := 1
			for _, v := range nums {
				m *= v
			}
			curr += m
		}

		fmt.Println(nums)
		// fmt.Println(curr)

		res += curr
	}

	return res
}
