package main

import (
	"fmt"
	"log"

	"github.com/vanjiii/aoc/pkg"
)

func main() {
	ints, err := pkg.ReadInts("input")
	if err != nil {
		log.Fatalf("fail to read input: %v", err)
		return
	}
	fmt.Println(solve(ints))
}

func solve(ints []int) int {
	prev := ints[0]
	incr := 0

	for i := 1; i < len(ints); i++ {
		curr := ints[i]
		if curr > prev {
			incr++
		}
		prev = curr
	}
	return incr
}
