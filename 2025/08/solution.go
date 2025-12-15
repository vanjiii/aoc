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

	fmt.Println(partone(rows))
}

func partone(rows []string) int {
	res := 0

	for _, row := range rows {
		x, y := strings.Split(row, ",")[0], strings.Split(row, ",")[1]

		for _, row := range rows {
			x1, y1 := strings.Split(row, ",")[0], strings.Split(row, ",")[1]

			rr := abs(x, x1) + 1
			cc := abs(y, y1) + 1

			if rr*cc > res {
				res = rr * cc
			}
		}
	}

	return res
}

func abs(as, bs string) int {
	a, _ := strconv.Atoi(as)
	b, _ := strconv.Atoi(bs)

	if a-b > 0 {
		return a - b
	} else {
		return b - a
	}
}
