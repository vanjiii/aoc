package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
	sum := 0
	ints := []int{}
	for _, row := range rows {
		arr := strings.Split(row, "-")
		smaller, larger := arr[0], arr[1]
		smalleri, _ := strconv.Atoi(smaller)
		largeri, _ := strconv.Atoi(larger)

		for smalleri < largeri {
			ints = append(ints, smalleri)
			smalleri++
		}
	}

	for i := range len(ints) - 1 {
		if isRepeated(fmt.Sprintf("%v", ints[i])) {
			sum += ints[i]
		}
	}

	return sum
}

func partone(rows []string) int {
	sum := 0
	ints := []int{}
	for _, row := range rows {
		arr := strings.Split(row, "-")
		smaller, larger := arr[0], arr[1]
		smalleri, _ := strconv.Atoi(smaller)
		largeri, _ := strconv.Atoi(larger)

		for smalleri < largeri {
			ints = append(ints, smalleri)
			smalleri++
		}
	}

	for i := range len(ints) - 1 {
		if isPalindrom(fmt.Sprintf("%v", ints[i])) {
			sum += ints[i]
		}
	}

	return sum
}

func isRepeated(s string) bool {
	ss := s + s
	middle := ss[1 : len(ss)-1]

	return strings.Contains(middle, s)
}

func isPalindrom(s string) bool {
	ln := len(s)

	if ln%2 == 1 {
		return false
	}

	r, l := s[0:ln/2], s[ln/2:]

	return r == l
}
