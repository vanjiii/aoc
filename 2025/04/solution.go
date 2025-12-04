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

func partone(grid [][]rune) int {
	lrow := len(grid)
	lcol := len(grid[0])

	res := 0

	treshhold := 4
	for row := range lrow {
		for col := range lcol {
			target := grid[row][col]

			if target == '.' {
				continue
			}

			neighbours := 0

			if row-1 >= 0 && grid[row-1][col] == '@' {
				neighbours++
			}
			if row+1 < lrow && grid[row+1][col] == '@' {
				neighbours++
			}

			if row-1 >= 0 && col-1 >= 0 && grid[row-1][col-1] == '@' {
				neighbours++
			}
			if col-1 >= 0 && grid[row][col-1] == '@' {
				neighbours++
			}
			if row+1 < lrow && col-1 >= 0 && grid[row+1][col-1] == '@' {
				neighbours++
			}

			if row-1 >= 0 && col+1 < lcol && grid[row-1][col+1] == '@' {
				neighbours++
			}
			if col+1 < lcol && grid[row][col+1] == '@' {
				neighbours++
			}
			if row+1 < lrow && col+1 < lcol && grid[row+1][col+1] == '@' {
				neighbours++
			}

			if neighbours < treshhold {
				res++
			}

		}
	}

	return res
}
