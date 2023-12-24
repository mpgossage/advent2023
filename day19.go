package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// this is looking like a fun challenge
// finally caved to do a regex rather than string munging
// parsing was tedious, rest was ok
// part B was ohhh *****!
// very obviously not brute forceable with 256T combinations to try
// looks like some kind of linear range challenge
// this took some SERIOUS thought, but it was very much a cheer when I got the answer first time

type workflow struct {
	cat     rune // category (x,m,a,s), space for no condition
	val     int  // value
	greater bool // true:>, false:<
	result  string
}

func parse_machines(lines []string) (steps map[string][]workflow, parts []map[rune]int) {
	process_parts := false
	re := regexp.MustCompile("([xmas])([<>])(\\d+):([a-zA-Z]+)")
	steps = make(map[string][]workflow)

	for _, line := range lines {
		if len(line) == 0 {
			process_parts = true
			continue
		}
		if process_parts {
			// {x=787,m=2655,a=1222,s=2876}
			arr := strings.Split(line[1:len(line)-1], ",")
			part := make(map[rune]int)
			for _, a := range arr {
				kv := strings.Split(a, "=")
				part[rune(kv[0][0])], _ = strconv.Atoi(kv[1])
			}
			parts = append(parts, part)
		} else {
			// px{a<2006:qkq,m>2090:A,rfg}
			kv := strings.Split(line, "{")
			workflow_name := kv[0]
			var flows []workflow
			arr := strings.Split(strings.TrimRight(kv[1], "}"), ",")
			//fmt.Printf("%v\n",arr)
			for _, a := range arr {
				if strings.Contains(a, ":") {
					// a<2006:qkq
					parse := re.FindStringSubmatch(a)
					if len(parse) != 5 {
						fmt.Printf("Error '%s' parsed to %v\n", a, parse)
					}
					//fmt.Printf("'%s' %v\n",a,parse)
					// parse[0] is the full string
					wf := workflow{}
					wf.cat = rune(parse[1][0])
					wf.greater = parse[2] == ">"
					wf.val, _ = strconv.Atoi(parse[3])
					wf.result = parse[4]
					flows = append(flows, wf)
				} else {
					// just another wf
					flows = append(flows, workflow{cat: ' ', val: 0, result: a})
				}
			}
			steps[workflow_name] = flows
		}
	}
	return
}

func process_workflow(flow []workflow, part map[rune]int) string {
	for _, wf := range flow {
		// empty cat is auto flow
		if wf.cat == ' ' {
			return wf.result
		}
		val := part[wf.cat]
		if (wf.greater && val > wf.val) || (wf.greater == false && val < wf.val) {
			return wf.result
		}
	}
	return ""
}

func validate_part(steps map[string][]workflow, part map[rune]int) bool {
	loc := "in"
	for loc != "R" && loc != "A" {
		loc = process_workflow(steps[loc], part)
	}
	return loc == "A"
}

func day19a(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	//fmt.Printf("%v\n", lines)
	steps, parts := parse_machines(lines)
	/*fmt.Printf("%v\n",steps)
	fmt.Printf("%v\n",parts)

	part:=parts[0]
	loc:=process_workflow(steps["in"],part)
	fmt.Printf("%v\n",loc)
	fmt.Printf("validate %v\n",validate_part(steps,part))*/
	total := 0
	for _, part := range parts {
		if validate_part(steps, part) {
			for _, v := range part {
				total += v
			}
		}
	}
	fmt.Printf("total %d\n", total)
}

type linear_range struct {
	flow   string // which state its in
	x1, x2 int    // min max values for X (inclusive)
	m1, m2 int
	a1, a2 int
	s1, s2 int
}

// given range low..high(inclusive) & split condition (val, greater_than)
// returns low & high for pass & fail
// note: if range fully passes/fails then the other range will have low & high ==-1
func split_linear_range(low, high int, val int, greater_than bool) (low_pass, high_pass, low_fail, high_fail int) {
	// 6 possibilites
	if greater_than {
		if val < low {
			// all pass
			return low, high, -1, -1
		} else if val >= high {
			// all fail
			return -1, -1, low, high
		}
		// its a split: low values fail, high passes
		return val + 1, high, low, val
	}
	// less than
	if val <= low {
		// all fail
		return -1, -1, low, high
	} else if val > high {
		// all pass
		return low, high, -1, -1
	}
	// its a split: low values pass, high fails
	return low, val - 1, val, high
}

