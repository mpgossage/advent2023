package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Part A was simple
// I expected trouble from part B, but not this kind
// going to try direct approach, but I have doubts
// probably going to have to be a full poly size routine
// this failed because I didn't know how to account for the width of the path
// Eventually looked online and got an answer in ~5 minutes

func lavaduct_to_coords(lines []string) []Coord {
	var result []Coord
	loc := Coord{0, 0}
	for _, l := range lines {
		arr := strings.Split(l, " ")
		var dir Compass
		switch arr[0] {
		case "U":
			dir = North
		case "D":
			dir = South
		case "L":
			dir = West
		case "R":
			dir = East
		}
		dist, _ := strconv.Atoi(arr[1])
		// move dist steps in dir
		for i := 0; i < dist; i++ {
			loc = MoveCoordCompass(loc, dir)
			result = append(result, loc)
		}
	}
	return result
}

func day18a(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	// for now simplest thing that works
	// parse data and change into a polygon
	// can then use point_in_poly() from day10
	// will need to add all points separately
	path := lavaduct_to_coords(lines)
	fmt.Printf("%v\n", path)

	// to poly & find the grid size are same time
	// & map lookup for finding stuff on the line
	maxx, minx, maxy, miny := 0, 0, 0, 0
	line := make(map[Coord]bool)
	var poly [][]float32
	for _, p := range path {
		poly = append(poly, []float32{float32(p.x) + 0.001, float32(p.y) + 0.001})
		maxx = Max(maxx, p.x)
		minx = Min(minx, p.x)
		maxy = Max(maxy, p.y)
		miny = Min(miny, p.y)
		line[p] = true
	}

	total := 0
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			// quick check if on the line
			if _, found := line[Coord{x, y}]; found {
				total++
			} else if point_in_poly(poly, float32(x), float32(y)) {
				total++
			}
		}
	}

	fmt.Printf("total %d\n", total)
}

func lavaduct_to_coords2(lines []string) []Coord {
	loc := Coord{0, 0}
	var result []Coord
	for _, l := range lines {
		arr := strings.Split(l, " ")
		// item is "(#70c710)"  we want 70c71 & 0 from it
		dist, _ := strconv.ParseInt(arr[2][2:7], 16, 32)
		var dir Compass
		switch arr[2][7] {
		case '0':
			dir = East
		case '1':
			dir = South
		case '2':
			dir = West
		case '3':
			dir = North
		}
		fmt.Printf("'%s' %d %d\n", arr[2], dir, dist)
		for i := 0; i < int(dist); i++ {
			loc = MoveCoordCompass(loc, dir)
		}
		result = append(result, loc)
	}
	return result
}

func day18b(filename string) {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	path := lavaduct_to_coords2(lines)
	fmt.Printf("%v\n", path)

	// poly area algorithm:
	area := 0
	ln := len(path)
	for i := 0; i < ln; i++ {
		j := (i + 1) % ln
		area += path[i].x * path[j].y
		area -= path[i].y * path[j].x
	}
	area /= 2
	// ISSUE: considering the thickness of the cut!
	// just add the perimeter of the area+1
	// (mentioned in comments of https://www.youtube.com/watch?v=UNimgm_ogrw )
	perimeter := 0
	for i := 0; i < ln; i++ {
		j := (i + 1) % ln
		// PS. Manhatten is a bit of an overkill, but good enough
		perimeter += ManhattenDistance(path[i], path[j])
	}

	fmt.Printf("total %d\n", area+perimeter/2+1)
}
