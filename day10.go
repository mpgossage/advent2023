package main

import (
	"fmt"
)

var pipe_exits =map[uint8][]Compass{
	'|': {North, South},
	'-': {East, West},
	'L': {North,East},
	'J': {North, West},
	'7': {South, West},
	'F': {South, East},
}

func locate_pipe_start_point(grid []string) (pos Coord, exits []Compass) {
	height, width := len(grid), len(grid[0])
	// find 'S'
	for y, line := range grid {
		for x, c := range line {
			if c == 'S' {
				pos = Coord{x, y}
				break
			}
		}
	}
	// looking a 4 adjacent places for inputs
	if pos.x > 0 {
		// check to the west of start, if either exit goes east, then we must have a west exit
		e, found := pipe_exits[grid[pos.y][pos.x-1]]
		if found && (e[0] == East || e[1] == East) {
			exits = append(exits, West)
		}
	}
	if pos.x < width-1 {
		// check east
		e, found := pipe_exits[grid[pos.y][pos.x+1]]
		if found && (e[0] == West || e[1] == West) {
			exits = append(exits, East)
		}
	}
	if pos.y > 0 {
		// check north
		e, found := pipe_exits[grid[pos.y-1][pos.x]]
		if found && (e[0] == South || e[1] == South) {
			exits = append(exits, North)
		}
	}
	if pos.y < height-1 {
		// check south
		e, found := pipe_exits[grid[pos.y+1][pos.x]]
		if found && (e[0] == North || e[1] == North) {
			exits = append(exits, South)
		}
	}
	return
}

func traverse_pipe(pos CoordCompass, grid []string) (newpos CoordCompass)  {
	newpos.Coord = MoveCoordCompass(pos.Coord,pos.dir)
	opposite_direction := OppositeCompass(pos.dir)
	exits := pipe_exits[grid[newpos.y][newpos.x]]
	// we are looking for the exit which is not the opposite direction
	// eg. if we moved east, then we should ignore the west & take the other direction
	for _,d:=range exits {
		if d!= opposite_direction {
			newpos.dir = d
		}
	}
	return
}

func day10a(filename string) {
	grid, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	pos,exits:=locate_pipe_start_point(grid)
	fmt.Printf("pos %v exits %v\n", pos, exits)
	var routes[]CoordCompass
	for _,exit := range exits{
		routes =append(routes, CoordCompass{Coord: pos, dir: exit})
	}
	step:=0
	for {
		fmt.Printf("step: %d routes %v\n", step, routes)
		step++
		for i,r:=range routes {
			routes[i]=traverse_pipe(r,grid)
		}
		// assuming/requiring len=2
		if routes[0].Coord == routes[1].Coord {
			break
		}
	}
	fmt.Printf("step: %d routes %v\n", step, routes)
	fmt.Printf("result %d\n", step)
}

func point_in_poly(polygon [][]float32, x,y float32) bool {
	// https://www.algorithms-and-technologies.com/point_in_polygon/python
	odd := false
	j:= len(polygon) -1
	for i, polyi :=range polygon {
		polyj:=polygon[j]
		if ((polyi[1]> y) != (polyj[1] > y)) &&
					(x < (polyj[0]-polyi[0]) * (y - polyi[1])/(polyj[1]-polyi[1])+polyi[0]) {
			odd = !odd
		}
		j = i
	}
	return odd
}

func day10b(filename string) {
	grid, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	pos, exits := locate_pipe_start_point(grid)
	fmt.Printf("pos %v exits %v\n", pos, exits)

	start := CoordCompass{Coord: pos, dir: exits[0]}
	// creating a FP path for a point in poly routine
	// it requires a slight offset at point in poly will fail if lines are exactly horizontal
	var path [][]float32
	// list of places in the path (for removal)
	ipath := make(map[Coord]bool)
	current := start
	for {
		current = traverse_pipe(current, grid)
		path = append(path, []float32{float32(current.x) + 0.001, float32(current.y) + 0.001})
		ipath[current.Coord] = true
		if current.Coord == start.Coord {
			break
		}
	}
	fmt.Printf("path %v\n", path)
	// for all cells (excluding those on the path), check if in/out of area
	result:=0
	for y := range grid {
		for x:= range grid[0] {
			if _, found := ipath[Coord{x,y}]; found == false {
				in := point_in_poly(path, float32(x), float32(y))
				//fmt.Printf("point %d,%d in %v\n", x,y,in)
				if in {
					result++
				}
			}
		}
	}
	fmt.Printf("result %d\n",result)
}