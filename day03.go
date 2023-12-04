package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Number struct {
	st, ed int // start & end index (inclusive)
	value int // actual numeric value
}

func find_numbers(source string) []Number {
	ValidDigits := "1234567890"
	var numbers []Number
	// iterate through the source finding a digit
	var n = Number{-1,-1,0}
	for idx, _ := range source {
		// contains() doesn't work on single char/rune so we use a substr
		s:= source[idx:idx+1]
		if strings.ContainsAny(s, ValidDigits) && n.st == -1 {
			// start found
			n.st = idx
		}
		if strings.ContainsAny(s, ValidDigits) == false && n.st != -1 {
			// end found
			n.ed = idx - 1
			n.value, _ = strconv.Atoi(source[n.st:idx])
			numbers = append(numbers, n)
			n = Number{-1,-1,0}
		}
	}
	// if ending item
	if n.st != -1 {
		n.ed = len(source) - 1
		n.value, _ = strconv.Atoi(source[n.st:])
		numbers = append(numbers, n)
	}

	return numbers
}

func is_valid_number(n Number, y int, grid []string) bool {
	// its easier to have a full list of symbols and use ContainsAny
	// than having a list of digits & finding whats not there
	Symbols := "!\"#$%&'()*+,-/:;<=>?@[\\]_{|}"

	// simple lambda
	trimstring := func(s string, st, ed int) string {
		if st < 0 {
			st = 0
		}
		if ed >= len(s) {
			ed = len(s)
		}
		return s[st:ed]
	}

	// above
	if y > 1 {
		above := trimstring(grid[y-1], n.st-1, n.ed+2)
		if strings.ContainsAny(above, Symbols) {
			return true
		}
	}
	line := trimstring(grid[y], n.st-1, n.ed+2)
	if strings.ContainsAny(line, Symbols) {
		return true
	}
	// below
	if y < len(grid)-1 {
		below := trimstring(grid[y+1], n.st-1, n.ed+2)
		if strings.ContainsAny(below, Symbols) {
			return true
		}
	}
	return false
}

func day03a(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	var total = 0
	for y, line := range lines {
		for _, num := range find_numbers(line) {
			if is_valid_number(num, y, lines) {
				total += num.value
			}
		}
	}
	fmt.Printf("day03a %d", total)
}

func day03b(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	// flipping logic: find all the * then check all the numbers and find which are near the *
	var stars = make(map[Coord][]int)
	for y, line := range lines {
		for x, c := range line {
			if c == '*' {
				stars[Coord{x, y}] = nil
			}
		}
	}
	//fmt.Printf("stars %v\n", stars)

	// get each number and if its near a star, add it to the stars list
	for y, line := range lines {
		for _, num := range find_numbers(line) {
			for star, nearby := range stars {
				if star.y >= y-1 && star.y <= y+1 && star.x >= num.st-1 && star.x <= num.ed+1 {
					//fmt.Printf("star %v is near number %v\n", star, num)
					stars[star] = append(nearby, num.value)
				}
			}
		}
	}
	//fmt.Printf("stars %v\n", stars)

	// sum those which have exactly 2 nearby numbers
	var total = 0
	for _, nearby := range stars {
		if len(nearby) == 2 {
			total += nearby[0] * nearby[1]
		}
	}
	fmt.Printf("result %d", total)
}