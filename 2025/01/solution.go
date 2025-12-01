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
	position := 50
	zeros := 0

	for _, row := range rows {
		fmt.Println()
		fmt.Printf("pos: %v\n", position)
		fmt.Printf("zeros: %v\n", zeros)
		fmt.Printf("row: %v\n", row)

		direction := row[0]
		curr, err := strconv.Atoi(row[1:])
		if err != nil {
			return -1
		}

		zeros += curr / 100
		curr = curr % 100

		if direction == 'L' {
			curr *= -1
		}

		old := position
		position += curr
		if position < 0 {
			position += 100
			if old != 0 {
				zeros++
			}
		} else if position >= 100 {
			position -= 100
			zeros++
		} else if position == 0 {
			zeros++
		}

		fmt.Printf("new zeros: %v\n", zeros)
		fmt.Println()
	}

	return zeros
}

func partone(rows []string) int {
	position := 50
	zeros := 0

	for _, row := range rows {
		rot := row[0]
		curr, _ := strconv.Atoi(row[1:])

		curr = curr % 100

		if rot == 'L' {
			curr *= -1
		}

		position += curr
		if position < 0 {
			position += 100
		} else if position >= 100 {
			position -= 100
		}

		if position == 0 {
			zeros++
		}
	}

	return zeros
}
