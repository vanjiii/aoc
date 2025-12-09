package main

import (
	"fmt"
	"log"

	"github.com/vanjiii/aoc/pkg"
)

func main() {
	grid, err := pkg.ReadGrid("sample")
	if err != nil {
		log.Fatalf("fail to read input: %v", err)
		return
	}

	res := 0
	gr := grid

	for {
		// printm(gr)
		curr := 0
		curr, gr = partone(gr)
		if 0 != curr {
			res += curr
		} else {
			break
		}
	}

	fmt.Println(res)
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

func partone(grid [][]rune) (int, [][]rune) {
	lrow := len(grid)
	lcol := len(grid[0])

	matrix := make([][]rune, lrow)
	for i := range matrix {
		matrix[i] = make([]rune, lcol)
	}

	res := 0

	treshhold := 4
	for row := range lrow {
		for col := range lcol {
			target := grid[row][col]

			matrix[row][col] = target
			if target == '.' || target == 'x' {
				continue
			}

			rows := []int{}
			cols := []int{}

			neighbours := 0

			if row-1 >= 0 && grid[row-1][col] == '@' {
				rows = append(rows, row-1)
				cols = append(cols, col)
				neighbours++
			}
			if row+1 < lrow && grid[row+1][col] == '@' {
				rows = append(rows, row+1)
				cols = append(cols, col)
				neighbours++
			}

			if row-1 >= 0 && col-1 >= 0 && grid[row-1][col-1] == '@' {
				rows = append(rows, row-1)
				cols = append(cols, col-1)
				neighbours++
			}
			if col-1 >= 0 && grid[row][col-1] == '@' {
				rows = append(rows, row)
				cols = append(cols, col-1)
				neighbours++
			}
			if row+1 < lrow && col-1 >= 0 && grid[row+1][col-1] == '@' {
				rows = append(rows, row+1)
				cols = append(cols, col-1)
				neighbours++
			}

			if row-1 >= 0 && col+1 < lcol && grid[row-1][col+1] == '@' {
				rows = append(rows, row-1)
				cols = append(cols, col+1)
				neighbours++
			}
			if col+1 < lcol && grid[row][col+1] == '@' {
				rows = append(rows, row)
				cols = append(cols, col+1)
				neighbours++
			}
			if row+1 < lrow && col+1 < lcol && grid[row+1][col+1] == '@' {
				rows = append(rows, row+1)
				cols = append(cols, col+1)
				neighbours++
			}

			if neighbours < treshhold {
				res++
				matrix[row][col] = 'x'
			}

		}
	}

	return res, matrix
}
