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
	default:
		fmt.Printf("no such day %s", day)
	}

}

