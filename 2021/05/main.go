package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/vanjiii/aoc/pkg"
)

func main() {
	rows, err := pkg.ReadStrings("input_real")
	if err != nil {
		log.Fatalf("fail to read input: %v", err)
		return
	}

	fmt.Println(solve(rows, false))
	fmt.Println(solve(rows, true))
}

type Point struct {
	x, y int
}

type Lines [][]Point

func (l Lines) print() {
	for i := range l {
		for j := range l[i] {
			fmt.Printf("%d ", l[i][j])
		}
		fmt.Println()
	}
}

func solve(rows []string, diagonal bool) int {
	points := make(Lines, len(rows))
	for i, r := range rows {
		arr := strings.Split(r, " -> ")

		line := []Point{
			str2point(arr[0]),
			str2point(arr[1]),
		}

		points[i] = line
	}

	// points.print()

	final := make([][]int, 1_000)
	for i := range final {
		final[i] = make([]int, 1_000)
	}

	for i := range points {
		p1 := points[i][0]
		p2 := points[i][1]

		if p1.x == p2.x {
			maxy := max(p1.y, p2.y)
			miny := min(p1.y, p2.y)
			for y := miny; y <= maxy; y++ {
				final[y][p1.x]++
			}
		} else if p1.y == p2.y {
			maxx := max(p1.x, p2.x)
			minx := min(p1.x, p2.x)
			for x := minx; x <= maxx; x++ {
				final[p1.y][x]++
			}
		} else if diagonal {
			maxx := max(p1.x, p2.x)
			minx := min(p1.x, p2.x)
			distance := maxx - minx

			for i := 0; i < distance+1; i++ {
				x := p1.x
				y := p1.y

				if p1.x < p2.x {
					x = x + i
				} else {
					x = x - i
				}

				if p1.y < p2.y {
					y = y + i
				} else {
					y = y - i
				}

				final[y][x]++
			}
		}
	}

	out := 0
	for i := range final {
		for j := range final[i] {
			if final[i][j] > 1 {
				out++
			}
		}
	}

	return out
}

func str2point(s string) Point {
	arr := strings.Split(s, ",")

	x, _ := strconv.Atoi(arr[0])
	y, _ := strconv.Atoi(arr[1])

	return Point{
		x: x,
		y: y,
	}
}
