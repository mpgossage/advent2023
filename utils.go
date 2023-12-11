package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func StringToIntArray(str, sep string) []int {
	var result []int
	for _, s := range strings.Split(str, sep) {
		if val, err := strconv.Atoi(s); err == nil {
			result = append(result, val)
		}
	}
	return result
}

func LeastCommonMultiple(n,m int) int {
	// https://stackoverflow.com/questions/3154454/what-is-the-most-efficient-way-to-calculate-the-least-common-multiple-of-two-int
	n1, m1 := n, m
	for m1 != n1 {
		if m1 > n1 {
			n1 += n
		} else {
			m1 += m
		}
	}
	return m1
}

type Coord struct {
	x,y int
}