package main

import (
	"fmt"
	"log"

	"github.com/vanjiii/aoc/pkg"
)

func main() {
	rows, err := pkg.ReadStrings("sample")
	if err != nil {
		log.Fatalf("fail to read input: %v", err)
		return
	}

	fmt.Println(partone(rows))
}

func partone(rows []string) string {
	for _, row := range rows {
	}
}
