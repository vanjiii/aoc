package main

import (
	"fmt"
	"log"

	"github.com/vanjiii/aoc/pkg"
)

func main() {
	grid, err := pkg.ReadGrid("input")
	if err != nil {
		log.Fatalf("fail to read input: %v", err)
		return
	}

	fmt.Println(partone(grid))

}

func partone(matrix [][]rune) int {
	res := 0

	for row := 1; row < len(matrix); row++ {
		for col := range len(matrix[0]) {
			above := matrix[row-1][col]
			current := matrix[row][col]

			if current == '.' {
				if above == 'S' || above == '|' {
					matrix[row][col] = '|'
				}
			}

			if current == '^' && above == '|' {
				matrix[row][col-1] = '|'
				matrix[row][col+1] = '|'
				res++
			}
		}
	}

	// printm(matrix)

	return res
}

func printm(m [][]rune) {
	fmt.Println("<<< EOT")
	for row := range len(m) {
		for col := range len(m[0]) {
			fmt.Printf("%v ", string(m[row][col]))
		}
		fmt.Println()
	}
	fmt.Println("EOT")
}
