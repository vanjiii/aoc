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

	fmt.Println(parttwo(rows))
}

func parttwo(rows []string) int {
	res := 0

	for _, row := range rows {
		elements := []byte{}

		startidx := 0
		el := byte('x')
		for i := range 12 {
			startidx, el = find(row, startidx, 12-i)
			elements = append(elements, el)
		}

		tmp, _ := strconv.Atoi(string(elements))
		fmt.Println(tmp)
		res += tmp

	}

	return res
}

func find(s string, startidx, endidx int) (int, byte) {
	ll := len(s) - endidx

	maxel := byte('0')
	j := -1
	for i := startidx; i <= ll; i++ {
		if maxel < s[i] {
			maxel = s[i]
			j = i + 1
		}
	}

	return j, maxel
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
