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
	input := []int{}
	inputstr := strings.Split(rows[0], ",")
	for i := range inputstr {
		val, _ := strconv.Atoi(inputstr[i])
		input = append(input, val)
	}

	fmt.Printf("%+v \n", solve(input, 256))
	// fmt.Printf("%+v \n", solve([]int{3, 4, 3, 1, 2}, 18))

}

// track all the fish states (using map) at once
func solve(fishes []int, days int) int {
	// more like every possible fish age
	states := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}

	// init the map state
	for i := 0; i < len(fishes); i++ {
		states[fishes[i]]++
	}

	for i := 0; i < days; i++ {
		// fmt.Printf("day: %v: %+v\n", i, states)
		eight := 0
		six := 0
		for k := 0; k < 9; k++ {
			if k == 0 {
				// we save those for later because the may fack up the
				// iteration
				six = states[0]
				eight += states[0]
			} else {
				states[k-1] = states[k]
			}

		}
		states[6] += six
		states[8] = eight
	}

	count := 0
	for k := range states {
		count += states[k]
	}

	return count
}

// track the state of every fish (individually) every day
func stupidSolve(fishes []int, days int) int {
	for i := 0; i < days; i++ {
		lenfishes := len(fishes)
		for j := 0; j < lenfishes; j++ {
			if fishes[j] == 0 {
				fishes[j] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[j]--
			}
		}

		// fmt.Printf("%+v \n", fishes)
	}

	return len(fishes)
}
