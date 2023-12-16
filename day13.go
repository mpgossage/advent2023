package main

import (
	"fmt"
)

// part A failed:
// 20700 too low
// several didn't have an mirror (I wonder..)

// Part B is a pain:
// the example given is incorrect
// with the smudge change the grid can have 2 mirrors (vert & horiz)
// but only the very is reported
// changing design to check for vert & only report that if its found
// (worked but not happy about question)

func parse_multi_grid(filename string) [][]string {
	lines, err := LoadLines(filename)
	if err != nil {
		fmt.Printf("%v", err)
		return nil
	}

	var result [][]string =nil

	var grid []string= nil
	for _,line:=range lines {
		if len(line) == 0 {
			result = append(result, grid)
			grid = nil
		} else {
			grid=append(grid, line)
		}
	}
	result = append(result, grid)
	return result
}

func look_for_grid_mirrors(grid []string) int {
	height:=len(grid)
	width:=len(grid[0])

	result:=0
	// look for a vertical line between index 1 & width-2 (it cannot be on the edge)
	for x:= 1; x< width; x++ {
		match:=true
		for lx:=0; match && lx < x; lx++ {
			rx:=2*x-1-lx
			if rx>=width {
				continue
			}
			//fmt.Printf("x=%d compare %d and %d\n",x,lx,rx)
			for y:= 0; y<height;y++{
				if grid[y][lx]!=grid[y][rx]{
					match=false
					break
				}
			}
		}
		if match {
			fmt.Printf("found vert mirror between %d-%d\n",x,x+1)
			// note: puzzles grids are 1+, so we use x
			result = x
			break
		}
	}
	// look for a horizontal line between index 1 & height-2 (it cannot be on the edge)
	for y:= 1; y< height; y++ {
		match:=true
		for ty:=0; match && ty<y;ty++{
			by:=2*y-1-ty
			if by>=height {
				continue
			}
			//fmt.Printf("y=%d compare %d and %d\n",y,ty,by)
			if grid[ty]!=grid[by]{
				match=false
				break
			}
		}
		if match {
			fmt.Printf("found horiz mirror between %d-%d\n",y,y+1)
			// note: puzzles grids are 1+, so we use y
			result += y*100
			break
		}
	}

	return result
}

func day13a(filename string) {
	grids:= parse_multi_grid(filename)
	fmt.Printf("found %d grids\n", len(grids))

	total:=0
	for g,grid:=range grids{
		val:=look_for_grid_mirrors(grid)
		fmt.Printf("examine grid %d result %d\n",g, val)
		total+=val
	}
	//look_for_grid_mirrors(grids[1])
	fmt.Printf("total %d\n",total)
}

func look_for_grid_near_mirrors(grid []string) int {
	// this is the same as above, but it looks for the differences for a given side
	// if the difference is 0 then its a real mirror
	// if the difference is 1 then its a smudge mirror
	// assuming the actual location of smudge don't matter
	height:=len(grid)
	width:=len(grid[0])

	// look for a horizontal line between index 1 & height-2 (it cannot be on the edge)
	for y:= 1; y< height; y++ {
		diff:=0
		for ty:=0; ty<y;ty++{
			by:=2*y-1-ty
			if by>=height {
				continue
			}
			//fmt.Printf("y=%d compare %d and %d\n",y,ty,by)
			for x:= 0; x<width;x++{
				if grid[ty][x]!=grid[by][x] {
					diff++
				}
			}
		}
		if diff==1 {
			fmt.Printf("found smudge horiz mirror between %d-%d\n",y,y+1)
			// note: puzzles grids are 1+, so we use y
			return y*100
		}
	}

	// look for a vertical line between index 1 & width-2 (it cannot be on the edge)
	for x:= 1; x< width; x++ {
		diff:=0
		for lx:=0; lx < x; lx++ {
			rx:=2*x-1-lx
			if rx>=width {
				continue
			}
			//fmt.Printf("x=%d compare %d and %d\n",x,lx,rx)
			for y:= 0; y<height;y++{
				if grid[y][lx]!=grid[y][rx] {
					diff++
				}
			}
		}
		if diff==1 {
			fmt.Printf("found smudge vert mirror between %d-%d\n",x,x+1)
			// note: puzzles grids are 1+, so we use x
			return x
		}
	}

	return 0
}

func day13b(filename string) {
	grids:= parse_multi_grid(filename)
	fmt.Printf("found %d grids\n", len(grids))

	total:=0
	for g,grid:=range grids{
		val:=look_for_grid_near_mirrors(grid)
		fmt.Printf("examine grid %d result %d\n",g, val)
		total+=val
	}
	fmt.Printf("total %d\n",total)
}
