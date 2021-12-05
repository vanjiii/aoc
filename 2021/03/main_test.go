package main

import (
	"testing"

	"github.com/vanjiii/aoc/pkg"
)

func TestSolve(t *testing.T) {
	numbers := []string{
		"00100", "11110", "10110",
		"10111", "10101", "01111",
		"00111", "11100", "10000",
		"11001", "00010", "01010",
	}

	pkg.AssertInt(t, solve1(numbers), 198)
}
