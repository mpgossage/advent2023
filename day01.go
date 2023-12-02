package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day01a(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	var total = 0
	for _,line := range lines {
		// simplest possible, this will fail for a multi digit number & if errors occur
		s_idx := strings.IndexAny(line ,"0123456789")
		e_idx := strings.LastIndexAny(line ,"0123456789")
		s_val,_ := strconv.Atoi(line[s_idx:s_idx+1])
		e_val,_ := strconv.Atoi(line[e_idx:e_idx+1])
		total += s_val*10 + e_val
	}
	fmt.Printf("%d", total)
}

func day01b(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	var total = 0
	for _,line := range lines {
		// look for ints
		var s_val = 0
		var e_val = 0

		var s_idx = strings.IndexAny(line ,"0123456789")
		if s_idx != -1 {
			s_val,_ = strconv.Atoi(line[s_idx:s_idx+1])
		} else {
			// its at the end of line, all will be earlier than this
			s_idx = len(line)
		}
		var e_idx = strings.LastIndexAny(line ,"0123456789")
		if e_idx != -1 {
			e_val,_ = strconv.Atoi(line[e_idx:e_idx+1])
		}
		// now see if there is a hidden number string before this
		digits := map[int]string {
			1 : "one",
			2 : "two",
			3 : "three",
			4 : "four",
			5 : "five",
			6 : "six",
			7 : "seven",
			8 : "eight",
			9 : "nine",
		}
		for v,digit := range digits {
			st_idx := strings.Index(line, digit)
			if st_idx != -1 && st_idx < s_idx {
				s_idx = st_idx
				s_val = v
			}
			ed_idx := strings.LastIndex(line, digit)
			if ed_idx != -1 && ed_idx > e_idx {
				e_idx = ed_idx
				e_val = v
			}
		}

		total += s_val*10 + e_val
	}
	fmt.Printf("%d", total)
}
