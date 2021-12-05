package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/vanjiii/aoc/pkg"
)

type bit struct {
	zero, one int
}

func main() {
	numbers, err := pkg.ReadStrings("input")
	if err != nil {
		log.Fatalf("fail to read input: %v", err)
		return
	}

	fmt.Println(solve1(numbers))
	fmt.Println(solve2(numbers))
}

func solve2(nums []string) int {
	zeros := []string{}
	ones := []string{}
	numbers := nums

	counter := 0
	for {
		for _, v := range numbers {
			switch string(v[counter]) {
			case "1":
				ones = append(ones, v)
			case "0":
				zeros = append(zeros, v)
			}
		}

		numbers = left(ones, zeros)
		ones = []string{}
		zeros = []string{}

		if len(numbers) == 1 {
			break
		}

		counter++
	}

	// oxigen generator rating
	fmt.Println(numbers)

	znums := nums
	ones = []string{}
	zeros = []string{}

	counter = 0
	for {
		for _, v := range znums {
			switch string(v[counter]) {
			case "1":
				ones = append(ones, v)
			case "0":
				zeros = append(zeros, v)
			}
		}

		znums = lefto(ones, zeros)
		ones = []string{}
		zeros = []string{}

		if len(znums) == 1 {
			break
		}

		counter++
	}
	fmt.Println(znums)

	return bin2dec(znums[0]) * bin2dec(numbers[0])
}

func lefto(o, z []string) []string {
	if len(o) < len(z) {
		return o
	}

	return z
}

func left(o, z []string) []string {
	if len(o) < len(z) {
		return z
	}

	return o
}

func solve1(numbers []string) int {
	bits := make([]*bit, len(numbers[0]))

	for i := range bits {
		bits[i] = new(bit)
	}

	for _, n := range numbers {
		for i := range bits {
			count(string(n[i]), bits[i])
		}
	}

	var sb strings.Builder
	for i := range bits {
		sb.WriteString(bits[i].getB())
	}

	gamarate := bin2dec(sb.String())

	shift := (1 << len(numbers[0])) - 1

	epsilon := gamarate ^ shift

	return gamarate * epsilon
}

func (b bit) getB() string {
	if b.one > b.zero {
		return "1"
	}

	return "0"
}

func count(bit string, counter *bit) {
	switch bit {
	case "1":
		counter.one++
	case "0":
		counter.zero++
	}
}

func bin2dec(number string) int {
	i, err := strconv.ParseInt(number, 2, 0)
	if err != nil {
		panic(err)
	}
	return int(i)
}
