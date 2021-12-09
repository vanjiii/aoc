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

	fmt.Println(solve1(rows))
	fmt.Println(solve2(rows))
}

type digit struct {
	value  string
	marked bool
}

type board [5][5]digit

func solve1(rows []string) int {
	inputnums := strings.Split(rows[0], ",")
	boards := []board{}

	for i := 2; i < len(rows)-4; {
		boards = append(boards, createBoard(
			[]string{rows[i], rows[i+1], rows[i+2], rows[i+3], rows[i+4]},
		))

		i = i + 6
	}

	for _, tnum := range inputnums {
		for i := 0; i < len(boards); i++ {
			boards[i].addNum(tnum)
			if boards[i].win() {
				return boards[i].calcUnmarker(tnum)
			}
		}
	}

	return -1
}

func solve2(rows []string) int {
	inputnums := strings.Split(rows[0], ",")
	boards := []board{}

	for i := 2; i < len(rows)-4; {
		boards = append(boards, createBoard(
			[]string{rows[i], rows[i+1], rows[i+2], rows[i+3], rows[i+4]},
		))

		i = i + 6
	}

	for _, tnum := range inputnums {
		idx := []int{}
		for i := 0; i < len(boards); i++ {
			boards[i].addNum(tnum)

			if boards[i].win() {
				if len(boards) == 1 {
					return boards[i].calcUnmarker(tnum)
				}
				idx = append(idx, i)
			}
		}

		for i := len(idx) - 1; i >= 0; i-- {
			boards = remove(boards, idx[i])
		}
	}

	return -1
}

func remove(s []board, i int) []board {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (b board) calcUnmarker(num string) int {
	i, _ := strconv.Atoi(num)
	sum := 0
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if !b[row][col].marked {
				bb, _ := strconv.Atoi(b[row][col].value)
				sum += bb
			}
		}
	}

	return sum * i
}

func (b board) print() {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			fmt.Print(b[row][col])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (b board) win() bool {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if !b[row][col].marked {
				break
			}
			if col == 4 {
				return true
			}
		}
	}

	for col := 0; col < 5; col++ {
		for row := 0; row < 5; row++ {
			if !b[row][col].marked {
				break
			}
			if row == 4 {
				return true
			}
		}
	}

	return false
}

func (b *board) addNum(n string) {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if b[row][col].value == n {
				b[row][col].marked = true
			}
		}
	}
}

func createBoard(rows []string) board {
	b := [5][5]digit{}

	b[0] = str2ints(rows[0])
	b[1] = str2ints(rows[1])
	b[2] = str2ints(rows[2])
	b[3] = str2ints(rows[3])
	b[4] = str2ints(rows[4])

	return b
}

func str2ints(r string) [5]digit {
	dg := [5]digit{}

	r = strings.TrimSpace(r)
	nums := strings.Split(r, " ")
	var numbers []string
	for i := range nums {
		if nums[i] != "" {
			numbers = append(numbers, nums[i])
		}
	}

	for i := range numbers {
		dg[i] = digit{numbers[i], false}
	}

	return dg
}
