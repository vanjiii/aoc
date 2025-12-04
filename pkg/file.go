package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadInts(file string) ([]int, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("fail to open file: %w", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	input := make([]int, 0)
	for sc.Scan() {
		num, err := strconv.Atoi(sc.Text())
		if err != nil {
			return nil, fmt.Errorf("fail to convert line: %v, err: %w", sc.Text(), err)
		}

		input = append(input, num)
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("scan file error: %w", err)
	}

	return input, nil
}

func ReadStrings(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("fail to open file: %w", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	input := make([]string, 0)
	for sc.Scan() {
		input = append(input, sc.Text())
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("scan file error: %w", err)
	}

	return input, nil
}

func ReadGrid(file string) ([][]rune, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}
