package main

import (
	"fmt"
	"os"
)

/*
to run this: rmb on advent2023 and run=>go build advent2023
to run all tests: rmb on advent2023 and run=>go test advent2023
(ps you can command line all these)
not 100% happy about
 */

func main() {

	var day = "1b"
	var filename = "data/test01b.txt"
	if len(os.Args) < 2 {
		fmt.Println("advent2023 <day> <filename>")
		fmt.Println("  using fallback for now")
	} else {
		day = os.Args[1]
		filename = os.Args[2]
	}

	fmt.Printf("advent2023 day:%s file:%s\n", day, filename)
	switch day {
	case "1a":
		day01a(filename)
	case "1b" :
		day01b(filename)
	case "2a" :
		day02a(filename)
	case "2b" :
		day02b(filename)
	case "3a" :
		day03a(filename)
	case "3b" :
		day03b(filename)
	case "4a" :
		day04a(filename)
	case "4b" :
		day04b(filename)
	case "5a" :
		day05a(filename)
	case "5b" :
		day05b(filename)
	case "6a" :
		day06a(filename)
	case "6b" :
		day06b(filename)
	case "7a" :
		day07a(filename)
	case "7b" :
		day07b(filename)
	case "8a" :
		day08a(filename)
	case "8b" :
		day08b(filename)
	case "9a" :
		day09a(filename)
	case "9b" :
		day09b(filename)
	case "10a" :
		day10a(filename)
	case "10b" :
		day10b(filename)
	case "11a" :
		day11a(filename)
	case "11b" :
		day11b(filename)
	case "12a" :
		day12a(filename)
	case "12b" :
		day12b(filename)
	case "13a" :
		day13a(filename)
	case "13b" :
		day13b(filename)
	case "14a" :
		day14a(filename)
	case "14b" :
		day14b(filename)
	case "15a" :
		day15a(filename)
	case "15b" :
		day15b(filename)
	case "16a" :
		day16a(filename)
	case "16b" :
		day16b(filename)
	case "17a" :
		day17a(filename)
	case "17b" :
		day17b(filename)
	case "18a" :
		day18a(filename)
	case "18b" :
		day18b(filename)
	case "19a" :
		day19a(filename)
	case "19b" :
		day19b(filename)
	default:
		fmt.Printf("no such day %s", day)
	}

}