func process_linear_range(lrange linear_range, wf workflow) []linear_range {
	// split based upon the category to test

	var result []linear_range
	switch wf.cat {
	case ' ':
		// if unconditional its easy
		lrange.flow = wf.result
		return []linear_range{lrange}
	case 'x':
		low_pass, high_pass, low_fail, high_fail := split_linear_range(lrange.x1, lrange.x2, wf.val, wf.greater)
		if low_fail != -1 {
			lrange.x1, lrange.x2 = low_fail, high_fail
			result = append(result, lrange)
		}
		if low_pass != -1 {
			lrange.x1, lrange.x2 = low_pass, high_pass
			lrange.flow = wf.result // the pass result
			result = append(result, lrange)
		}
	case 'm':
		low_pass, high_pass, low_fail, high_fail := split_linear_range(lrange.m1, lrange.m2, wf.val, wf.greater)
		if low_fail != -1 {
			lrange.m1, lrange.m2 = low_fail, high_fail
			result = append(result, lrange)
		}
		if low_pass != -1 {
			lrange.m1, lrange.m2 = low_pass, high_pass
			lrange.flow = wf.result // the pass result
			result = append(result, lrange)
		}
	case 'a':
		low_pass, high_pass, low_fail, high_fail := split_linear_range(lrange.a1, lrange.a2, wf.val, wf.greater)
		if low_fail != -1 {
			lrange.a1, lrange.a2 = low_fail, high_fail
			result = append(result, lrange)
		}
		if low_pass != -1 {
			lrange.a1, lrange.a2 = low_pass, high_pass
			lrange.flow = wf.result // the pass result
			result = append(result, lrange)
		}
	case 's':
		low_pass, high_pass, low_fail, high_fail := split_linear_range(lrange.s1, lrange.s2, wf.val, wf.greater)
		if low_fail != -1 {
			lrange.s1, lrange.s2 = low_fail, high_fail
			result = append(result, lrange)
		}
		if low_pass != -1 {
			lrange.s1, lrange.s2 = low_pass, high_pass
			lrange.flow = wf.result // the pass result
			result = append(result, lrange)
		}
	}
	return result
}

func process_linear_range_set(lrange linear_range, flows []workflow) []linear_range {
	flow_name := lrange.flow
	var processed []linear_range
	todo := []linear_range{lrange}
	// for each step in the flows process
	for _, flow := range flows {
		for _, t := range todo {
			if t.flow == flow_name {
				processed = append(processed, process_linear_range(t, flow)...)
			} else {
				// its already jumped out
				processed = append(processed, t)
			}
		}
		// copy back
		todo = processed
		processed = nil
	}
	return todo
}

func day19b(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	steps, _ := parse_machines(lines)

	start_range := linear_range{flow: "in", x1: 1, x2: 4000, m1: 1, m2: 4000, a1: 1, a2: 4000, s1: 1, s2: 4000}

	/*	fmt.Printf("start %v\n",start_range)
		//result:=process_linear_range(start_range,steps["in"][0])
		result:=process_linear_range_set(start_range,steps["in"])
		fmt.Printf("result %v\n",result)*/

	ranges := []linear_range{start_range}
	var accept, reject []linear_range
	// start processing
	start_time := time.Now()
	for count := 0; len(ranges) > 0; count++ {
		// pop
		var r linear_range
		r, ranges = ranges[len(ranges)-1], ranges[:len(ranges)-1]
		// checking now if accept/reject
		if r.flow == "A" {
			accept = append(accept, r)
		} else if r.flow == "R" {
			reject = append(reject, r)
		} else {
			ranges = append(ranges, process_linear_range_set(r, steps[r.flow])...)
		}
		if count%1000 == 0 {
			fmt.Printf("count %d todo %d acc %d rej %d time %.2f\n", count, len(ranges), len(accept), len(reject), time.Now().Sub(start_time).Seconds())
		}
	}
	fmt.Printf("final todo %d acc %d rej %d time %.2f\n", len(ranges), len(accept), len(reject), time.Now().Sub(start_time).Seconds())
	total := int64(0)
	for _, acc := range accept {
		fmt.Printf("%v\n", acc)
		total += int64(1+acc.x2-acc.x1) * int64(1+acc.m2-acc.m1) * int64(1+acc.a2-acc.a1) * int64(1+acc.s2-acc.s1)
	}
	fmt.Printf("total %d\n", total)
}
