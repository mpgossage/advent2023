package main

import "fmt"

func expand_galaxies(grid []string, galaxies []Coord, amount int) []Coord  {
	// expand galaxy:
	// look for the empty lines & then shift right/down accordingly
	var empty_x, empty_y []int
	for x:= range grid[0] {
		empty := true
		for y:= range grid {
			if grid[y][x]=='#' {
				empty = false
				break
			}
		}
		if empty {
			empty_x=append(empty_x,x)
		}
	}
	for y:= range grid {
		empty := true
		for x:= range grid[0] {
			if grid[y][x]=='#' {
				empty = false
				break
			}
		}
		if empty {
			empty_y=append(empty_y,y)
		}
	}
	fmt.Printf("emptyx=%v emptyy=%v\n",empty_x,empty_y)

	// expand right & down
	for i,g:= range galaxies {
		newg := g
		for _,x:=range empty_x {
			if g.x>x {
				newg.x+=amount
			}
		}
		for _,y:=range empty_y{
			if g.y>y {
				newg.y+=amount
			}
		}
		galaxies[i]=newg
	}
	return galaxies
}

func day11a(filename string)  {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	var galaxies []Coord
	for y, line := range lines {
		for x,c:= range line {
			if c == '#' {
				galaxies= append(galaxies, Coord{x,y})
			}
		}
	}
	fmt.Printf("Galaxies %v\n",galaxies)

	galaxies = expand_galaxies(lines, galaxies, 1)
	fmt.Printf("Galaxies %v\n",galaxies)
	// should be sorted so thats not needed

	// total distances
	total:=0
	for ia,a:=range galaxies {
		for ib,b := range galaxies {
			if ib>ia {
				total+=ManhattenDistance(a,b)
			}
		}
	}
	fmt.Printf("result %d\n", total)
}

func day11b(filename string) {
	// I'm surprised I got the algol right in part A
	// it was not do use a grid, but instead to use the coords
	// if I had done a grid, I would have terabyte grids!!
	// I need to split the expand fn out with a parameter on the amount

	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	var galaxies []Coord
	for y, line := range lines {
		for x,c:= range line {
			if c == '#' {
				galaxies= append(galaxies, Coord{x,y})
			}
		}
	}
	fmt.Printf("Galaxies %v\n",galaxies)
	// 10 times as big, so its 10-1 (9)
	galaxies = expand_galaxies(lines, galaxies, 1000000-1)
	fmt.Printf("Galaxies %v\n",galaxies)

	// total distances
	total:=0
	for ia,a:=range galaxies {
		for ib,b := range galaxies {
			if ib>ia {
				total+=ManhattenDistance(a,b)
			}
		}
	}
	fmt.Printf("result %d\n", total)
}