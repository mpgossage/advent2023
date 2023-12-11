package main

import (
	"fmt"
	"strings"
	"time"
)

func parse_camel_path(filename string) (steps string, nodes map[string][]string)  {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	steps = lines[0]
	nodes = make(map[string][]string)

	for _, line := range lines[2:] {
		line = strings.ReplaceAll(line, "= (","")
		line = strings.ReplaceAll(line, ",","")
		line = strings.ReplaceAll(line, ")","")
		path := strings.Split(line, " ")
		nodes[path[0]] = path[1:]
	}
	return
}

func traverse_camel_path(loc string, dir uint8, nodes map[string][]string) string {
	if dir == 'L' {
		return nodes[loc][0]
	}
	return nodes[loc][1]
}

func day08a(filename string)  {
	steps,nodes:= parse_camel_path(filename)
	fmt.Printf("steps %s\nnodes %+v\n", steps, nodes)

	steps_ln := len(steps)
	var count int
	loc,end:= "AAA","ZZZ"
	for count = 0; loc != end; count++ {
		dir := steps[count % steps_ln]
		fmt.Printf("step %d loc %s dir %c\n", count, loc, dir)
		loc  = traverse_camel_path(loc,dir,nodes)
	}
	fmt.Printf("step %d loc %s\n", count, loc)
	fmt.Printf("result %d\n", count)
}

func is_camel_path_ended(locs []string) bool {
	for _,loc := range locs {
		if strings.HasSuffix(loc, "Z") == false {
			return false
		}
	}
	return true
}

func _day08b(filename string)  {
	// my brain hurts reading the explanation, but at least the example is clear
	// however the algol is too slow, after 10 mins, its done 2G steps and not reached the end
	steps,nodes:= parse_camel_path(filename)
	fmt.Printf("steps %s\nnodes %+v\n", steps, nodes)

	var starts []string
	for node := range nodes {
		if strings.HasSuffix(node, "A"){
			starts = append(starts, node)
		}
	}

	start_time := time.Now()
	steps_ln := len(steps)
	var count int
	locs := starts
	for count = 0; is_camel_path_ended(locs) == false; count++ {
		dir := steps[count%steps_ln]
		//fmt.Printf("step %d locs %v dir %c\n", count, locs, dir)
		for i := range locs {
			locs[i]=traverse_camel_path(locs[i], dir, nodes)
		}
		if count % 1000000 == 0 {
			fmt.Printf("step %d time %.2fs\n", count, time.Now().Sub(start_time).Seconds())
		}
	}
	fmt.Printf("step %d locs %v\n", count, locs)
	fmt.Printf("result %d\n", count)
}

func find_repeated_camel_path(loc string, steps string, nodes map[string][]string) int {
	steps_ln := len(steps)
	var ends []int
	for count := 0; len(ends)<3; count++ {
		dir := steps[count%steps_ln]
		loc = traverse_camel_path(loc, dir, nodes)
		if strings.HasSuffix(loc, "Z") {
			ends = append(ends, count+1)
		}
	}
	// look at ends
	// originally I was expected it not be to be a simple periodic, but it was
	// so take the easy solution
	if len(ends)>=3 {
		d1:= ends[1]-ends[0]
		d2:= ends[2]-ends[1]
		//fmt.Printf("delta %d %d\n", d1,d2)
		if d1 == d2 {
			return d1
		}
	}
	return 0 // error
}


func day08b(filename string) {
	steps, nodes := parse_camel_path(filename)

	var starts []string
	for node := range nodes {
		if strings.HasSuffix(node, "A") {
			starts = append(starts, node)
		}
	}
	fmt.Printf("steps %s\nnodes %+v\nstarts %v\n", steps, nodes, starts)

	// logic change: the path will eventually reach a loop
	// of len L and it will pass the exit on point(s) A
	// so find them

	result:=1

	for idx, loc := range starts {
		val:=find_repeated_camel_path(loc, steps, nodes)
		fmt.Printf("path %d repeated after %d steps\n", idx, val)
		// lowest common multiplier
		result = LeastCommonMultiple(result, val)
	}
	fmt.Printf("result %d\n", result)
}
