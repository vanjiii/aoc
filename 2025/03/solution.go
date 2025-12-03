package main

import (
	"fmt"
	"log"
	"strconv"

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

func parttwo(rows []string) int {
	res := 0

	for _, row := range rows {

		for i := range 12 {
			res += find(row, 12-i)
		}

	}

	return res
}

func find(s string, idx int) int {
	for i := len(s)-idx-1

	return 0
}

func partone(rows []string) int {
	res := 0
	for _, row := range rows {
		idx := 0
		maxi := byte('0')
		for i := range len(row) - 1 {
			if maxi < row[i] {
				idx = i
				maxi = row[i]
			}
		}

		maxy := byte('0')
		for i := idx + 1; i < len(row); i++ {
			if maxy < row[i] {
				idx = i
				maxy = row[i]
			}
		}

		tmp, _ := strconv.Atoi(string([]byte{maxi, maxy}))
		// fmt.Println(tmp)
		res += tmp
	}

	return res
}
