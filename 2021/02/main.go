package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/vanjiii/aoc/pkg"
)

func main() {
	cmds, err := pkg.ReadStrings("input")
	if err != nil {
		log.Fatalf("fail to read input: %v", err)
		return
	}

	fmt.Println(solve1(cmds))

}

type coord struct {
	x, y int
}

func solve1(cmds []string) int {
	coords := new(coord)

	for _, v := range cmds {
		coords.applyCmd(v)
	}

	return coords.x * coords.y
}

func (c *coord) applyCmd(cmd string) {
	parts := strings.Split(cmd, " ")
	action := parts[0]
	step, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	switch action {
	case "forward":
		c.x = c.x + step
	case "down":
		c.y = c.y + step
	case "up":
		c.y = c.y - step
	default:
		panic(cmd)
	}
}
