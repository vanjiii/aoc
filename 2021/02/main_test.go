package main

import (
	"testing"

	"github.com/vanjiii/aoc/pkg"
)

func Test01(t *testing.T) {
	cmds := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	res := solve1(cmds)

	if res != 150 {
		t.Fatal(res)
	}
}

func Test02(t *testing.T) {
	cmds := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}

	res := solve2(cmds)

	pkg.AssertInt(t, res, 900)
}
