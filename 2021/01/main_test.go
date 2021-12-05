package main

import "testing"

func Test01(t *testing.T) {
	ints := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	res := solve1(ints)

	if res != 7 {
		t.Error(res)
	}
}
func Test02(t *testing.T) {
	ints := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	res := solve2(ints)

	if res != 5 {
		t.Error(res)
	}
}
