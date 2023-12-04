package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parse_draw(s string) (int, int, int){
	// use a map for simple storage rather than a lookup on name
	totals := map[string]int{ "red":0, "green":0, "blue":0 }

	// given "3 blue, 4 red" we split on comma then space
	for _, item := range strings.Split(s, ","){
		v := strings.Fields(item)
		amt, _ := strconv.Atoi(v[0])
		totals[v[1]] += amt
	}

	return totals["red"], totals["green"], totals["blue"]
}

func day02a(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	max_red := 12
	max_green := 13
	max_blue := 14

	var total = 0
	for _, line := range lines {
		splitline := strings.Split(line, ":")
		//fmt.Printf("splitline %v\n", splitline)
		// v1[0] = "Game 5"
		gameid,_ := strconv.Atoi(splitline[0][5:])
		//fmt.Printf("game %d sl '%s'\n", gameid, splitline[0][5:])

		var possible = true
		for _, pull := range strings.Split(splitline[1],";") {
			red, green, blue := parse_draw(pull)
			if red > max_red || green > max_green || blue > max_blue {
				possible = false
			}
		}

		//fmt.Printf("game %d possible %v\n", gameid, possible)
		if possible {
			total += gameid
		}
	}
	fmt.Printf("total %d\n", total)
}

func day02b(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	var total = 0
	for _, line := range lines {
		splitline := strings.Split(line, ":")
		// v1[0] = "Game 5"

		var min_red = 0
		var min_green = 0
		var min_blue = 0

		for _, pull := range strings.Split(splitline[1],";") {
			red, green, blue := parse_draw(pull)
			min_red = Max(min_red, red)
			min_green = Max(min_green, green)
			min_blue = Max(min_blue, blue)
		}
		total += min_red * min_green * min_blue
	}
	fmt.Printf("total %d\n", total)
}