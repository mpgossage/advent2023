package main

import (
	"bufio"
	"os"
)

// testing: do I need a comment
func LoadLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result [] string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err:= scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

// golang doesn't have a max/min for int!!!
// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// returns is difference between a & b is <= dist
// turned out to be more trouble than it was worth
// & I needed to write a whole bundle of unit tests for it
func INear(a,b int, dist int) bool {
	delta := a-b
	if delta >= 0 {
		return delta <= dist
	} else {
		return -delta <= dist
	}
}

type Coord struct {
	x,y int
}