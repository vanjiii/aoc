package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/vanjiii/aoc/pkg"
)

func main() {
	numbers, err := pkg.ReadStrings("input")
	if err != nil {
		log.Fatalf("fail to read input: %v", err)
		return
	}

	fmt.Println(solve1(numbers))
}

type bit struct {
	zero, one int
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
