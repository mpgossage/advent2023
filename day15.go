package main

import (
	"fmt"
	"strconv"
	"strings"
)

// What is this?
// Part A hash fn easy
// Part B looks very valuely like a hashtable, but strange
// took a lot of reading of the question to understand, but once read it all, it was easy
// sort of like a bucket hash

func deer_hash(str string) int {
	val := 0
	for _, c := range str {
		val += int(c)
		val *= 17
		val %= 256
	}
	return val
}

func day15a(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	//fmt.Printf("test HASH=%d\n", deer_hash("HASH"))

	total := 0
	for _, a := range strings.Split(lines[0], ",") {
		v := deer_hash(a)
		fmt.Printf("hash '%s'=%d\n", a, v)
		total += v
	}
	fmt.Printf("total %d\n", total)
}

func day15b(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	type lens struct {
		label string
		value int
	}
	var boxes [256][]lens

	for _, a := range strings.Split(lines[0], ",") {
		label := ""
		val := -1
		if idx := strings.IndexAny(a, "-="); idx >= 0 {
			label = a[:idx]
			if a[idx] == '=' {
				val, _ = strconv.Atoi(a[idx+1:])
			}
		}

		//fmt.Printf("label %s hash %d, val %d\n",label, deer_hash(label),val)
		box_idx := deer_hash(label)
		if val == -1 {
			// delete request
			var newbox []lens
			for _, b := range boxes[box_idx] {
				if b.label != label {
					newbox = append(newbox, b)
				}
			}
			boxes[box_idx] = newbox
		} else {
			found := false
			for i, b := range boxes[box_idx] {
				if b.label == label {
					boxes[box_idx][i].value = val
					found = true
				}
			}
			if found == false {
				boxes[box_idx] = append(boxes[box_idx], lens{label: label, value: val})
			}
		}
	}
	// print
	for i, b := range boxes {
		if len(b) > 0 {
			fmt.Printf("box %d %v\n", i, b)
		}
	}
	// score
	total := 0
	for i, b := range boxes {
		for j, l := range b {
			total += (i + 1) * (j + 1) * l.value
		}
	}
	fmt.Printf("total %d\n", total)
}
