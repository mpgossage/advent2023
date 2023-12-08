package main

import (
	"fmt"
	"strings"
	"time"
)

type SeedMapping struct {
	Name string
	Mapping [][]int
}

func (s SeedMapping) Transform(x int) int {
	for _, m := range s.Mapping {
		start, end, delta := m[1], m[1]+m[2]-1, m[0]-m[1]
		//fmt.Printf("start %d end %d delta %d\n", start, end, delta)
		if start<=x && x<=end {
			return x+delta
		}
	}
	return x
}


func ParseSeeds(filename string) (seeds []int, mappings []SeedMapping) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	seeds = StringToIntArray(strings.ReplaceAll(lines[0], "seeds:",""), " ")

	// discard the "seeds" line & empty line
	var mapping SeedMapping
	for _, line := range lines[2:] {
		if len(line) == 0 {
			mappings = append(mappings, mapping)
			mapping.Mapping = nil
			continue
		}

		if strings.ContainsAny(line[:1], "1234567890") {
			mapping.Mapping = append(mapping.Mapping, StringToIntArray(line, " "))
		} else {
			mapping.Name = line
		}
	}
	mappings = append(mappings, mapping)
	return
}

func day05a(filename string)  {
	seeds, mappings := ParseSeeds(filename)
	fmt.Printf("%+v\n%v\n", seeds, mappings)
	/*	seed := 50//seeds[0]
		v:= mappings[0].Transform(seed)
		fmt.Printf("seed %d transformed to %d\n", seed, v)
	*/

	nearest := 1000 * 1000 * 1000 * 1000 // arbitrary big num

	for _, seed := range seeds {
		//fmt.Printf("seed %d\n", seed)
		for _,m := range mappings {
			seed = m.Transform(seed)
			//fmt.Printf("transform %s to %d\n", m.Name, seed)
		}
		nearest = Min(nearest, seed)
	}

	fmt.Printf("nearest %d", nearest)
}

func day05b(filename string)  {
	// suspecting this will not work, but lets try the simplest method and see what happens
	// took 5 minutes, but worked
	seeds, mappings := ParseSeeds(filename)
	fmt.Printf("%+v\n%v\n", seeds, mappings)

	nearest := 1000 * 1000 * 1000 * 1000 // arbitrary big num

	startTime := time.Now()
	// seeds are now in pairs
	numPairs := len(seeds) / 2
	for i:=0; i< numPairs; i++ {
		start, end := seeds[2*i], seeds[2*i]+seeds[2*i+1]
		for seed :=start; seed < end; seed++ {
			val:= seed
			//fmt.Printf("seed %d", val)
			for _,m := range mappings {
				val = m.Transform(val)
			}
			//fmt.Printf(" becomes %d\n", val)
			nearest = Min(nearest, val)
		}
		// ticker
		taken := time.Now().Sub(startTime)
		fmt.Printf("pair %d/%d nearest %d taken %.f seconds\n", i, numPairs, nearest, taken.Seconds())
	}


	fmt.Printf("nearest %d", nearest)
}