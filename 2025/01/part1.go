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

	fmt.Println(partone(rows))
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
