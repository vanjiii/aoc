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
	fmt.Println(solve2(cmds))
}

type coord struct {
	x, y, aim int
}

func solve2(cmds []string) int {
	coords := new(coord)

	for _, v := range cmds {
		coords.applyCmd2(v)
	}

	return coords.x * coords.y
}

func (c *coord) applyCmd2(cmd string) {
	parts := strings.Split(cmd, " ")
	action := parts[0]
	step, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	switch action {
	case "forward":
		c.y = c.y + (c.aim * step)
		c.x = c.x + step
	case "down":
		c.aim = c.aim + step
	case "up":
		c.aim = c.aim - step
	default:
		panic(cmd)
	}
}

func solve1(cmds []string) int {
	coords := new(coord)

	for _, v := range cmds {
		coords.applyCmd1(v)
	}

	return coords.x * coords.y
}

func (c *coord) applyCmd1(cmd string) {
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
