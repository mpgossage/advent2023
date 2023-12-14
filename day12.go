package main

import (
	"fmt"
	"strings"
	"time"
)

// looking at this, I'm thinking "Nonogram" which is familiar territory
// as expected the part B increases the complexity and it dies (3 mins for the first 24 items)
// therefore do for memorisation as an optimisation (took less than a second :-) )

func split_nono_line(line string) (string, []int)  {
	v:= strings.SplitN(line, " ", 2)
	return v[0],StringToIntArray(v[1], ",")
}

// matches all the chars of 'test' to 'line'
// if line is longer than test, the extra are ignored
func nono_match(line string, test string) bool {
	if len(test) > len(line) {
		return false
	}
	for i := range test {
		l := line[i]
		t := test[i]
		// if l is know, then it be the same
		if (l == '.' || l == '#') && t != l {
			return false
		}
		// otherwise assume ok
	}
	return true
}

func nono_arrangements(line string, vals []int) int  {
	//fmt.Printf("nono_arrangements %s %v\n", line,vals)

	// this is a recursive problem: actual source data has len(vals)==6, so having 6 versions is messy
	// nono_arrangements("??", [1]) has 2 possible solutions (".#" & "#.")
	// nono_arrangements("??", [2]) has 1 possible solutions ("##")
	// nono_arrangements("?.", [2]) has 0 possible solutions (impossible)
	// nono_arrangements("##.??", [2,1]) could be considered
	//    nono_arrangements("##", [2]) + nono_arrangements("??", [1])
	// note: the . is not part of calculations
	if len(vals) == 1 {
		val := vals[0]
		ln := len(line)
		if val > ln {
			return 0
		}

		result := 0
		for spc := 0; spc+val <= ln; spc++ {
			// the proposed line is '.'*spc + '#'*val +some '.'
			s := strings.Repeat(".", spc) + strings.Repeat("#", val) + strings.Repeat(".", ln-spc-val)
			//fmt.Printf("Val %d spc %d '%s'\n", val, spc, s)
			if nono_match(line, s) {
				result++
			}
		}
		return result
	}
	// multiple (here is gets fun)
	// for a set of values {a,b,c}, len must be at least a+b+c+2 (the number of gaps)
	// so we have a certain amount of movement room
	sum:=len(vals)-1
	for _,v:=range vals{
		sum+=v
	}
	ln := len(line)
	if sum > ln {
		return 0
	}
	val:=vals[0]
	result := 0
	for spc := 0; spc+sum <= ln; spc++ {
		// the proposed line is '.'*spc + '#'*val
		s := strings.Repeat(".", spc) + strings.Repeat("#", val)
		ln_s:=spc+val

		// see if 's' could fit in the line, BUT: the next char must not be a block
		nm := nono_match(line, s)
		lm:=line[ln_s]!='#'
		//fmt.Printf("match '%s' '%s' %v clear space %v\n", line,s,nm,lm)
		//if nono_match(line, s) && line[ln_s+1]!='#' {
		if nm && lm {
			// test the rest & get it options.
			// if it returns 0 then the rest didn't match
			v:=nono_arrangements(line[ln_s+1:], vals[1:])
			//fmt.Printf("nono_arrangements '%s' %v gives %v\n", line[ln_s+1:], vals[1:], v)
			result+= v
		}
	}
	return result
}

func nono_arrangements_memorised(line string, vals []int, memorise map[string]int) int  {

	if len(vals) == 1 {
		val := vals[0]
		ln := len(line)
		if val > ln {
			return 0
		}

		lookup:=fmt.Sprintf("%s-%v",line,vals)
		if score,found:= memorise[lookup]; found {
			return score
		}

		result := 0
		for spc := 0; spc+val <= ln; spc++ {
			// the proposed line is '.'*spc + '#'*val +some '.'
			s := strings.Repeat(".", spc) + strings.Repeat("#", val) + strings.Repeat(".", ln-spc-val)
			//fmt.Printf("Val %d spc %d '%s'\n", val, spc, s)
			if nono_match(line, s) {
				result++
			}
		}
		memorise[lookup]=result
		return result
	}
	// multiple (here is gets fun)
	// for a set of values {a,b,c}, len must be at least a+b+c+2 (the number of gaps)
	// so we have a certain amount of movement room
	sum:=len(vals)-1
	for _,v:=range vals{
		sum+=v
	}
	ln := len(line)
	if sum > ln {
		return 0
	}

	lookup:=fmt.Sprintf("%s-%v",line,vals)
	if score,found:= memorise[lookup]; found {
		return score
	}

	val:=vals[0]
	result := 0
	for spc := 0; spc+sum <= ln; spc++ {
		// the proposed line is '.'*spc + '#'*val
		s := strings.Repeat(".", spc) + strings.Repeat("#", val)
		ln_s:=spc+val

		// see if 's' could fit in the line, BUT: the next char must not be a block
		nm := nono_match(line, s)
		lm:=line[ln_s]!='#'
		//fmt.Printf("match '%s' '%s' %v clear space %v\n", line,s,nm,lm)
		//if nono_match(line, s) && line[ln_s+1]!='#' {
		if nm && lm {
			// test the rest & get it options.
			// if it returns 0 then the rest didn't match
			v:=nono_arrangements_memorised(line[ln_s+1:], vals[1:], memorise)
			//fmt.Printf("nono_arrangements '%s' %v gives %v\n", line[ln_s+1:], vals[1:], v)
			result+= v
		}
	}
	memorise[lookup]=result
	return result
}

func day12a(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	total:=0
	for _, line := range lines {
		total+= nono_arrangements(split_nono_line(line))
	}
	fmt.Printf("total %d\n", total)
}

func day12b(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	start_time := time.Now()
	cache:=make(map[string]int)

	total := 0
	for idx, line := range lines {
		l, v := split_nono_line(line)
		l2 := ""
		var v2 []int
		for i := 0; i < 5; i++ {
			if i > 0 {
				l2 += "?"
			}
			l2 += l

			v2 = append(v2, v...)
		}

		total += nono_arrangements_memorised(l2, v2, cache)
		taken := time.Now().Sub(start_time).Seconds()
		fmt.Printf("Line %d time taken %.2fs #cache %d\n", idx, taken, len(cache))
	}
	fmt.Printf("total %d\n", total)
}

