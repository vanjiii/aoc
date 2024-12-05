package main

import (
	"fmt"
	"log"
	"math"
	"slices"
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

	// fmt.Printf("%+v \n", solveb([]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}))
	fmt.Printf("%+v \n", solve([]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, false))

	fmt.Printf("a: %+v \n", solve(input, false))
	// fmt.Printf("a: %+v \n", solveb(input))
}

// solve useMedian or useMean
func solve(pos []int, useMedian bool) int {
	// sort the array
	// find the median
	// calculate the distance from the median
	// return the distance
	slices.Sort(pos)

	delimiter := 0
	if useMedian {
		delimiter = pos[len(pos)/2]
	} else {
		sum := 0
		for i := range pos {
			sum += pos[i]
		}
		delimiter = sum / len(pos)
	}
	ans := 0
	for i := range pos {
		ans += int(math.Abs(float64(pos[i] - delimiter)))
	}

	return ans
}

// bruteforce
func solveb(pos []int) int {
	res := math.MaxInt32
	for i := 0; i < len(pos); i++ {
		currentFuel := 0
		for j := 0; j < len(pos); j++ {
			c := math.Abs(float64(pos[j]) - float64(pos[i]))
			currentFuel += int(c)
		}

		if currentFuel < res {
			res = currentFuel
		}
	}

	return res
}
