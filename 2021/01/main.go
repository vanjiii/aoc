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

	fmt.Println(solve1(ints))
}

func solve1(ints []int) int {
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

func solve2(ints []int) int {
	incr := 0

	for i := 0; i < len(ints)-3; i++ {
		curr := ints[i] + ints[i+1] + ints[i+2]
		next := ints[i+1] + ints[i+2] + ints[i+3]

		if curr < next {
			incr++
		}
	}
	return incr
}
