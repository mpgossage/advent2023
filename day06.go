package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseRaces(filename string) (times []int, distances []int) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	times = StringToIntArray(lines[0][10:], " ")
	distances = StringToIntArray(lines[1][10:], " ")
	return
}

func day06a(filename string)  {
	times, distances := ParseRaces(filename)
	fmt.Printf("times %v\ndistances %v\n", times,distances)

	total := 1 // 1 as multiplicative

	for idx, time := range times {
		distance := distances[idx]
		// simplest method which works: brute force
		// this would be a binary search candidate, but thats later
		winning:=0
		for t := 1; t< time; t++{
			hold:=t
			move:= time-t
			travelled := move * hold
			//fmt.Printf("held %d move %d travelled %d\n", hold, move, travelled)
			if travelled > distance {
				winning ++
			}
		}
		total *= winning
	}
	fmt.Printf("result %d\n",total)
}

func ParseRaces2(filename string) (time int, distance int) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	timestr := strings.ReplaceAll(lines[0][10:], " ","")
	diststr := strings.ReplaceAll(lines[1][10:], " ","")
	time, _ = strconv.Atoi(timestr)
	distance, _ = strconv.Atoi(diststr)
	return
}

func day06b(filename string) {
	time, distance := ParseRaces2(filename)
	fmt.Printf("time %d distance %d\n", time, distance)

	// simplest method which works: brute force
	// this would be a binary search candidate, but thats later
	// (turned out to be a doddle, didn't need anything complex)
	winning := 0
	for t := 1; t < time; t++ {
		hold := t
		move := time - t
		travelled := move * hold
		//fmt.Printf("held %d move %d travelled %d\n", hold, move, travelled)
		if travelled > distance {
			winning++
		}
	}
	fmt.Printf("result %d\n", winning)
}